package main

import (
	"log"
	"net/http"

	"rogeriods/txedit/routes"
)

func main() {
	log.Println("Server running on port 8080")

	http.HandleFunc("/", routes.Index)
	http.HandleFunc("/new", routes.New)
	http.HandleFunc("/edit", routes.Edit)
	http.HandleFunc("/insert", routes.Insert)
	http.HandleFunc("/update", routes.Update)
	http.HandleFunc("/delete", routes.Delete)

	http.ListenAndServe(":8080", nil)
}
