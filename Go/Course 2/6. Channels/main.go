package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	c := make(chan string)

	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://youtube.com",
		"http://codemagic.io",
		"http://github.com",
		"http://slack.com",
		"http://orkut.com",
		"http://behance.com",
	}

	for _, url := range urls {
		go printSiteStatus(url, c)
	}

	for url := range c {
		go func(url string) {
			time.Sleep(time.Second * 5)
			go printSiteStatus(url, c)
		}(url)

	}

}

func printSiteStatus(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println(url, "might be DOWN")
		c <- url
		return
	}

	fmt.Println(url, "is UP")
	c <- url
}
