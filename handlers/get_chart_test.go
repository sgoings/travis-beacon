package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/sgoings/travis-beacon/chart"
	"github.com/sgoings/travis-beacon/storage"
)

var chartRouter *mux.Router
var db storage.DB

func init() {
	db = storage.NewInMemoryDB()
	chartRouter = NewRouter(db).(*mux.Router)
}

func TestNonexistentChart(t *testing.T) {
	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/charts/name", nil)
	if err != nil {
		t.Fatal("Creating 'GET /charts/name' request failed!")
	}

	request.Header.Add("Content-Type", "application/json")

	chartRouter.ServeHTTP(response, request)

	if response.Code != http.StatusNotFound {
		t.Fatal("Nonexistent chart was found?")
	}
}

func TestExistentChart(t *testing.T) {
	db.Set(chart.Chart{Name: "redis-standalone", Status: 75})

	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/charts/redis-standalone", nil)
	if err != nil {
		t.Fatal("Creating 'GET /charts/redis-standalone' request failed!")
	}

	request.Header.Add("Content-Type", "application/json")

	chartRouter.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatal("redis-standalone chart not found")
	}
}
