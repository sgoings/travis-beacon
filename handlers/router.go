package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sgoings/travis-beacon/storage"
)

// NewRouter can route traffic to any handler
func NewRouter(db storage.DB) http.Handler {
	router := mux.NewRouter()
	addChartHandlers(router, db)
	addTravisHandlers(router, db)

	return router
}

func addChartHandlers(router *mux.Router, db storage.DB) {
	subRouter := router.PathPrefix("/charts").
		Headers("Content-Type", "application/json").
		Subrouter()

	subRouter.Handle("/", GetAllCharts(db)).Methods("GET")
	subRouter.Handle("/{name}", GetChart(db)).Methods("GET")
}

func addTravisHandlers(router *mux.Router, db storage.DB) {
	subRouter := router.PathPrefix("/charts").
		Headers("Content-Type", "application/x-www-form-urlencoded").
		MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		travisToken := r.Header.Get("Authorization")
		if travisToken == "" {
			return false
		}
		return true
	}).Subrouter()

	subRouter.Handle("/", UpdateCharts(db)).Methods("POST")
}
