package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()

	r := router.Gerar()

	fmt.Printf("Rodando a API na porta %d\n", config.Porta)

	str_porta := fmt.Sprintf(":%d", config.Porta)
	log.Fatal(http.ListenAndServe(str_porta, r))
}
