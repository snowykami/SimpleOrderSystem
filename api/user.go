package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func Register(c *gin.Context) {
	username := c.PostForm("username") // 用户名
	password := c.PostForm("password") // 密码明文

	status, msg, err := UserRegister(username, password)
	if err != nil {
		c.JSON(200, gin.H{
			"message": msg,
			"status":  "error",
		})
		return

	}

	if status {
		c.JSON(200, gin.H{
			"message": "register",
			"status":  "ok",
		})
	} else {
		c.JSON(200, gin.H{
			"message": msg,
			"status":  "error",
		})
	}
}

func Login(c *gin.Context) {
	username := c.PostForm("username") // 用户名
	password := c.PostForm("password") // 密码明文

	user, msg := UserLogin(username, password)
	if user == nil {
		c.JSON(200, gin.H{
			"message": msg,
			"status":  "error",
		})
		return
	}

	jwtString, err := generateJWT(strconv.FormatInt(user.ID, 10), secretKey)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "generate jwt error",
			"status":  "error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "login",
		"status":  "ok",
		"data": gin.H{
			"token": jwtString,
		},
	})
	user.SecretKey = secretKey
	DB.Save(&user)
}

func JoinAct(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	auth, msg, user := Authorize(tokenString)
	if !auth {
		c.JSON(401, gin.H{
			"message": msg,
			"status":  "Unauthorized",
		})
		return
	}

	actID, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	var activity Activity
	DB.First(&activity, actID)

	err := DB.Model(&activity).Association("Users").Append(&user)
	if err != nil {
		return
	}

	c.JSON(200, gin.H{
		"message": "join activity",
		"status":  "ok",
	})

}

func AuthTest(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	claims, err := parseJWT(tokenString)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "auth failed",
			"status":  "error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "auth",
		"status":  "ok",
		"data":    claims,
	})
}
