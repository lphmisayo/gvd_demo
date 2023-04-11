package utils

import (
	"errors"
	"gin_vue_blog_server/global"
	"gin_vue_blog_server/model/system/request"
	"github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	TokenKey []byte
}

var (
	ExpiredToken     = errors.New("令牌已过期")
	NotValidYetToken = errors.New("令牌未生效")
	MalFormedToken   = errors.New("令牌格式错误")
	InvalidToken     = errors.New("非法令牌")
)

func NewJwt() *JWT {
	return &JWT{
		[]byte(global.Config.JWT.TokenKey),
	}
}

// ParseToken 解析Token
func (j *JWT) ParseToken(tokenString string) (*request.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.TokenKey, nil
	})
	//分类错误集
	if err != nil {
		if validerr, ok := err.(*jwt.ValidationError); ok {
			if validerr.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, MalFormedToken
			} else if validerr.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ExpiredToken
			} else if validerr.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, NotValidYetToken
			} else {
				return nil, InvalidToken
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.Claims); ok && token.Valid {
			return claims, nil
		}
		return nil, InvalidToken
	} else {
		return nil, InvalidToken
	}
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims request.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.TokenKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims request.Claims) (string, error) {
	v, err, _ := global.ConcurrencyControl.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}
