package main

import (
	"errors"
	"fmt"
	"log"
	"net"
)

var mainPort = ":8080"

func main() {
	ln, err := NewTCPListener(mainPort)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = ln.Close()
		if err != nil {
			log.Default().Println(err)
		}
	}()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		err = handleTCPConn(conn)
		if err != nil {
			log.Default().Println(err)
			return
		}
	}
}

func NewTCPListener(port string) (ln net.Listener, err error) {
	ln, err = net.Listen("tcp", port)
	if err != nil {
		return nil, fmt.Errorf("listening to port %s : %v", port, err)
	}

	return ln, nil
}

func handleTCPConn(conn net.Conn) (err error) {
	defer func() {
		err = errors.Join(err, conn.Close())
	}()

	_, err = conn.Write([]byte("OK\n"))
	if err != nil {
		return fmt.Errorf("writing to conn: %v", err)
	}

	return err
}
