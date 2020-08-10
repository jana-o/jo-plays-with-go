package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	} //start new go routine and add c

	for l := range c {
		go func(link string) {
			time.Sleep(3 * time.Second)
			checkLink(link, c)
		}(l) //invoke func literal
	}
}

//checkLink checks links and adds c chan string as input
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "must be down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}
