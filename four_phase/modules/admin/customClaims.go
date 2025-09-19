package admin

import (
	"github.com/dgrijalva/jwt-go"
)

// 自定义claims（包含用户信息）
type CustomClaims struct {
	Id             uint   `json:"user_id"`  // 用户ID
	Username           string `json:"username"` // 用户名
	jwt.StandardClaims        // 标准claims（包含过期时间等）
}
