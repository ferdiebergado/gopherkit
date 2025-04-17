package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"

	ghttp "github.com/ferdiebergado/gopherkit/http"
)

const (
	templatesDir = "templates"
	layoutFile   = "layout.html"
)

// Sends a JSON response
func JSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set(ghttp.HeaderContentType, ghttp.MimeJSON)
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		ServerError(w, fmt.Errorf("encode json: %w", err))
	}
}

func ServerError(w http.ResponseWriter, err error) {
	slog.Error("server error", "reason", err, "stack_trace", string(debug.Stack()))
	http.Error(w, "An error occurred.", http.StatusInternalServerError)
}

// Sends an HTML response
func HTML(w http.ResponseWriter, data any, templateFiles ...string) {
	layoutTemplate := filepath.Join(templatesDir, layoutFile)
	targetTemplates := []string{layoutTemplate}

	targetTemplates = append(targetTemplates, templateFiles...)

	funcMap := templateFuncs()

	templates, err := template.New("template").Funcs(funcMap).ParseFiles(targetTemplates...)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer

	if err := templates.ExecuteTemplate(&buf, layoutFile, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = buf.WriteTo(w)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Creates a template funcmap
func templateFuncs() template.FuncMap {
	return template.FuncMap{
		"attr": func(s string) template.HTMLAttr {
			return template.HTMLAttr(s)
		},
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
		"url": func(s string) template.URL {
			return template.URL(s)
		},
		"js": func(s string) template.JS {
			return template.JS(s)
		},
		"jsstr": func(s string) template.JSStr {
			return template.JSStr(s)
		},
		"css": func(s string) template.CSS {
			return template.CSS(s)
		},
	}
}

// ParsePages recursively parses a given directory containing html templates.
// Each page is parsed against the layout template.
// It returns a map containing the name of the template as key and the parsed template as the value.
func ParsePages(templateDir string, layoutTmpl *template.Template) (map[string]*template.Template, error) {
	tmplMap := make(map[string]*template.Template)
	err := fs.WalkDir(os.DirFS(templateDir), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		const suffix = ".html"
		if !d.IsDir() && strings.HasSuffix(path, suffix) {
			name := strings.TrimPrefix(path, "/")
			name = strings.TrimSuffix(name, suffix)
			tmplMap[name] = template.Must(template.Must(layoutTmpl.Clone()).ParseFiles(filepath.Join(templateDir, path)))
			slog.Debug("parsed page", "path", path, "name", name)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("load pages templates: %w", err)
	}

	return tmplMap, nil
}
