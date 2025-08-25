// Package render is created for rendering the html pages and caches it.
package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/MohummedSoliman/booking/pkg/config"
	"github.com/MohummedSoliman/booking/pkg/models"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, td *models.TemplateData, temName string) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[temName]
	if !ok {
		log.Fatal("Could not get tempate from templates cache")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td, r)
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.*")
	if err != nil {
		log.Println(err)
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		t, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.*")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			t, err = t.ParseGlob("./templates/*.layout.*")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = t
	}
	return myCache, nil
}
