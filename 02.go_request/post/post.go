package post

import (
	"fmt"

	go_request "packet-demo/02.go_request"

	"github.com/parnurzeal/gorequest"
	"github.com/sirupsen/logrus"
)

// Post - gorequest post test
func Post() {
	// gorequest init
	req := gorequest.New()

	// API router
	url := "https://jsonplaceholder.typicode.com/posts"

	// Post
	req.Post(url)

	// Post data
	d := &go_request.TestBody{
		Title:  "YauTz",
		Body:   "YauTz Body",
		UserID: 123456,
	}

	// send
	req.Send(d)

	// response
	v := new(go_request.TestBody)
	resp, _, errs := req.EndStruct(&v)

	// error check
	for _, err := range errs {
		if err != nil {
			logrus.Error(err)
			break
		}
	}

	fmt.Println(resp)
	fmt.Println()
	fmt.Println("res: ", v)
}
