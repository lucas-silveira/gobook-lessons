package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		method2(res, url)
	}
}

// func method1(res *http.Response, url string) {
// 	// res.Body is a stream
// 	// io.ReadAll load the entire data into memory
// 	b, err := io.ReadAll(res.Body)
// 	res.Body.Close() // Close the stream connection

// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
// 	}

// 	// b is an array of bytes (uint8)
// 	fmt.Println(reflect.TypeOf(b))

// 	fmt.Printf("%s", b)
// }

func method2(res *http.Response, url string) {
	// res.Body is a stream
	// io.Copy copy the data stream without a buffer
	_, err := io.Copy(os.Stdout, res.Body)
	res.Body.Close() // Close the stream connection

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
	}
}
