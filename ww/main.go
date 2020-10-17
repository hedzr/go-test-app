package main

import (
	"fmt"
	"github.com/hedzr/go-test-app/ww/handlers"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", &indexHandler{content: "hello world!"})
	http.HandleFunc("/health-check", handlers.HealthCheckHandler)
	log.Fatal(http.ListenAndServe(":8111", nil))
}

type indexHandler struct {
	content string
}

func (ih *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, ih.content)
}
