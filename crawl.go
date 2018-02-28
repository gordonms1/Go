// Webcrawler build
// Built using the tutorial https://jdanger.com/build-a-web-crawler-in-go.html
package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/jackdanger/collectlinks"
)

var visited = make(map[string]bool)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Please specify the start page")
		os.Exit(1)
	}

	queue := make(chan string)

	go func ()  {

		queue <- args[0]

	}()

	for uri := range queue {

		enqueue(uri, queue)
	}
}

func enqueue(uri string, queue chan string)  {
	fmt.Println("fetching", uri)
	visited[uri] = true
	transport := &http.Transport{
		TLSClientConfig: tlsConfig{
		InsecureSkipVerify: true
	},
}
client := http.Client{Transport: transport}
resp, err := http.Get(uri)
if err != nil {
	return
}
	defer resp.Body.Close()

	links := collectlinks.All(resp.Body)

	for _, link := range links {
		absolute := fixUrl(link, uri)
		if uri != "" {
			if !visited[absolute] {
				go func() { queue <- absolute } ()
			}
		}
	}
}

func fixUrl(href, base string) (string)  {
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uri = baseUrl.ResolveReference(uri)
	return uri.String()
}

// And now you can start at any html page you like and slowly explore the entire world wide
// web.
//
// go run crawl.go http://anywebsite.com
