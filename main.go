package main

import (
	"fmt"
	"net/http"
)

func main() {
	sites := []string{
		"http://zomato.com",
		"http://codechef.com",
		"http://github.com",
		"http://flipkart.com",
	}

	checkLink(sites)
}

func checkLink(s []string) {
	for _, link := range s {
		_, err := http.Get(link)

		if err != nil {
			fmt.Println(link, "is down :(")
			return
		}

		fmt.Println(link, "is up :)")

	}
}
