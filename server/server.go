package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Works so far.")
	})

	log.Print("Server listening on port ", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
