package server

import (
	"log"
	"net"
	"net/http/httputil"
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
			request := <-globalChannel

			bytes, err := httputil.DumpRequest(&request, true)
			if err != nil {
				log.Fatal("Failed to dump request to bytes")
			}

			conn.Write(bytes)
		}(conn)

		log.Println("waiting...")
	}

}
