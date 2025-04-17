package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func checkAndSaveBody(url string, c chan string) {

	resp, err := http.Get(url)
	if err != nil {
		s :=	fmt.Sprintf("%s is DOWN!\n", url)
		s += fmt.Sprintf("Error: %v\n",err)
		c <- s // sending into the channel
	} else {
		s := fmt.Sprintf("%s -> Status code: %d \n", url, resp.StatusCode)
		s += fmt.Sprintf("%s is UP\n",url)
		c <- s
	}

}

func main() {

	urls := []string{"https://www.golang.org", "https://www.google.com", "https://www.medium.com"}

	c := make(chan string)

	for _, url := range urls {
		// 4.
		// Launch the goroutines
		go checkAndSaveBody(url, c)
	}
	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())

	// for i := 0; i < len(urls); i++{
	// 	fmt.Println(<-c)
	// }

	// for {
	// 	go checkAndSaveBody(<-c,c)
	// 	fmt.Println(strings.Repeat("#",30))
	// 	time.Sleep(time.Second * 10)
	// }


	for url := range c {
		go func(u string){
			time.Sleep(time.Second * 2)
			checkAndSaveBody(u,c)
		}(url)

	}

}

