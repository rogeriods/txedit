package main

import (
	"log"
	"net/http"

	"rogeriods/txedit/controllers"
	"rogeriods/txedit/routes"
)

func main() {
	log.Println("Server running on port 8080")

	http.HandleFunc("/", routes.Index)
	http.HandleFunc("/new", routes.New)
	http.HandleFunc("/edit", routes.Edit)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/delete", controllers.Delete)

	http.ListenAndServe(":8080", nil)
}
