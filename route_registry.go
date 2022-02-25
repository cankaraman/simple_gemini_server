package main

var routes map[string]func() *Response = map[string]func() *Response{
	"/": Home,
}

func Home() *Response{
	return NewResponse(Success, nil)

}
