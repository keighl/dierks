package dierks

import (
	"io/ioutil"
	"strings"
	"testing"
)

func Test_StatusCode(t *testing.T) {
	test := func(t *testing.T, builder ResponseBuilder) {
		server, client := builder.Start()
		defer server.Close()

		resp, err := client.Get("http://google.com")
		if err != nil {
			t.Errorf(err.Error())
		}

		if resp.StatusCode != 201 {
			t.Errorf("Expected 201 status, got: %d", resp.StatusCode)
		}
	}
	test(t, Res().Status(201))
}

////////////////////
////////////////////
////////////////////

func Test_ContentType(t *testing.T) {
	test := func(t *testing.T, builder ResponseBuilder, result string) {
		server, client := builder.Start()
		defer server.Close()

		resp, err := client.Get("http://google.com")
		if err != nil {
			t.Errorf(err.Error())
		}

		if resp.Header.Get("Content-Type") != result {
			t.Errorf("Expected %s, got: %s", result, resp.Header.Get("Content-Type"))
		}
	}
	test(t, Res().ContentType("application/xml"), "application/xml")
	test(t, Res().XML(), "application/xml")
	test(t, Res().JSON(), "application/json")
	test(t, Res().JSONAPI(), "application/vnd.api+json")
}

////////////////////
////////////////////
////////////////////

func Test_Body(t *testing.T) {
	content := `{"data":false}`
	test := func(t *testing.T, builder ResponseBuilder) {
		server, client := builder.Start()
		defer server.Close()

		resp, err := client.Get("http://google.com")
		if err != nil {
			t.Errorf(err.Error())
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf(err.Error())
		}
		defer resp.Body.Close()

		bodyString := strings.TrimSpace(string(body))
		if bodyString != content {
			t.Errorf("Expected body %s, got: %s", content, bodyString)
		}
	}

	test(t, Res().Body(content))
	test(t, Res().BodyData([]byte(content)))
}

////////////////////
////////////////////
////////////////////

func Test_Header(t *testing.T) {
	test := func(t *testing.T, builder ResponseBuilder) {
		server, client := builder.Start()
		defer server.Close()

		resp, err := client.Get("http://google.com")
		if err != nil {
			t.Errorf(err.Error())
		}

		if resp.Header.Get("Authorization") != "Bearer CHEESE" {
			t.Errorf("Expected auth header `Bearer CHEESE`, got: %s", resp.Header.Get("Authorization"))
		}
	}
	test(t, Res().Header("Authorization", "Bearer CHEESE"))
}
