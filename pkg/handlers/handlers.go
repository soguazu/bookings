package handlers

import (
	"net/http"

	"github.com/soguazu/bookings/pkg/config"
	"github.com/soguazu/bookings/pkg/modules"
	"github.com/soguazu/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepository(a *config.AppConfig) *Repository {

	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(res http.ResponseWriter, req *http.Request) {
	remoteIP := req.RemoteAddr
	m.App.Session.Put(req.Context(), "remote_ip", remoteIP)
	render.RenderTemplates(res, "home.page.gohtml", &modules.TemplateData{})
}

func (m *Repository) About(res http.ResponseWriter, req *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, well again."
	remoteIP := m.App.Session.GetString(req.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplates(res, "about.page.gohtml", &modules.TemplateData{
		StringMap: stringMap,
	})
}
