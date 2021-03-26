package post

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

// Post -
func Post() {
	// new a Resty client
	client := resty.New()

	// API router
	url := "https://jsonplaceholder.typicode.com/posts"

	// new request
	resp := client.R()

	// set header
	resp.SetHeader(ContentType, JSON)

	// Post data
	d := &go_resty.TestBody{
		Title:  "YauTz",
		Body:   "YauTz Body",
		UserID: 123456,
	}

	// set post body
	resp.EnableTrace().SetBody(d)

	// set response struct
	response := new(go_resty.TestBody)
	resp.SetResult(&response)

	// post
	_, err := resp.Post(url)
	if err != nil {
		spew.Dump(err)
	}

	spew.Dump(response)
}
