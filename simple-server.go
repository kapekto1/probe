package main

import (
	"fmt"
	"net/http"
	"time"
	"os"
	"html/template"
)

var temp1 *template.Template

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func handler1(w http.ResponseWriter, r *http.Request) {
	err := temp1.Execute(w,nil)
	check(err)
	}


	func handler2(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	time_now :=	start.Format(time.Stamp)
	hostname, _ := os.Hostname()
	ip_addr := r.RemoteAddr
	fmt.Fprintf(w, "%s... You've hit %s from %s\n", time_now, hostname, ip_addr)
	fmt.Println(time.Since(start))
}

func main() {
	var err error
	temp1, err = template.ParseFiles("temp1.gohtml")
	check(err)
	http.HandleFunc("/", handler1)
	http.HandleFunc("/probe", handler2)
	http.ListenAndServe(":8080", nil)
}

