package main

var Routes map[string]func() *Response = map[string]func() *Response{
	"/other": Other,
}

func Other() *Response {
	f, err := GetFile("other.gmi")

	if err != nil {
		return NewResponse(NotFound, nil)
	}

	return NewResponse(Success, f)

}
