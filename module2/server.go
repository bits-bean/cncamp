package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", echoHeader)
	http.HandleFunc("/healthz", healthz)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func echoHeader(respWriter http.ResponseWriter, req *http.Request) {
	for key, value := range req.Header {
		respWriter.Header().Set(key, strings.Join(value, ","))
	}
	respWriter.Header().Set("VERSION", os.Getenv("VERSION"))
	code := http.StatusOK
	respWriter.WriteHeader(code)
	fmt.Printf("remoteAddr: %s, status code: %v\n", req.RemoteAddr, code)
}

func healthz(respWriter http.ResponseWriter, req *http.Request) {
	respWriter.WriteHeader(http.StatusOK)
}
