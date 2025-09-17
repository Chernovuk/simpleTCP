package main

import (
	"fmt"
	"log"
	"net"
)

var url = "localhost:8080"

func main() {
	conn, err := DialTCP(url)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Default().Println(err)
		}
	}()

	err = ReadFromConn(conn)
	if err != nil {
		log.Default().Println(err)
		return
	}
}

func DialTCP(url string) (net.Conn, error) {
	conn, err := net.Dial("tcp", url)
	if err != nil {
		return nil, fmt.Errorf("dialing: %v", err)
	}

	return conn, nil
}

func ReadFromConn(conn net.Conn) error {
	buff := make([]byte, 1024)
	nBytes, err := conn.Read(buff)
	if err != nil {
		return fmt.Errorf("reading from conn: %v", err)
	}
	if checkResponse(string(buff[:nBytes])) {
		fmt.Println("Successful one-time connection")
	} else {
		fmt.Println("Wrong response!")
	}

	return nil
}

func checkResponse(resp string) bool {
	return resp == "OK\n"
}
