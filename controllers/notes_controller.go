package controllers

import (
	"database/sql"

	"rogeriods/txedit/models"
	"rogeriods/txedit/utils"
)

// ========================================================== //
// This package is only responsible for database transactions //
// ========================================================== //

func InsOrUp(note models.Note, isNew bool) {
	db := utils.DBConn()

	var insData *sql.Stmt
	var err error

	// Treat if is update or insert new
	if !isNew {
		insData, err = db.Prepare("UPDATE notes SET title = ?, note = ? WHERE id = ?")
		if err != nil {
			panic(err.Error())
		}
		insData.Exec(note.Title, note.Note, note.ID)
	} else {
		insData, err = db.Prepare("INSERT INTO notes (id, title, note) VALUES (?, ?, ?)")
		if err != nil {
			panic(err.Error())
		}
		insData.Exec(note.ID, note.Title, note.Note)
	}

	defer db.Close()
}

func Delete(id string) {
	db := utils.DBConn()

	delForm, err := db.Prepare("DELETE FROM notes WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)

	defer db.Close()
}

func SelectAll() []models.Note {
	db := utils.DBConn()

	selDB, err := db.Query("SELECT * FROM notes")
	if err != nil {
		panic(err.Error())
	}

	note := models.Note{}
	notes := []models.Note{}

	// Fill the array
	for selDB.Next() {
		err := selDB.Scan(&note.ID, &note.Note, &note.Title)
		if err != nil {
			panic(err.Error())
		}
		notes = append(notes, note)
	}

	defer db.Close()

	return notes
}

func SelectById(id string) models.Note {
	db := utils.DBConn()

	note := models.Note{}

	selDB, err := db.Query("SELECT * FROM notes WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {
		err := selDB.Scan(&note.ID, &note.Note, &note.Title)
		if err != nil {
			panic(err.Error())
		}
	}

	defer db.Close()

	return note
}
