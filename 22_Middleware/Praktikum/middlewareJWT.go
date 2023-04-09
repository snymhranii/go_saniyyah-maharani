package middlewares

import {
	"project_structure/constants"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
}

func CreateToken(userId int)  {string, error} {
	claims := jwt.NapClaims{}
	claims["authorized"] = true
	claims["userId"] = true
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	byteSecret := []byte(constants.SECRET_JWT)
	return token.SignedString(byteSecret)
	
}

func ExtractTokenUserId(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].{int}
}

return 0

}
