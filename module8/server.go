package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	http.HandleFunc("/", echoHeader)
	http.HandleFunc("/healthz", healthz)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)
	go func() {
		fmt.Println(<-sigs)
	}()
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
