// Concurrent website status checker. 
// Checks all links until each link is accessed sucessfully 10 times.

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	go func() {
		for x := 0; x < 10; x++ {
			for _, link := range links {
				checkLink(link, c)
			}
		}
		close(c)
	}()
	
	for l := range c {
		go func() {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}()
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // <- Blocking call
	
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	
	fmt.Println(link, "is up!")
}
