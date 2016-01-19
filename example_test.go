package dierks_test

import (
	"github.com/keighl/dierks"
	"io/ioutil"
	"fmt"
)

func ExampleBody() {
	server, client := dierks.Res().Body(`{"data":false}`).Start()
	defer server.Close()

	resp, _ := client.Get("http://google.com")

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	// Output: {"data":false}
	// 200
	// application/json
}

func ExampleHeader() {
	server, client := dierks.Res().
		Header("Authorization", "Bearer XXXXX").
		Header("X-CLIENT-ID", "XXXXXXXXXX").
		Body(`...`).Start()
	defer server.Close()

	resp, _ := client.Get("http://google.com")

	fmt.Println(resp.Header.Get("Authorization"))
	fmt.Println(resp.Header.Get("X-CLIENT-ID"))
	// Output: Bearer XXXXX
	// XXXXXXXXXX
}

func ExampleContentType() {
	server, client := dierks.Res().Body(`<data>false</data>`).
		ContentType("application/xml").
		// or .XML() .JSON()
		Start()
	defer server.Close()

	resp, _ := client.Get("http://google.com")

	fmt.Println(resp.Header.Get("Content-Type"))
	// Output: application/xml
}

func ExampleStatus() {
	server, client := dierks.Res().Status(304).Start()
	defer server.Close()

	resp, _ := client.Get("http://google.com")

	fmt.Println(resp.StatusCode)
	// Output: 304
}
