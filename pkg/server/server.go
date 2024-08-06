package server


func NewServer() {
	go httpServe()
	tcpServe()
}

