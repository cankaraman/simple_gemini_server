package main

import (
	"crypto/x509"
	"errors"
	"net/url"
	"os"
	"strings"
)

const Success string = "20"
const TemporaryFailure string = "40"
const PermanentFailure string = "50"

const NotFound string = "51"
const StatusBadRequest string = "59"
const DefaultHome string = "index.gmi"

// TODO implement  60 & 61 use cases

type Response struct {
	status string
	meta string
	body   *os.File
}

type Request struct {
	header string
	certs []*x509.Certificate
	
}

type GeminiUrlParser interface {
	GetRelativePath() (string, error)
}

//accept input
//
func (r Request) GetRelativePath() (string, error) {
	rawUrl := strings.Replace(r.header, "\r\n", "", -1)
	parsed, err := url.Parse(rawUrl)
	// TODO return error that correspons to a code
	if err != nil {
		return "", err
	}

	if parsed.Scheme != "" && parsed.Scheme != "gemini" {
		return "", errors.New("unsported scheme")
	}

	return strings.Trim(parsed.Path, "/"), nil
}

func NewRequest(header string, 	certs []*x509.Certificate) *Request {
	return &Request{header, certs}
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
