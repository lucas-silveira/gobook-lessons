package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	/*
		- Internamente, o servidor executa cada requisição
			em uma gorrotina separada.
		- Um padrão de handler que termine com uma barra
			corresponde a qualquer URL que tenha o padrão
			como prefixo.
	*/
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
