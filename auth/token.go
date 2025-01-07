package auth

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func GenerarToken(correo string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = correo
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return jwtToken.SignedString([]byte(os.Getenv("API_SECRET")))

}
