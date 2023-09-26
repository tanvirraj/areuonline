package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	links := []string{
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://stackoverflow.blog/",
		"https://go.dev/",
		"https://amazon.com",
		"https://tanvirrajHossain.com",
	}

	c := make(chan string)

	for _, link := range links {
		go CheckLink(link, c)

	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

	fmt.Print("\n\n")
	fmt.Println("Total response time", time.Since(start))
}

func CheckLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		c <- link + "site is down"
		return
	}

	c <- link + "site is up"

}
