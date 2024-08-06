package server

import (
	"log"
	"net/http"
	"net/http/httputil"
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
	bytes, err := httputil.DumpRequest(request, true)
	if err != nil {
		writter.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	log.Println("before")
	globalChannel <- bytes
	log.Println("after")

	writter.WriteHeader(http.StatusOK)
	writter.Write([]byte("Endpoint reached"))
}
