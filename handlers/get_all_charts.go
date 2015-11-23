package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sgoings/travis-beacon/storage"
)

// GetAllCharts responds with a list of charts and their quality metrics
func GetAllCharts(db storage.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		returning := db.GetAll()

		log.Printf("[DEBUG] Returning %d charts", returning)

		json.NewEncoder(w).Encode(returning)
	})
}
