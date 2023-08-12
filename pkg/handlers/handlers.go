package handlers

import (
	"net/http"
	"github.com/brajesh/bookings/pkg/config"
	"github.com/brajesh/bookings/pkg/render"
	"github.com/brajesh/bookings/pkg/models"
)
// repo the repository used by the handlers
var Repo *Repository

// repository is the repository type
type Repository struct{
	App *config.AppConfig
}
// ṆewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App: a,
	}
}
// New handlers sets the repository for the handlers
func NewHandlers(r *Repository){
	Repo=r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request){
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(),"remote_ip",remoteIP)

	render.RenderTemplate(w,"home.page.html",&models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request){
	// perform some logic
	stringMap:=make(map[string]string)
	stringMap["test"]="hello, again." 

	remoteIP := m.App.Session.GetString(r.Context(),"remote_ip")

	stringMap["remote_ip"]=remoteIP

	// send data to the template
	render.RenderTemplate(w,"about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}