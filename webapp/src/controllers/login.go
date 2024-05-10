package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"webapp/src/modelos"
	"webapp/src/respostas"
)

// FazerLogin utiliza o email e senha do usuário para autenticar na aplicação
func FazerLogin(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	usuario, erro := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		respostas.JSON(rw, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	response, erro := http.Post("http://localhost:5000/login", "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(rw, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(rw, response)
		return
	}
	var dadosAutenticacao modelos.DadosAutenticacao
	if erro = json.NewDecoder(response.Body).Decode((&dadosAutenticacao)); erro != nil {
		respostas.JSON(rw, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

}
