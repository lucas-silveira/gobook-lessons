package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	// Chamadas concorrentes
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // inicia uma gorrotina
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // recebe do canal ch
	}

	// O tempo total gasto será equivalente ao
	// tempo da request mais demorada
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	res, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) // envia para o canal ch
		return
	}

	nbytes, err := io.Copy(io.Discard, res.Body)
	res.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
