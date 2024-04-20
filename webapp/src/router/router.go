package router

import (
	"webapp/src/router/rotas"

	"github.com/gorilla/mux"
)

//Gerar returna um router com todas as rotas confguradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
