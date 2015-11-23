package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sgoings/travis-beacon/chart"
	"github.com/sgoings/travis-beacon/storage"
	"github.com/sgoings/travis-beacon/travis"
)

// UpdateCharts updates all charts
func UpdateCharts(db storage.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		startToken := "[helm-lint] <start>"
		endToken := "[helm-lint] <end>"

		if vars["startToken"] != "" {
			startToken = vars["startToken"]
		}

		if vars["endToken"] != "" {
			endToken = vars["endToken"]
		}

		var travisData map[string]travis.WebhookPayload
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&travisData); err != nil {
			log.Printf("[ERROR] json couldn't be decoded: %s\n", err)
		}

		travisLog := travisData["payload"].Matrix[0].Log
		output := extractJSONFromLog(startToken, endToken, travisLog)

		for _, value := range output {
			db.Set(value)
		}
	})
}

func extractJSONFromLog(startKey string, endKey string, logOutput string) map[string]chart.Chart {
	startIndex := strings.LastIndex(logOutput, startKey) + len(startKey)
	endIndex := strings.LastIndex(logOutput, endKey)

	charts := make(map[string]chart.Chart)
	err := json.Unmarshal([]byte(logOutput[startIndex:endIndex]), &charts)

	if json.Unmarshal([]byte(logOutput[startIndex:endIndex]), &charts); err != nil {
		log.Printf("[ERROR] json couldn't be transformed into a list of charts %s\n", err)
	}

	return charts
}
