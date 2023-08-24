package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

/*
Um mutex é um bloqueio de exclusão mútua.
Apenas uma thread pode segurar o bloqueio.
Os mutexes são usados para proteger dados
compartilhados do acesso simultâneo.
*/
var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

/*
	Para evitar race-conditions ao tentar atualizar
	a variável global count, utilizamos um mutex
*/

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()

	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
