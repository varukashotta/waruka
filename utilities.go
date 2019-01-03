package utilities

import (
	"errors"
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

/*
 * init function added
 */
func init() {
	fmt.Println("rectangle package initialized")
}

//Today is another day in paradise
func Today(name string) {
	log.Println("Hello", name)
}

//DoRequest - send request and fetch response
func DoRequest(url string) (*fasthttp.Response, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	client.Do(req, resp)

	if resp.Header.StatusCode() != 200 {
		return nil, errors.New("Fetching Error")
	}

	return resp, nil
}
