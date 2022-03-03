package main

import (
	"crypto/x509"
	"errors"
	"net/url"
	"os"
	"strings"
)

const DefaultHome string = "index.gmi"

type Response struct {
	status string
	meta   string
	body   *os.File
}

type Request struct {
	header, rawUrl string
	url            *url.URL
	certs          []*x509.Certificate
}

type GeminiUrlParser interface {
	GetRelativePath() (string, error)
}

//accept input
//
func (r Request) GetRelativePath() (string, error) {

	return strings.Trim(r.url.Path, "/"), nil
}

func NewRequest(header string, certs []*x509.Certificate) (*Request, error) {
	rawUrl := strings.Replace(header, "\r\n", "", -1)
	parsed, err := url.Parse(rawUrl)
	// 2. TODO return error that correspons to a code
	if err != nil {
		return nil, err
	}

	if parsed.Scheme != "" && parsed.Scheme != "gemini" {
		return nil, errors.New("unsported scheme")
	}
	return &Request{header, rawUrl, parsed, certs}, nil
}

func NewResponse(status string, body *os.File) *Response {
	return &Response{status, "", body}
}

func GetFile(rp string) (*os.File, error) {
	if rp == "" {
		return os.Open("./" + DefaultHome)
	}

	return os.Open("./" + rp + ".gmi")
}
