package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server(ln net.Listener) { 
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("received", msg)
	}
	c.Close()
}

func client(msg string) {
	c, err := net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("sending", msg)

	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}

	c.Close()
}

func main() {

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	go server(ln)
	go client("yoo my n word")

	select{}
}
