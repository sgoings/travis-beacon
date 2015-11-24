package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/sgoings/travis-beacon/handlers"
	"github.com/sgoings/travis-beacon/storage"
)

func main() {
	router := handlers.NewRouter(storage.NewInMemoryDB())

	port := 8000
	serveString := ":" + strconv.Itoa(port)

	log.Printf("[INFO] Listening on %s", serveString)
	http.ListenAndServe(serveString, router)
}
