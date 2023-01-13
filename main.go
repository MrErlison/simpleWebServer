package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	portFlag string
	dirFlag  string
)

// Initialize flags with default values
func init() {
	flag.StringVar(&portFlag, "p", "8000", "bind to this port")
	flag.StringVar(&dirFlag, "d", "./", "serve this directory")

}

// logHandler is a middleware function that logs information
// about the incoming request
func logHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func main() {
	flag.Parse()
	fmt.Printf("Serving HTTP on http://127.0.0.1:%s\n", portFlag)
	http.ListenAndServe(":"+portFlag, logHandler(http.FileServer(http.Dir(dirFlag))))
}
