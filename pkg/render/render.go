// Package render is created for rendering the html pages and caches it.
package render

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
	"time"

	"github.com/MohummedSoliman/booking/pkg/config"
	"github.com/MohummedSoliman/booking/pkg/models"
	"github.com/justinas/nosurf"
)

var (
	app          *config.AppConfig
	pathTemplate = "./templates"
)

var functions = template.FuncMap{
	"humanDate":  HumanDate,
	"formatDate": FormatDate,
}

func NewTemplate(a *config.AppConfig) {
	app = a
}

// HumanDate return time in yyyy-mm-dd format
func HumanDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func FormatDate(t time.Time, format string) string {
	return t.Format(format)
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	td.IsAuthenticated = app.Session.Exists(r.Context(), "user_id")
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, td *models.TemplateData, temName string) error {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[temName]
	if !ok {
		return errors.New("can not get templates from cache")
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
		return err
	}
	return nil
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.*", pathTemplate))
	if err != nil {
		log.Println(err)
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		t, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.*", pathTemplate))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			t, err = t.ParseGlob(fmt.Sprintf("%s/*.layout.*", pathTemplate))
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = t
	}
	return myCache, nil
}
