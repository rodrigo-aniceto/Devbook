package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// CarregarTelaDeLogin carrega a tela de login
func CarregarTelaDeLogin(rw http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(rw, "login.html", nil)
}

// CarregarPaginaDeCadastroDeUsuario carrega a pagina de cadastro de usuarios
func CarregarPaginaDeCadastroDeUsuario(rw http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(rw, "cadastro.html", nil)
}
