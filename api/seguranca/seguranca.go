package seguranca

import "golang.org/x/crypto/bcrypt"

//Hash coverte uma senha em um hash dela
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerificarSenha compara uma senha e um hash e verifica se são iguais
func VerificarSenha(senhaComHash, senha string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senha))
}
