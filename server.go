package main

import (
	"bufio"
	"crypto/tls"
	"io"
	"log"
	"net"
)

func RunServer(domain, port, crt, key string) {
	log.SetFlags(log.Lshortfile)
	//use example certs
	cer, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		log.Println(err)
		return
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	ln, err := tls.Listen("tcp", domain+":"+port, config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	r := bufio.NewReader(conn)
	msg, err := r.ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}

	req := NewRequest(msg)

	handleResponse(req, conn)

	log.Println(msg)
}

func handleResponse(req *Request, conn net.Conn) {

	defer conn.Close()
	res := getResponse(req)

	_, err := conn.Write([]byte(res.status + "\r\n"))
	if err != nil {
		log.Println(err)
		return
	}
	io.Copy(conn, res.body)

}

func getResponse(req *Request) *Response {
	rp := req.GetRelativePath()

	file, err := GetFile(rp)

	if err != nil {
		return NewResponse(NotFound, nil)
	}

	return NewResponse(Success, file)
}
