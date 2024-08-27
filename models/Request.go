package models

import (
	"github.com/mkqis/chttp"
	"github.com/mkqis/chttp/cookiejar"
	"github.com/mkqis/requests/url"
)

type Request struct {
	Method  string
	Url     string
	Params  *url.Params
	Headers *http.Header
	Cookies *cookiejar.Jar
	Data    *url.Values
	Files   *url.Files
	Body    string
	Json    map[string]interface{}
	Auth    []string
}

func (req *Request) Prepare() *PrepareRequest {
	p := &PrepareRequest{}
	p.Prepare(
		req.Method,
		req.Url,
		req.Params,
		req.Headers,
		req.Cookies,
		req.Data,
		req.Files,
		req.Json,
		req.Body,
		req.Auth,
	)
	return p
}
