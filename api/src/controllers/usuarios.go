package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CriarUsuario(rw http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}
	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}
	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioID, erro := repositorio.Criar(usuario)
	if erro != nil {
		log.Fatal(erro)
	}
	rw.Write([]byte(fmt.Sprintf("Id inserido: %d", usuarioID)))
}

func BuscarUsuarios(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Buscando todos os usuários"))
}

func BuscarUsuario(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Buscando um usuário"))
}

func AtualizarUsuario(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Atualizando usuário"))
}

func DeletarUsuario(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Deletando usuário"))
}
