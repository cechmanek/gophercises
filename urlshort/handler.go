package urlshort

import (
	"net/http"
	yaml "gopkg.in/yaml.v2"
)

// MapHandler returns a http.HandlerFunc (which also implements http.Handler)
// that wil attempt to map any paths (keys in the map) to thier corresponding
// URL (values that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallbak http.Handler will be
// called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		// if we can match a path... run the ok branch, which returns
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		// else..
	 	fallback.ServeHTTP(w,r)
	}
}

// YAMLHandler will parse the provided YAML and then return a http.HandlerFun
// (which also implements httpHandler) that will attempt to map any paths to
// their correspdning URL. If the path is not provide in the YAML, then the
// fallback http.Handler wil be called instead.
func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// first, parse yaml
	var pathURLs []pathURL
	err := yaml.Unmarshal(yamlBytes, &pathURLs)
	if err != nil {
		return nil, err
	}

	// second conver yaml array into map
	pathsToUrls := make(map[string]string)
	for _, pu := range pathURLs {
		pathsToUrls[pu.Path] = pu.URL
	}

	// third return map handler using the map
	return MapHandler(pathsToUrls, fallback), nil
}

type pathURL struct {
	Path string `yaml:"path"`
	URL string `yaml:"url"`
}