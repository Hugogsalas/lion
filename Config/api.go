package main

import (
	"log"
	"net/http"
	"../Controllers"
	"github.com/gorilla/mux"
	"fmt"
)

func main() {
	fmt.Println("jajaja k creysi")
	fmt.Println("Iniciamos el server")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/User/create", controllers.CreateUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
