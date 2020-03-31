package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/farzamalam/gopher-exercises/go-shorturl/handler"
)

func main() {
	mux := defaultMux()

	pathsToUrl := map[string]string{
		"/github":   "https://github.com/farzamalam",
		"/linkedin": "https://www.linkedin.com/in/farzamalam/",
	}
	maphandler := handler.MapHandler(pathsToUrl, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandler, err := handler.YAMLHandler([]byte(yaml), maphandler)
	if err != nil {
		fmt.Println("Error in YAMLHandler : ", err)
	}
	fmt.Println("Starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", yamlHandler))
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.DefaultHandler)
	return mux
}
