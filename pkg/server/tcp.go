package server

import (
	"log"
	"net"
)

func tcpServe() {
	listen, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		log.Fatal("An error occured while opening port on 0.0.0.0:8081")
	}
	defer listen.Close()

	log.Println("Server is open on 0.0.0.0:8081")
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("An error occured while receiving connection")
		}

		go func(c net.Conn) {
			bytes := <-globalChannel
			conn.Write(bytes)
		}(conn)

		log.Println("waiting...")
	}

}
