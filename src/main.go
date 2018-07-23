package main

import (
  "net/http"
  "fmt"
   "strconv"
  "github.com/gorilla/mux"
  "github.com/BurntSushi/toml"
)

type Config struct {
  Enable  bool
  Port    int
  Dir     string
}

func newRouter(dir string) *mux.Router {
  router := mux.NewRouter()
  staticFileDirectory := http.Dir(dir)
  staticFileHandler := http.StripPrefix("/", http.FileServer(staticFileDirectory))
  router.PathPrefix("/").Handler(staticFileHandler).Methods("GET")
  return router
}

func main(){
  var conf Config
  if _, err := toml.DecodeFile("./config/config.toml", &conf); err != nil {
    fmt.Println(err)
  }
  if (conf.Enable){
    router := newRouter(conf.Dir)
    port := ":" + strconv.Itoa(conf.Port)
    http.ListenAndServe(port, router)
  }
}
