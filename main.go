package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

func newRouter() *mux.Router {
  router := mux.NewRouter()
  router.HandleFunc("/", handler).Methods("GET")

  staticFileDirectory := http.Dir("./assets/")
  staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
  router.PathPrefix("/assets").Handler(staticFileHandler).Methods("GET")

  return router
}

func main(){
  router := newRouter()
  http.ListenAndServe(":8080", router)
}

func handler(w http.ResponseWriter, router *http.Request){
  fmt.Fprintf(w, "SAMARKAND:8080")
}
