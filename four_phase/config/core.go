package config

import (
	"fmt"
	"go_work/four_phase/modules/admin"
	"os"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error
var Conf *Config

// var Key string = "123456789abc"

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		return
	}
	fmt.Printf("当前工作目录: %s\n", cwd)
	Conf, err = LoadConfig(cwd + "/config/config.yml")
	if err != nil {
		fmt.Println("获取配置文件失败", err)
		return
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Conf.Database.Username, Conf.Database.Password, Conf.Database.Url, Conf.Database.Database)
	fmt.Println("数据库连接地址", dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败", err)
		return
	}
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
			return []byte(Conf.SecretKey), nil
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
