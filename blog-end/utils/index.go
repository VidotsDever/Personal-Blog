package utils

import (
	"github.com/dgrijalva/jwt-go"
	"strings"
)

func GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})
	return token.SignedString([]byte("i-love-you"))
}

func QueryStringToMap(query string) map[string]string {
	items := strings.Split(query, "&")
	objs := make(map[string]string)
	for i := 0; i < len(items); i++  {
		item := items[i]
		kws := strings.Split(item, "=")
		if len(kws) == 2 {
			objs[kws[0]] = kws[1]
		}
	}
	return objs
}
