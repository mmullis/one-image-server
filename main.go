package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type", "text/html")
	data, err := Asset("resources/onlyimage.jpg")
	if err != nil {
		data = []byte("image missing") // asset not found
	}

	switch r.Method {
	case "GET":
		w.Write(data)
	default:
		fmt.Fprintf(w, "Unsupported!")
	}
}

func main() {
	port := flag.String("p", "9666", "port to serve on")
	flag.Parse()
	http.HandleFunc("/", indexHandler)
	log.Printf("Serving on HTTP port: %s\n", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
