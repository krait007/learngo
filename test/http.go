package testdemo

import (
	"fmt"
	"net/http"
	"time"
)

func demoHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Fprintln(w, t.YearDay())
}

func StartServer() {
	http.HandleFunc("/test", demoHandler)
	http.ListenAndServe(":9999", nil)
}
