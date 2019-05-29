package common

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"wc_inner_server/common/code"
	"wc_inner_server/model/table"
)

func Md5Compute(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(code.SecretKey),nil
	}
}


func ParseToken(token string) (token2 *jwt.Token, err error) {
	user := &model.UserAdmin{}
	tokenJwt, err := jwt.Parse(token, secret())

	if tokenJwt != nil {
		claim,_ := tokenJwt.Claims.(jwt.MapClaims)

		user.No =claim["no"].(string)
	}

	return tokenJwt, err
}
