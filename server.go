package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/cankaraman/gemini_server/status"
)

func RunServer(domain, port, crt, key string) {
	log.SetFlags(log.Lshortfile)
	//use example certs
	cer, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		log.Println(err)
		return
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}, 
	InsecureSkipVerify: true, ClientAuth: tls.RequestClientCert} // ClientAuth: tls.VerifyClientCertIfGiven,

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
			break
		}

		tlsConn, ok := conn.(*tls.Conn)
		if !ok {
			break
		}

		tlsConn.Handshake()
		tlsState := tlsConn.ConnectionState()

		certs := tlsState.PeerCertificates
		fmt.Println(certs)

		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn, certs)
	}

}

func handleConnection(conn net.Conn, certs []*x509.Certificate) {
	// gemini request are at most 1024 bytes
	buf := make([]byte, 1024)

	_, err := conn.Read(buf)
	if err != nil {
		log.Println(err)
		handleBadRequest(conn)
		return
	}

	rawRequest := string(bytes.Trim(buf, "\x00"))
	log.Println(rawRequest)

	req, err := NewRequest(rawRequest, certs)
	if err != nil {
		log.Println(err)
		handleBadRequest(conn)
		return
	}

	handleResponse(req, conn)
}

func handleBadRequest(conn net.Conn) {

	res := NewResponse(status.BadRequest, nil)
	sendResponse(res, conn)
}

func handleResponse(req *Request, conn net.Conn) {
	defer conn.Close()
	res := getResponse(req)

	sendResponse(res, conn)
}

func sendResponse(res *Response, conn net.Conn) {
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
		return NewResponse(status.BadRequest, nil)
	}

	if Routes[rp] != nil {
		return Routes[rp](req)
	}

	file, err := GetFile(rp)

	if err != nil {
		return NewResponse(status.NotFound, nil)
	}

	return NewResponse(status.Success, file)
}
