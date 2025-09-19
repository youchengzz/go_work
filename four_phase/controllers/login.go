package controllers

import (
	"fmt"
	core "go_work/four_phase/config"
	"go_work/four_phase/modules/admin"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
}

func (LoginController) Register(c *gin.Context) {
	user := admin.User{}
	c.ShouldBindJSON(&user)
	fmt.Println(user)
	var num int64
	if err := core.DB.Table("user").Where("username = ?", user.Username).Count(&num).Error; err != nil {
		core.Err("register", err)
		return
	}
	if num > 0 {
		core.Err500("用户名已被注册", c)
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)
	registerRrr := core.DB.Create(&user).Error
	core.Err("register", registerRrr)
}

func (LoginController) Login(c *gin.Context) {
	user := admin.User{}
	c.ShouldBind(&user)
	fmt.Printf("登录信息%#v", user)
	storedUser := admin.User{}
	tx := core.DB.Where("username = ?", user.Username).First(&storedUser)
	if tx.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"success": false,
			"message": "登录失败",
		})
		return
	}
	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       storedUser.BaseModel.Id,
		"username": storedUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(core.Conf.SecretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"success": true,
		"message": "登录成功",
		"token":   tokenString,
	})
}
