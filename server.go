package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
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

	config := &tls.Config{Certificates: []tls.Certificate{cer}, InsecureSkipVerify: true, ClientAuth: tls.RequestClientCert} // ClientAuth: tls.VerifyClientCertIfGiven,

	ln, err := tls.Listen("tcp", domain+":"+port, config)

	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()

	for {
		//var conn tls.Conn
		//var err error
		conn, err := ln.Accept()

		if err != nil {

			fmt.Println(err)
		}
		//tlsC tls.conn :=

		tlsConn:= conn.(*tls.Conn)
		tlsConn, ok := conn.(*tls.Conn)


		fmt.Println(ok)
		tlsState := tlsConn.ConnectionState()

		fmt.Println(tlsState)
		if tlsConn.ConnectionState().HandshakeComplete {
			fmt.Println(tlsConn.RemoteAddr().String())
		}

		//var tConn tls.Conn  = tls.Conn(conn)

		certs := tlsConn.ConnectionState().PeerCertificates
		fmt.Println(certs)

		//
		//conn.ConnectionState()
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
	// TODO request for certicication
	// TODO accept input

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
	rp, err := req.GetRelativePath()
	if err != nil {
		return NewResponse(StatusBadRequest, nil)
	}

	file, err := GetFile(rp)

	if err != nil {
		return NewResponse(NotFound, nil)
	}

	return NewResponse(Success, file)
}
