// Webcrawler build
// Built using the tutorial https://jdanger.com/build-a-web-crawler-in-go.html
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://triforcecode.com")

	fmt.Println("http transport error is:", err)

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println("read error is:", err)

	fmt.Println(string(body))
}
