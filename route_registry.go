package main

import (
	"github.com/cankaraman/gemini_server/status"
)

var Routes map[string]func(*Request) *Response = map[string]func(*Request) *Response{
	"other": Other,
	"other/input": OtherInput,
}

func Other(req *Request) *Response {
	f, err := GetFile("other.gmi")

	if err != nil {
		return NewResponse(status.NotFound, nil)
	}

	return NewResponse(status.Success, f)

}

func OtherInput(req *Request) *Response {
	return NewResponse(status.Input, nil)
}