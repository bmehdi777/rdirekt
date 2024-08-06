package server

var globalChannel chan []byte

func NewServer() {
	globalChannel = make(chan []byte)
	go httpServe()
	tcpServe()
}

