package main

import (
	"log"
	"net"
	"time"
)

func main() {

	addr, err := net.ResolveTCPAddr("tcp", ":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalf("listen error:%s", err.Error())
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("accept error:%s", err.Error())
			continue
		}

		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) {
	log.Printf("new connection: %s, %s", conn.LocalAddr(), conn.RemoteAddr())
	defer conn.Close()

	tcpConn, ok := conn.(*net.TCPConn)
	if ok {
		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(time.Second)
	}

	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}

		_, err = conn.Write(buf[0:n])
		if err != nil {
			return
		}
	}
}
