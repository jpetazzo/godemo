package main

import (
	"fmt"
	"net/http"
	"os"
  "strings"
)

func serve(w http.ResponseWriter, r *http.Request) {
	ua := r.UserAgent()
	fmt.Printf("Got a request from %s\n", ua)
	if strings.HasPrefix(ua, "curl/") {
    w.Write([]byte("This request was served by "))
    w.Write([]byte(os.Getenv("HOSTNAME")))
    w.Write([]byte("\n"))
	} else {
    http.Redirect(w, r, "https://www.youtube.com/watch?v=dQw4w9WgXcQ&t=43", 302)
	}
}

func main() {
	fmt.Printf("HTTP server is up and running.\n")
	http.HandleFunc("/", serve)
	listen := os.Getenv("LISTEN")
	if listen == "" {
		listen = ":8888"
	}
	if err := http.ListenAndServe(listen, nil); err != nil {
		panic(err)
	}
}
