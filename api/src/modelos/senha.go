package modelos

// Senha representa o formato da requisição de auteração de senha
type Senha struct {
	Nova  string `json:"Nova"`
	Atual string `json:"Atual"`
}
