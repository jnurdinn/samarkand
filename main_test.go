//main_test.go

package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestHandler(t *testing.T){
  //Form HTTP Request
  req, err := http.NewRequest("GET", "", nil)

  //If error persist
  if err != nil {
    t.Fatal(err)
  }

  //Recoder as target of http request
  recorder := httptest.NewRecorder()

  //HTTP handler
  hf := http.HandlerFunc(handler)

  //Serve HTTP request to recoder
  hf.ServeHTTP(recorder, req)

  //Check status code
  if status := recorder.Code; status != http.StatusOK{
    t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
  }

  //Check response body
  expected := `SAMARKAND:8080`
  actual := recorder.Body.String()
  if actual != expected {
    t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
  }
}
