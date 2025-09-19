package core

import (
	"fmt"
	"go_work/four_phase/modules/admin"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

var Key string = "123456789abc"

func init() {
	dsn := "root:hcrz1234@tcp(192.168.1.123:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

// 从token中解析用户信息
func ParseToken(tokenString string) (*admin.CustomClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(
		tokenString,
		&admin.CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// 验证签名方法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(Key), nil
		},
	)

	// 验证token有效性
	if err != nil {
		return nil, err
	}

	// 提取claims中的用户信息
	if claims, ok := token.Claims.(*admin.CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
