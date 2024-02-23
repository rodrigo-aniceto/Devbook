package autenticacao

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//CriarToken com as permissoes de login
func CriarToken(usuarioId uint64) (string, error) {
	permissoes := jwt.MapClaims{}

	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix() //expiração em 6 horas
	permissoes["usuarioId"] = usuarioId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte("abc")) // secret
}
