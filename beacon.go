package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sgoings/travis-beacon/handlers"
)

func addChartHandlers(router *mux.Router) {
	subRouter := router.PathPrefix("/charts").
		Headers("Content-Type", "application/json").
		Subrouter()

	subRouter.HandleFunc("/", handlers.GetCharts).Methods("GET")
	subRouter.HandleFunc("/{name}", handlers.GetChart).Methods("GET")
}

func addTravisHandlers(router *mux.Router) {
	subRouter := router.PathPrefix("/charts").
		Headers("Content-Type", "application/x-www-form-urlencoded").
		MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		travisToken := r.Header.Get("Authorization")
		if travisToken == "" {
			return false
		}
		return true
	}).Subrouter()

	subRouter.HandleFunc("/", handlers.UpdateCharts).Methods("POST")
}

func addHandlers(router *mux.Router) {
	addChartHandlers(router)
	addTravisHandlers(router)
}

func main() {
	r := mux.NewRouter()
	addHandlers(r)

	port := 8000
	serveString := ":" + strconv.Itoa(port)

	log.Printf("[INFO] Listening on %s", serveString)
	http.ListenAndServe(serveString, r)
}
