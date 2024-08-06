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
	buffer := make([]byte, 1024*1024);
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("An error occured while receiving connection")
		}

		if length, err := conn.Read(buffer); err != nil {
			log.Fatal("Can't read from client")
		} else {
			log.Println("Data received : ", string(buffer))
			log.Println("Length received : ", length)

			conn.Write([]byte("Hello there!"))
		}

		log.Println("waiting...")
	}

}
