package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Error    error
	Response *http.Response
}

// If you goroutine can provides errors, those errors should be tightly coupled with your result type.
func main() {
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)

			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{err, resp}

				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()

		return results
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"http://baidu.com", "http://x.x"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("Error:%v\n", result.Error)
			continue
		}
		fmt.Printf("Response :%v\n", result.Response.Status)
	}

	fmt.Println("Done")
}
