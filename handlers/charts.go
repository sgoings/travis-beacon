package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// TravisWebhookBody is the body of the POST
type TravisWebhookBody struct {
	Payload TravisPayload
}

// TravisPayload is the stuff
type TravisPayload struct {
	Matrix []TravisJob
}

// TravisJob gives nice access to Job
type TravisJob struct {
	Log string
}

// ChartQuality describes a Chart's quality
type ChartQuality struct {
	Name   string
	Status int `json:"status"`
}

// ChartQualityList is used to decode the output
// from helm lint in the Travis job log
type ChartQualityList struct {
	Charts map[string]ChartQuality
}

var chartList = new(ChartQualityList)

// GetChart responds with a single chart and its quality metrics
func GetChart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	json.NewEncoder(w).Encode(chartList.Charts[name])
}

// GetCharts responds with a list of charts and their quality metrics
func GetCharts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(chartList)
}

// UpdateCharts updates all charts
func UpdateCharts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	startToken := "[helm-lint] <start>"
	endToken := "[helm-lint] <end>"

	if vars["startToken"] != "" {
		startToken = vars["startToken"]
	}

	if vars["endToken"] != "" {
		endToken = vars["endToken"]
	}

	var travisData TravisWebhookBody
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&travisData); err != nil {
		log.Printf("[ERROR] json couldn't be decoded: %s\n", err)
	}

	travisLog := travisData.Payload.Matrix[0].Log
	output := extractJSONFromLog(startToken, endToken, travisLog)

	for _, value := range output.Charts {
		updateChart(value)
	}
}

func extractJSONFromLog(startKey string, endKey string, logOutput string) ChartQualityList {
	startIndex := strings.LastIndex(logOutput, startKey) + len(startKey)
	endIndex := strings.LastIndex(logOutput, endKey)

	var ChartQualityList ChartQualityList
	err := json.Unmarshal([]byte(logOutput[startIndex:endIndex]), &ChartQualityList)

	if json.Unmarshal([]byte(logOutput[startIndex:endIndex]), &ChartQualityList); err != nil {
		log.Printf("[ERROR] json couldn't be transformed into a list of charts %s\n", err)
	}

	return ChartQualityList
}

func updateChart(chart ChartQuality) {
	chartList.Charts[chart.Name] = chart
	log.Printf("[DEBUG] %s (%d/100)\n", chart.Name, chart.Status)
}
