package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ryanhoangt/go-web-bookings/pkg/config"
	"github.com/ryanhoangt/go-web-bookings/pkg/models"
)

var appConfig *config.AppConfig

// NewTemplateRenderer sets the config for the template
// renderer package
func NewTemplateRenderer(ac *config.AppConfig) {
	appConfig = ac
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmplName string, td *models.TemplateData) {
	var tmplCache map[string]*template.Template

	if appConfig.UseCache {
		tmplCache = appConfig.TemplateCache
	} else {
		tmplCache, _ = CreateTemplateCache()
	}

	tmpl, ok := tmplCache[tmplName]
	if !ok {
		log.Fatal("Could not get template from template cache.")
	}

	buf := new(bytes.Buffer)
	td = addDefaultData(td)
	_ = tmpl.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to response.", err)
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pagesPath, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}
	layoutsPath, err := filepath.Glob("./templates/*.layout.tmpl")
	if err != nil {
		return cache, err
	}

	for _, pagePath := range pagesPath {
		name := filepath.Base(pagePath)
		tmpl, err := template.New(name).ParseFiles(pagePath)
		if err != nil {
			return cache, err
		}

		if len(layoutsPath) > 0 {
			tmpl, err = tmpl.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return cache, err
			}
		}

		cache[name] = tmpl
	}

	return cache, nil
}

func addDefaultData(td *models.TemplateData) *models.TemplateData {
	// TODO:
	return td
}
