package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/unidev-platform/golang-core/xenv"
)

func main() {

	// simple handler that prints hostname, requested path and env variables
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling request")
		hostname, err := os.Hostname()
		if err != nil {
			log.Printf("%v", err)
			return
		}
		fmt.Fprintf(w, "%s path: %v \n ", hostname, r.URL.Path)
		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			fmt.Fprintf(w, "%v = %v \n", pair[0], pair[1])
		}
	})

	listenAddress := xenv.Env("LISTEN_ADDRESS", "0.0.0.0")
	listenPort := xenv.Env("LISTEN_PORT", "8080")
	log.Printf("Starting HTTP server on %s:%s", listenAddress, listenPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", listenAddress, listenPort), nil))

}
