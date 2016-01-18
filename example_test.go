package dierks

import (
	"io/ioutil"
	"testing"
	"fmt"
)

func Example_Basic(t *testing.T) {
	server, client := Res().Body(`{"data":false}`).Start()
	defer server.Close()

	resp, _ := client.Get("http://google.com")

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body)) // {"data":false}
	fmt.Println(resp.StatusCode) // 200
	fmt.Println(resp.Header.Get("ContentType")) // application/json
}

func Example_Headers(t *testing.T) {
	server, client := Res().
		Header("Authorization", "Bearer XXXXX").
		Header("X-CLIENT-ID", "XXXXXXXXXX").
		Body(`...`).Start()
	defer server.Close()

	resp, _ := client.Get("http://google.com")

	fmt.Println(resp.Header.Get("Authorization")) // Bearer XXXXX
	fmt.Println(resp.Header.Get("X-CLIENT-ID")) // XXXXXXXXXX
	fmt.Println(resp.StatusCode) // 200
}

func Example_ContentType(t *testing.T) {
	server, client := Res().Body(`<data>false</data>`).
		ContentType("application/xml").
		// or .XML() .JSON()
		Start()
	defer server.Close()

	resp, _ := client.Get("http://google.com")

	fmt.Println(resp.Header.Get("ContentType")) // application/xml
}

func Example_Status(t *testing.T) {
	server, client := Res().Status(301).Start()
	defer server.Close()

	resp, _ := client.Get("http://google.com")

	fmt.Println(resp.StatusCode) // 301
}
