package handlers

import (
	"net/http"

	"github.com/ryanhoangt/go-web-bookings/pkg/config"
	"github.com/ryanhoangt/go-web-bookings/pkg/models"
	"github.com/ryanhoangt/go-web-bookings/pkg/render"
)

// Repo the repository used by handlers
var Repo *Repository

type Repository struct {
	appConfig *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(ac *config.AppConfig) *Repository {
	return &Repository{
		appConfig: ac,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.appConfig.SessionMgr.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	stringMap["test"] = "Hello again"

	remoteIP := m.appConfig.SessionMgr.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
