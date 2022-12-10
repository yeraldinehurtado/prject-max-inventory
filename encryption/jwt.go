package encryption

import (
	"inventory/internal/models"

	"github.com/golang-jwt/jwt/v4"
)

func SignedLoginToken(u *models.User) (string, error) {

	// los claims en pocas palabras son los parametros o variables que queremos compartir con un cliente

	// el metodo de HS256 es viable si solamente el servidor que creó el token es quien se va a encargar de validar
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": u.Email,
		"name": u.Name,
	}) 
/* 
	// signedString lo que va a hacer es crear el jwtoken como string y lo va a firmar
	jwtString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	} */

	return token.SignedString([]byte(key))

}
