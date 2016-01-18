// Package dierks is a library for stubbing HTTP responses in your test suites with chainable methods. Perfect for testing API wrappers.
package dierks

import (
	"net/http"
	"net/http/httptest"
	"net/url"
)

// ResponseBuilder holds the body and content to be stubbed from an http client
type ResponseBuilder struct {
	statusCode   int
	contentType  string
	body         []byte
	headers      map[string]string
}

// Res builds a default ResponseBuilders (empty 200 JSON response)
func Res() ResponseBuilder {
	return ResponseBuilder{
		statusCode:   200,
		contentType:  "application/json",
		body:         []byte{},
		headers:      map[string]string{},
	}
}

// Status sets status code for a response
func (x ResponseBuilder) Status(code int) ResponseBuilder {
	x.statusCode = code
	return x
}

// Body sets a string body for a response
func (x ResponseBuilder) Body(val string) ResponseBuilder {
	return x.BodyData([]byte(val))
}

// BodyData sets a data body for a response
func (x ResponseBuilder) BodyData(val []byte) ResponseBuilder {
	x.body = val
	return x
}

// Header sets an arbitrary HTTP header for the response
func (x ResponseBuilder) Header(key, val string) ResponseBuilder {
	x.headers[key] = val
	return x
}

// ContentType sets the Content-Type header for the response
func (x ResponseBuilder) ContentType(val string) ResponseBuilder {
	x.contentType = val
	return x
}

// JSON sets the Content-Type header to application/json
func (x ResponseBuilder) JSON() ResponseBuilder {
	x.contentType = "application/json"
	return x
}

// XML sets the Content-Type header to application/xml
func (x ResponseBuilder) XML() ResponseBuilder {
	x.contentType = "application/xml"
	return x
}

// JSONAPI sets the Content-Type header to application/vnd.api+json
func (x ResponseBuilder) JSONAPI() ResponseBuilder {
	x.contentType = "application/vnd.api+json"
	return x
}

// Start returns a httptest.Server that responds accordingly, and an HTTPClient with an overidden transport which directly all request to the test server.
func (x ResponseBuilder) Start() (*httptest.Server, *http.Client) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", x.contentType)
		for k, v := range x.headers {
			w.Header().Set(k, v)
		}
		w.WriteHeader(x.statusCode)
		w.Write(x.body)
	}))

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}
	return server, &http.Client{Transport: transport}
}
