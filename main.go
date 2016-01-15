package main

import (
	"fmt"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello from Cisco Shipped!</h1>\n")
	fmt.Fprintf(w, "<h1>Hello from Sheikh! --- Published N Build form Dockerfile</h1>\n")
	time1 := time.Now().Second()
	fmt.Fprintf(w,"---Timer Started---\n")
	count := 1
	for time1 <=  time.Now().Second(){
		fmt.Fprintf(w,time1,"\n")
		if time1 == 59 {
			fmt.Fprintf(w,"***- ",count," Min. -***\n")
			count = count+1
		}
		time.Sleep(1*time.Second)
		
		time1 = time.Now().Second()
	}
}

func main() {
    http.HandleFunc("/", defaultHandler)
    fmt.Println("Example app listening at http://localhost:8888")
    http.ListenAndServe(":8888", nil)
}
