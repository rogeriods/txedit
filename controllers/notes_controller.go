package controllers

import (
	"net/http"

    "github.com/google/uuid"
    "rogeriods/txedit/utils"
)

func Insert(w http.ResponseWriter, r *http.Request) {
	db := utils.DBConn()
	if r.Method == "POST" {
        myId := uuid.NewString()
        title := r.FormValue("txtTitle")
		content := r.FormValue("txtContent")
		insForm, err := db.Prepare("INSERT INTO notes (id, title, note) VALUES (?, ?, ?)")
        if err != nil {
			panic(err.Error())
		}
		insForm.Exec(myId, title, content)
	}
	defer db.Close()
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := utils.DBConn()
	if r.Method == "POST" {
		id := r.FormValue("txtId")
        title := r.FormValue("txtTitle")
		content := r.FormValue("txtContent")
		insForm, err := db.Prepare("UPDATE notes SET title = ?, note = ? WHERE id = ?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(title, content, id)
	}
	defer db.Close()
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := utils.DBConn()
	id := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM notes WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)
	defer db.Close()
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
