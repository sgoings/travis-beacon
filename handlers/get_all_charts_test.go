package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/sgoings/travis-beacon/chart"
	"github.com/sgoings/travis-beacon/storage"
)

var allChartRouter *mux.Router
var allChartDB storage.DB

func init() {
	allChartDB = storage.NewInMemoryDB()
	allChartRouter = NewRouter(allChartDB).(*mux.Router)
}

func TestNoCharts(t *testing.T) {
	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/charts/", nil)
	if err != nil {
		t.Fatal("Creating 'GET /charts/' request failed!")
	}

	request.Header.Add("ContentType", "application/json")

	allChartRouter.ServeHTTP(response, request)

	if response.Code != http.StatusNotFound {
		t.Fatal("There were charts?", response.Code)
	}
}

func TestSomeCharts(t *testing.T) {
	allChartDB.Set(chart.Chart{Name: "redis-standalone", Status: 75})

	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/charts/", nil)
	if err != nil {
		t.Fatal("Creating 'GET /charts/' request failed!")
	}

	request.Header.Add("ContentType", "application/json")

	allChartRouter.ServeHTTP(response, request)

	if response.Code != http.StatusNotFound {
		t.Fatal("There were charts?", response.Code)
	}
}
