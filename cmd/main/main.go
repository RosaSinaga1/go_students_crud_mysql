package main

import (
	"fmt"
	"github.com/RosaSinaga1/go_students_crud_mysql/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func main() {
	// inisialisasi routes
	r := mux.NewRouter()
	routes.RegisterStudentsRoutes(r)
	http.Handle("/", r)
	fmt.Print("Starting Server localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
