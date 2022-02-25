package main

import "flag"

func main() {

	domain := flag.String("domain", "127.0.0.1", "domain name or ip address")
	port := flag.String("port", "1965", "port for the server")
	crtPath := flag.String("crt", "./selfsigned.crt", "crt file path")
	keyPath := flag.String("key", "./selfsigned.key", "key file path")
	flag.Parse()

	RunServer(*domain, *port, *crtPath, *keyPath)


	// meta field on response
	// 1024 byte long request url. Read max until \r\n
	// success meta is just content type
}



