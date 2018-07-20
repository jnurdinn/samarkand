//main_test.go

package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
)

//function for testing handler
func TestHandler(t *testing.T){
  req, err := http.NewRequest("GET", "", nil)
  if err != nil {
    t.Fatal(err)
  }

  recorder := httptest.NewRecorder()

  hf := http.HandlerFunc(handler)
  hf.ServeHTTP(recorder, req)

  if status := recorder.Code; status != http.StatusOK{
    t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
  }

  expected := `SAMARKAND:8080`
  actual := recorder.Body.String()
  if actual != expected {
    t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
  }
}

func TestRouter(t *testing.T){
  router := newRouter()
  mockServer := httptest.NewServer(router)

  resp, err := http.Get(mockServer.URL + "/index")
  if err != nil {
    t.Fatal(err)
  }

  if resp.StatusCode != http.StatusOK {
    t.Errorf("Status should be ok, got %d", resp.StatusCode)
  }

  defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
  respString := string(b)
  expected := "SAMARKAND:8080"

  if respString != expected {
    t.Errorf("Response should be %s, got %s", expected, respString)
  }
}
