package render

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func View(resp http.ResponseWriter, name string, data interface{}) error {
	filename := filepath.Join("./views", name+".html")
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		return err
	}

	resp.Header().Set("Content-Type", "text/html")
	return tmpl.Execute(resp, data)
}
