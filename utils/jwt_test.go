package utils

import (
	"testing"
)

func TestJwtParseToken(t *testing.T) {
	tokenString, err := JwtGenerateToken("2", 0)
	//'Accept' => 'application/json',
	//'Authorization' => 'Bearer '.$accessToken,
	claims, err := JwtParseToken(tokenString)
	if err == nil {
		t.Log(claims)
	} else {
		t.Fatal(err)
	}
}
