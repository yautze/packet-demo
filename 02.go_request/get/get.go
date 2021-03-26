package get

import (
	"fmt"

	go_request "packet-demo/02.go_request"

	"github.com/parnurzeal/gorequest"
	"github.com/sirupsen/logrus"
)

// GetOne - gorequest get one test
func GetOne() {
	// gorequest init
	req := gorequest.New()

	// API router
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// GET
	req.Get(url)

	// request set
	req.Header.Add("Content-Type", "application/json")
	req.Type("json")

	// Unmarshal body
	v := new(go_request.TestBody)
	_, _, errs := req.EndStruct(v)

	// error check
	for _, err := range errs {
		if err != nil {
			logrus.Error(err)
			break
		}
	}

	fmt.Println(v)
}

// GetAll - gorequest get lsit test
func GetAll() {
	// gorequest init
	req := gorequest.New()

	// API router
	url := "https://jsonplaceholder.typicode.com/posts"

	// GET
	req.Get(url)

	// request set
	req.Header.Add("Content-Type", "application/json")
	req.Type("json")

	// Unmarshal body
	v := make([]*go_request.TestBody, 0)
	_, _, errs := req.EndStruct(&v)

	// error check
	for _, err := range errs {
		if err != nil {
			logrus.Error(err)
			break
		}
	}

	for _, val := range v {
		fmt.Println(val)
		fmt.Println()
	}
}
