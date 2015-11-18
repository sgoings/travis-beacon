package main

import (
  "testing"

  "net/http"
  "net/http/httptest"
  "github.com/gorilla/mux"
)

var m *mux.Router

func init() {
  m = mux.NewRouter()
  AddRoutes(m)
}

func TestGetCharts(t *testing.T) {
  response := httptest.NewRecorder()
  request, err := http.NewRequest("GET", "/charts/", nil)
  if err != nil {
    t.Fatal("Creating 'GET /charts/' request failed!")
  }

  m.ServeHTTP(response, request)

  if response.Code != http.StatusOK {
    t.Fatal("Charts base route failed to return properly")
  }
}

func TestGetChart(t *testing.T) {
  response := httptest.NewRecorder()
  request, err := http.NewRequest("GET", "/charts/name", nil)
  if err != nil {
    t.Fatal("Creating 'GET /charts/name' request failed!")
  }

  m.ServeHTTP(response, request)

  if response.Code != http.StatusOK {
    t.Fatal("Individual chart failed to display properly")
  }
}


