package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

// Logger funçao intermediaria que escreve em log a função da rota sendo chamada
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(rw, r)
	}
}

// Autenticar verifica se o usuário fazendo a requsição está autenticado
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Autenticando...")
		next(rw, r)
	}
}
