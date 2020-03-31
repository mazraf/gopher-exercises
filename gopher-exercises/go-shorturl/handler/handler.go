package handler

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page.")
}

type pathURL struct {
	Path string `yaml:path`
	Url  string `yaml:url`
}

func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// Intialize and parse yaml into pathURL.
	var pathURLs []pathURL
	err := yaml.Unmarshal(yamlBytes, &pathURLs)
	if err != nil {

		return nil, err
	}
	// Convert YAML to map
	pathToMap := make(map[string]string)
	for _, pu := range pathURLs {
		pathToMap[pu.Path] = pu.Url
	}
	return MapHandler(pathToMap, fallback), nil
}
