package main

import (
	"fmt"
	"net/http"
	"time"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello from Cisco Shipped!</h1>\n")
	fmt.Fprintf(w, "<h1>Hello from Sheikh! --- Published N Build form Dockerfile--golang_job</h1>\n")
	time1 := time.Now().Second()
	fmt.Fprintf(w,"<div>---Timer Started---</div>\n")
	count := 1
	for time1 <=  time.Now().Second(){
		fmt.Println(w,time1)
		if time1 == 59 {
			fmt.Println(w,"***- ",count," Min. -***\n")
			count = count+1
		}
		time.Sleep(1*time.Second)
		if time1 > 50 {
			break;
		}
		time1 = time.Now().Second()
	}
}

func main() {
    http.HandleFunc("/", defaultHandler)
    fmt.Println("Example app listening at http://localhost:8888")
    http.ListenAndServe(":8888", nil)
}
