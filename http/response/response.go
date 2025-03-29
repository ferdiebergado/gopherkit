package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"path/filepath"
	"runtime/debug"

	ghttp "github.com/ferdiebergado/gopherkit/http"
)

const (
	templatesDir = "templates"
	layoutFile   = "layout.html"
)

// Sends a JSON response
func JSON(w http.ResponseWriter, r *http.Request, status int, v any) {
	w.Header().Set(ghttp.HeaderContentType, ghttp.MimeJSON)
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		ServerError(w, r, fmt.Errorf("encode json: %w", err))
	}
}

func ServerError(w http.ResponseWriter, r *http.Request, err error) {
	slog.Error("server error", "reason", err, "request", fmt.Sprint(r), "trace", string(debug.Stack()))
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
