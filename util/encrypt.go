package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
	"f.in/v/utils"
	"github.com/dgrijalva/jwt-go"
)

/*
md5加密
	before:加密前的字符串
	after:加密后的字符串
*/
func MD5(before string) (after string) {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(before))
	cipherStr := md5Ctx.Sum(nil)
	after = hex.EncodeToString(cipherStr)
	return
}

//获取token
func MakeToken(sub string) (token string) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
		Subject:   sub,
	})

	token, _ = t.SignedString(tokenKey)
	return
}

func UnmarshalToken(tokenStr string) (subStr string, expiresTime time.Time, err error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return tokenKey, nil
	})
	if err != nil {
		return "", time.Now(), err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["sub"].(string), time.Unix(utils.Int64Must(claims["exp"]), 0), nil
	} else {
		//return "", time.Now(), errors.New("uid验证失败")
	}
	return
}