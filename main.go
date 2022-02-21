package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/unidev-platform/golang-core/xenv"
)

func main() {

	// simple handler that prints hostname, requested path and env variables
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		sleep := xenv.Env("PROCESSING_SLEEP", "0")
		isleep, err := strconv.Atoi(sleep)
		if err != nil {
			log.Printf("%v", err)
			return
		}
		time.Sleep(time.Duration(isleep) * time.Millisecond)

		log.Println("Handling request")
		code := xenv.Env("HTTP_RESPONSE_CODE", "200")
		icode, err := strconv.Atoi(code)
		if err != nil {
			log.Printf("%v", err)
			return
		}
		w.WriteHeader(icode)
		hostname, err := os.Hostname()
		if err != nil {
			log.Printf("%v", err)
			return
		}
		name := xenv.Env("APP_NAME", "TestApp")
		fmt.Fprintf(w, "%s %s path: %v \n ", name, hostname, r.URL.Path)
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
