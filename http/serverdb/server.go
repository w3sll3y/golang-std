package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/usuarios/", UsuarioHandler)
	log.Println("Excecutando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
