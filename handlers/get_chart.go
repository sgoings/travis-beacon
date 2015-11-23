package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sgoings/travis-beacon/storage"
)

// GetChart responds with a single chart and its quality metrics
func GetChart(db storage.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]

		chart := db.Get(name)

		if chart.IsComplete() {
			log.Printf("[DEBUG] Chart being returned: %s\n", chart)
			json.NewEncoder(w).Encode(db.Get(name))
		} else {
			log.Printf("[DEBUG] Chart with name: %s not found", name)
			http.Error(w, "", http.StatusNotFound)
		}
	})
}
