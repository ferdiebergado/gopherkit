package response

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"
)

const (
	templatesDir = "templates"
	layoutFile   = "layout.html"
)

// Sends a JSON response
func JSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		ServerError(w)
	}
}

func ServerError(w http.ResponseWriter) {
	http.Error(w, "an error occurred.", http.StatusInternalServerError)
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
