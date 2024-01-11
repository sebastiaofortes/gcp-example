package main

import (
	"fmt"
	"net/http"
)

func main()  {

	fmt.Println("Iniciando app")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	fmt.Println("Iniciando servidor na porta 8080")

	http.ListenAndServe(":8080", nil)
}