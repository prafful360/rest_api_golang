package token

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	Models "github.com/rest_api/Models"
	bcrypt "golang.org/x/crypto/bcrypt"
)

var jwt_key = []byte("my_secret_key")

func HashPassword(pass_str string) ([]byte, string) {

	hash, err := bcrypt.GenerateFromPassword([]byte(pass_str), bcrypt.DefaultCost)

	if err != nil {
		return nil, err.Error()
	}

	return hash, ""
}

func ValidatePassword(pass_str string, pass_hashed []byte) bool {

	if err := bcrypt.CompareHashAndPassword(pass_hashed, []byte(pass_str)); err != nil {
		return false
	}

	return true
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken() string {
	var user Models.User

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwt_key)
	if err != nil {
		return err.Error()
	}

	return tokenString

}

func ValidateToken(token string) (bool, error) {

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwt_key, nil
	})

	if err != nil {
		return false, err
	}

	if !tkn.Valid {
		return false, errors.New("Not Authorized")
	}

	return true, nil
}
