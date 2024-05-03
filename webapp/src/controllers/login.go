package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	token, _ := io.ReadAll(response.Body)
	fmt.Println(response.StatusCode, string(token))
	/*
	   defer response.Body.Close()

	   	if response.StatusCode > 400 {
	   		respostas.TratarStatusCodeDeErro(rw, response)
	   		return
	   	}

	   respostas.JSON(rw, response.StatusCode, nil)
	*/
}
