package main

import (
	"fmt"
	"net/http"
	"time"
	"os"
)


type Message struct {
	hostname string
	time string
	client_ip string
	process_time time.Duration
	node_name string
}

func get_env_variables() string {
	return os.Getenv("MY_NODE_NAME")

}

func handler1(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	m1 := new(Message)
	m1.time = start.Format(time.Stamp)
	m1.hostname, _ = os.Hostname()
	m1.client_ip = r.RemoteAddr
	m1.node_name = get_env_variables()
	m1.process_time = time.Since(start)
	fmt.Fprintf(w, "> %s\n  you've hit %s\n  your ip: %s\n  processing time: %s\n", m1.time, m1.hostname, m1.client_ip, m1.process_time)
	fmt.Println(get_env_variables())
}

func main() {
	http.HandleFunc("/", handler1)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

