package main

import (
	"log"

	"github.com/cankaraman/gemini_server/status"
)

var Routes map[string]func(*Request) *Response = map[string]func(*Request) *Response{
	"other":       Other,
	"other/input": OtherInput,
}

//TODO architecture improvements
// usable server struct with config fields
// easier way to register routes like in django or node.js
//1. TODO redirect after input 

func Other(req *Request) *Response {

	if req.certs == nil || len(req.certs) == 0 {
		return NewResponse(status.ClientCertificateRequired, nil)
	}

	fp, err := req.GetRelativePath()

	if err != nil {
		return NewResponse(status.NotFound, nil)
	}

	f, err := GetFile(fp)

	if err != nil {
		return NewResponse(status.NotFound, nil)
	}

	return NewResponse(status.Success, f)

}

func OtherInput(req *Request) *Response {
	log.Println(req.url.RawQuery)
	return NewResponse(status.Input, nil)
}
