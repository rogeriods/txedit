package routes

import (
	"net/http"
	"text/template"

	"rogeriods/txedit/controllers"
	"rogeriods/txedit/models"

	"github.com/google/uuid"
)

var tmpl = template.Must(template.ParseGlob("templates/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	notes := controllers.SelectAll()
	tmpl.ExecuteTemplate(w, "Index", notes)
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("id")
	note := controllers.SelectById(nId)
	tmpl.ExecuteTemplate(w, "Edit", note)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	controllers.Delete(id)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

// ================= //
// Form POST actions //
// ================= //

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		note := models.Note{}
		note.ID = uuid.NewString()
		note.Title = r.FormValue("txtTitle")
		note.Note = r.FormValue("txtContent")

		controllers.InsOrUp(note, true)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		note := models.Note{}
		note.ID = r.FormValue("txtId")
		note.Title = r.FormValue("txtTitle")
		note.Note = r.FormValue("txtContent")

		controllers.InsOrUp(note, false)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
