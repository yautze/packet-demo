package get

import (
	go_resty "packet-demo/01.resty"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-resty/resty/v2"
)

//
const (
	// Header -
	ContentType = "Content-Type"

	// JSON -
	JSON = "application/json"
)

// GetOne -
func GetOne() {
	// new a Resty client
	client := resty.New()

	// API router
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// new request
	resp := client.R()

	// set header
	resp.SetHeader(ContentType, JSON)

	// set response struct
	response := new(go_resty.TestBody)
	resp.SetResult(&response)

	// post
	_, err := resp.Get(url)
	if err != nil {
		spew.Dump(err)
	}

	spew.Dump(response)
}

// GetAll -
func GetAll() {
	// new a Resty client
	client := resty.New()

	// API router
	url := "https://jsonplaceholder.typicode.com/posts"

	// new request
	resp := client.R()

	// set header
	resp.SetHeader(ContentType, JSON)

	// set response struct
	responses := make([]*go_resty.TestBody, 0)

	resp.SetResult(&responses)

	// post
	_, err := resp.Get(url)
	if err != nil {
		spew.Dump(err)
	}

	spew.Dump(responses)
}
