package main

import (
    "log"
    "crypto/tls"
)

func main() {
    log.SetFlags(log.Lshortfile)


	cer, err := tls.LoadX509KeyPair("../selfsigned.crt", "../selfsigned.key")
	if err != nil {
		log.Println(err)
		return
	}

	conf := &tls.Config{InsecureSkipVerify: true, Certificates: []tls.Certificate{cer}} // ClientAuth: tls.VerifyClientCertIfGiven,

    // conf := &tls.Config{
    //      //InsecureSkipVerify: true,
		 
    // }

    conn, err := tls.Dial("tcp", "127.0.0.1:1965", conf)
    if err != nil {
        log.Println(err)
        return
    }
    defer conn.Close()

    certs := conn.ConnectionState().PeerCertificates
    print(certs)

    n, err := conn.Write([]byte("hello\n"))
    if err != nil {
        log.Println(n, err)
        return
    }

    buf := make([]byte, 100)
    n, err = conn.Read(buf)
    if err != nil {
        log.Println(n, err)
        return
    }

    println(string(buf[:n]))
}