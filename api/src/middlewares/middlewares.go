package middlewares

import (
	"api/src/autenticacao"
	"api/src/respostas"
	"log"
	"net/http"
)

// Logger funçao intermediaria que escreve em log a função da rota sendo chamada
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(rw, r)
	}
}

// Autenticar verifica se o usuário fazendo a requsição está autenticado
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil {
			respostas.Erro(rw, http.StatusUnauthorized, erro)
			return
		}
		proximaFuncao(rw, r)
	}
}
