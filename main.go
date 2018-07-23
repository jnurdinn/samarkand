package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

func newRouter() *mux.Router {
  router := mux.NewRouter()
  staticFileDirectory := http.Dir("./assets/")
  staticFileHandler := http.StripPrefix("/", http.FileServer(staticFileDirectory))
  router.PathPrefix("/").Handler(staticFileHandler).Methods("GET")
  return router
}

func main(){
  router := newRouter()
  http.ListenAndServe(":8080", router)
}
