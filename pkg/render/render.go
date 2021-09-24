package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/soguazu/bookings/pkg/config"
	"github.com/soguazu/bookings/pkg/modules"
)

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

var functions = template.FuncMap{}

func AddDefaultData(td *modules.TemplateData) *modules.TemplateData {
	return td
}

func RenderTemplates(res http.ResponseWriter, tmpl string, td *modules.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplatesCache()
	}

	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("Page not found")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(res)

	if err != nil {
		println("Error writing template to browser", err)
	}

}

func CreateTemplatesCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.gohtml")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.gohtml")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
		}

		if err != nil {
			return myCache, err
		}

		myCache[name] = ts
	}

	return myCache, nil
}
