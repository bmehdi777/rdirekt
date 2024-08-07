package server

import (
	"net/http"
)

var globalChannel chan http.Request

func NewServer() {
	globalChannel = make(chan http.Request)
	go httpServe()
	tcpServe()
}

