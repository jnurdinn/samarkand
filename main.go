package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

func main(){
  //initiate new router
  router := mux.NewRouter()

  //declare method where path is valid
  router.HandleFunc("/index", handler).Methods("GET")

  //pass router after declaring routes, port 8080
  http.ListenAndServe(":8080", router)
}

func handler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "SAMARKAND:8080")
}
