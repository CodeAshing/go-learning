package urlshort

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

func NewBaseUrlMapper(urls map[string]string) func(string) (string, bool) {
	return func(path string) (string, bool) {
		url, ok := urls[path]
		return url, ok
	}
}

func NewYamlUrlMapper(filename string) (func(string) (string, bool), error) {
	content, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	var yml []map[string]string
	err = yaml.Unmarshal(content, &yml)

	mapping := make(map[string]string)

	if err != nil {
		return nil, err
	}

	for _, m := range yml {
		mapping[m["path"]] = m["url"]
	}

	return NewBaseUrlMapper(mapping), nil
}

func NewHttpRedirectHandler(mapper func(string) (string, bool), fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if url, ok := mapper(r.URL.Path); ok {
			log.Printf("Redirecting %s to %s\n", r.URL.Path, url)
			http.Redirect(w, r, url, http.StatusMovedPermanently)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}
