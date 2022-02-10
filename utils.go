package main

import (
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

type Response struct {
	status string
	body   *os.File
}

type Request struct {
	header string
}

type GeminiUrlParser interface {
	GetRelativePath() (string, error)
}

//accept input
//
func (r Request) GetRelativePath() (string, error) {
	rawUrl := strings.Replace(r.header, "\r\n", "", -1)
	parsed, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	if parsed.Scheme != "" && parsed.Scheme != "gemini" {
		return "", errors.New("unsoperted scheme")
	}

	return strings.Trim(parsed.Path, "/"), nil
}

func NewRequest(header string) *Request {
	return &Request{header}
}

func NewResponse(status string, body *os.File) *Response {
	return &Response{status, body}
}

func GetFile(rp string) (*os.File, error) {
	if rp == "" {
		return os.Open("./" + DefaultHome)
	}

	return os.Open("./" + rp + ".gmi")
}
