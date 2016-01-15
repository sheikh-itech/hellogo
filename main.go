package main

import (
	"fmt"
	"net/http"
	"time"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello from Cisco Shipped!</h1>\n")
	fmt.Fprintf(w, "<h1>Hello from Sheikh! --- Published N Build form Dockerfile--golang_job</h1>\n")
}

func main() {
    http.HandleFunc("/", defaultHandler)
    fmt.Println("Example app listening at http://localhost:8888")
    http.ListenAndServe(":8888", nil)
}
