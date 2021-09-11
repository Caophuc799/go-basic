package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://facebook.com",
		"http://golang.org",
		"http://tiki.com",
		"http://google.com",
	}

	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

	for _, link := range links {
		go repeatCheckLink(link, c)
	}
	// infinity loop
	for l := range c {
		// function literals
		go func(l string) {
			time.Sleep(5 * time.Second)
			go repeatCheckLink(l, c)
		}(l)
	}
	fmt.Println("End!")
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " might be down!")
		c <- "Might be down I think"
		return
	}
	c <- "It's up"
	fmt.Println(link + " is up!")
}

func repeatCheckLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " might be down!")
		c <- link
		return
	}
	c <- link
	fmt.Println(link + " is up!")
}
