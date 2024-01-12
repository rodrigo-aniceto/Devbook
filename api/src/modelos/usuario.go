package modelos

import "time"

//usuario representa um usuario da rede
type Usuario struct {
	ID       uint64    `json:"id,omitempty"` //o json omite esse campo se estiver vazio
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}
