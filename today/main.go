package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if err := http.ListenAndServe(":8085", mainHandler()); err != nil {
		fmt.Fprintf(os.Stderr, "ListenAndServe failed: %v\n", err)
	}
}

func mainHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintf(w, "Hello!")
	})
	return mux
}
