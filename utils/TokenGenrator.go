package utils

import (
	"crypto/sha1"
	"encoding/base64"
	//fmt --
	_ "fmt"
	"time"
)

//TokenGenerator -- genrate authToken
type TokenGenerator struct {
}

//GenerateToken -- genrate authToken
func (token TokenGenerator) GenerateToken(email string) string {
	var authKey string
	hash := sha1.New()
	t := time.Now().String()
	hash.Write([]byte(email + t))
	authKey = base64.URLEncoding.EncodeToString(hash.Sum(nil))
	//remove = index
	s1 := authKey
	if last := len(s1) - 1; last >= 0 && s1[last] == '=' {
		s1 = s1[:last]
	}
	return s1
}
