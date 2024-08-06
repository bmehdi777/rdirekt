package server

import (
	"net/http"
	"log"
)

func httpServe() {
	log.Println("Http server is up on http://localhost:8080")
	http.HandleFunc("/healthcheck", healthCheck)
	http.HandleFunc("/", redirectTcp)

	http.ListenAndServe(":8080",nil)
} 

func healthCheck(writter http.ResponseWriter, request *http.Request) {
	log.Println("Health check called")
	writter.WriteHeader(http.StatusOK)
}

func redirectTcp(writter http.ResponseWriter, request *http.Request) {
	writter.WriteHeader(http.StatusOK)
	writter.Write([]byte("Endpoint reached"))
}
