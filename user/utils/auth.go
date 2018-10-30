package utils

import (
	"io/ioutil"
	 "github.com/dgrijalva/jwt-go"
	 "strings"
)



func VerifyToken(token string) (string, error) {
	publicKeyPath := "./public.cert"

    keyData, err := ioutil.ReadFile(publicKeyPath)
    if err != nil {
        return "false", err
    }
    key, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
    if err != nil {
        return "false", err
    }
	claims := jwt.MapClaims{}
    
	
	toke, err := jwt.ParseWithClaims(token, claims, func(toke *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if toke != nil {
	 }	
	 if err != nil {
	 }
	
    parts := strings.Split(token, ".")
    err = jwt.SigningMethodRS256.Verify(strings.Join(parts[0:2], "."), parts[2], key)
    if err != nil {
        return "false", nil
    }
    return claims["user_name"].(string), nil
}


