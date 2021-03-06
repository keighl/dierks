# Dierks

[![Build Status](https://travis-ci.org/keighl/dierks.png?branch=master)](https://travis-ci.org/keighl/dierks) [![codecov.io](https://codecov.io/github/keighl/dierks/coverage.svg?branch=master)](https://codecov.io/github/keighl/dierks?branch=master) [![GoDoc](https://godoc.org/github.com/keighl/dierks?status.svg)](https://godoc.org/github.com/keighl/dierks)

Dierks is a library for stubbing HTTP responses in Golang tests using chained methods. Perfect for testing API wrappers!

### Installation

    go get -u github.com/keighl/dierks

### Basic Usage

Dierks stubs HTTP responses for requests made through an `http.Client`. Here's a simple example showing how to stub the response body for a request to google.com.

```go
// example_test.go

package dierks

import (
    "io/ioutil"
    "strings"
    "testing"
)

func TestGoogle(t *testing.T) {

    responsePayload := `{"user": {"id": 3, "name": "Kyle"}}`

    // Dierks generates a test server, and an http.Client.
    // Any request through the client will generate a response with
    // `responsePayload` in the body
    server, client := dierks.Res().Body(responsePayload).Start()
    defer server.Close()

    // Make a request through the client
    resp, _ := client.Get("http://google.com")

    // Let's look at what the response has...
    body, _ := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()

    // What?! Google returned my JSON! Cool!
    expect(t, string(body), responsePayload)
}
```
More realistically, you might use dierks to test a API wrapper-lib method.

```go
func TestGetUser(t *testing.T) {
    responsePayload := `{"user": {"id": 3, "name": "Kyle"}}`

    server, client := dierks.Res().Body(responsePayload).Start()
    defer server.Close()

    apiClient := &APIClient{CustomHTTPClient: client}

    user, err := apiClient.GetUser(3)

    expect(t, err, nil)
    expect(t, user.ID, 3)
    expect(t, user.Name, "Kyle")
}
```

### Building a Response

###

