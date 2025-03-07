package handlers

import (
	"github.com/bratovnikoff/go-booking-monolithic-web/pkg/config"
	"github.com/bratovnikoff/go-booking-monolithic-web/pkg/models"
	"net/http"

	"github.com/bratovnikoff/go-booking-monolithic-web/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	repo.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perfom some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send the data
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
