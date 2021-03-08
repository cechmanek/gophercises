package main

import (
"fmt"
"net/http"
"urlshort"
//urlshort "../urlshort"//"github.com/pohercies/urlshort"

)

func main() {
	mux := defaultMux()

	// build the mapHandler using the mux as the fallback
	pathsToUrls := map[string]string {
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc": "https://godoc.org/gopkg.in/yaml.v2",
	} 
	mapHandler := urlshort.MaptHandler(pathsToUrls, mux)

	//built the YAMLHandler using the mapHandler as the fallback
	// back ticks are like triple quoted strings in python
	yaml := `
	- path: /urlshort
	url: https://github.com/gophercises.urlshort
	- path: /urlshort-final
	url: https://github.com/gophercises/urlshort/tree/final
	`
	yamlHanlder, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting teh server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
} // end main

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello, world!")
}