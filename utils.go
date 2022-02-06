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
	body   *os.File
}

type Request struct {
	header string
}

type GeminiUrlParser interface {
	GetRelativePath() string
}

//clean url
//accept input
//getFlags
//
func (r Request) GetRelativePath() string {
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

func NewResponse(status string, body *os.File)*Response {
	return &Response{status, body}
}

func GetFile(rp string) (*os.File, error){
	//TODO return file pointer instead
	if rp == "" {
		return os.Open("./"+DefaultHome)
	}

	return os.Open("./"+rp+".gmi")
}
