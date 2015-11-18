package main

import (
  "fmt"
  "net/http"

  "github.com/gorilla/mux"
)

func GetChart(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  name := vars["name"]
  fmt.Fprintf(w, "Info about chart: %s\n", name)
}

func GetCharts(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("List of all charts\n"))
}

func AddChartRoutes(router *mux.Router) {
  subRouter := router.PathPrefix("/charts").Subrouter()

  subRouter.HandleFunc("/", GetCharts).Methods("GET")
  subRouter.HandleFunc("/{name}", GetChart).Methods("GET")
}

func AddRoutes(router *mux.Router) {
  AddChartRoutes(router)
}

func main() {
  r := mux.NewRouter()
  AddRoutes(r)

  http.ListenAndServe(":8000", r)
}
