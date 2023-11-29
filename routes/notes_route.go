package routes

import (
	"net/http"
	"text/template"

	"rogeriods/txedit/models"
	"rogeriods/txedit/utils"
)

var tmpl = template.Must(template.ParseGlob("templates/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := utils.DBConn()
	selDB, err := db.Query("SELECT * FROM notes")
	if err != nil {
		panic(err.Error())
	}

	note := models.Note{}
	res := []models.Note{}

	for selDB.Next() {
		var id, title, content string
		err := selDB.Scan(&id, &content, &title)
		if err != nil {
			panic(err.Error())
		}

		note.ID = id
		note.Note = content
        note.Title = title 
		res = append(res, note)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := utils.DBConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM notes WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	note := models.Note{}

	for selDB.Next() {
		var id, title, content string
		err := selDB.Scan(&id, &content, &title)
		if err != nil {
			panic(err.Error())
		}

		note.ID = id
        note.Title = title 
		note.Note = content
	}

	tmpl.ExecuteTemplate(w, "Edit", note)
	defer db.Close()
}
