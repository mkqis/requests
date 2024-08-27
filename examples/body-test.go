package main

import (
	"fmt"
	"github.com/mkqis/requests"
	"github.com/mkqis/requests/url"
)

func main() {
	req := url.NewRequest()
	headers := url.NewHeaders()
	headers.Set("User-Agent", "123")
	req.Headers = headers
	req.Body = "testdata"
	r, _ := requests.Post("https://httpbin.org/post", req)
	fmt.Println(r.Text)
}
