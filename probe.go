package main

import (
	"fmt"
	"net/http"
	"time"
	"os"
)

type Message struct {
	message string
	hostname string
	time string
	client_ip string
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}


func handler1(w http.ResponseWriter, r *http.Request) {
	m1 := new(Message)
	start := time.Now()
	m1.time = start.Format(time.Stamp)
	m1.hostname, _ = os.Hostname()
	m1.client_ip = r.RemoteAddr
	fmt.Fprintf(w, "%s... You've hit %s from %s \n", m1.time, m1.hostname, m1.client_ip)
	fmt.Println(time.Since(start))
}

func main() {
	var err error
	check(err)
	http.HandleFunc("/", handler1)
	http.ListenAndServe("0.0.0.0:8080", nil)
}

