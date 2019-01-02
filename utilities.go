package utilities

import (
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
func DoRequest(url string) string {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	client.Do(req, resp)

	bodyBytes := resp.Body()

	return string(bodyBytes)
	// User-Agent: fasthttp
	// Body:
}
