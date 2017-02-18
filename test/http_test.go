package testdemo

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	path := "http://localhost:9999/test"

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", path, nil)
	http.DefaultServeMux.ServeHTTP(res, req)

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Test Request Error %s", err)
	}

	nt := time.Now()
	if string(response) != string(nt.YearDay()) || err != nil {
		t.Errorf("Expected %s, got %s", string(nt.YearDay()), string(response))
	}
}
