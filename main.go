package main

import (
	"BankDatabase/model"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")

	model.InitialMigration()
	http.HandleFunc("/", serve)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serve(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello Mr.Amini!"))
}
