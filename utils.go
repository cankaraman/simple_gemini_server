package main

import (
	"os"
	"strings"
)

const Success string = "20"
const TemporaryFailure string = "40"
const PermanentFailure string = "50"
const NotFound string = "51"
const DefaultHome string = "index.gmi"

type Response struct {
	status string
	body   []byte
}

type Request struct {
	header string
}

type GeminiUrlParser interface {
	getRelativePath() string
}

func (r Request) getRelativePath() string {
	url := strings.Replace(r.header, "\r\n", "", -1)
	url = strings.Replace(url, "gemini://", "", -1)
	url = strings.Replace(url, "/../", "/", -1)
	if strings.Contains(url, "/") {
		return strings.Trim(url[strings.IndexByte(url, '/'):], "/")
	}
	return ""
}

func NewRequest(header string) *Request {
	return &Request{header}
}

func NewResponse(status string, body []byte) *Response {
	return &Response{status, body}
}

func GetFile(rp string) ([]byte, error){
	//TODO return file pointer instead
	if rp == "" {
		return os.ReadFile("./"+DefaultHome)
	}

	return os.ReadFile("./"+rp+".gmi")
}
