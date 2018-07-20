package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

func newRouter() *mux.Router {
  router := mux.NewRouter()
  router.HandleFunc("/index", handler).Methods("GET")
  return router
}

func main(){
  router := newRouter()
  http.ListenAndServe(":8080", router)
}

func handler(w http.ResponseWriter, router *http.Request){
  fmt.Fprintf(w, "SAMARKAND:8080")
}
