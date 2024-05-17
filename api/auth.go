package api

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// 方便演示用，实际开发中请使用数据库存储
const issuer = "redrocker"
const secretKey = "dGhpcyBpcyBhIHNlY3JldCBrZXkgZm9yIFRlc3Rpbmc="

const (
	adminAccount = "admin"
	adminPasswd  = "admin"
)

func generateJWT(userID string, secretKey string) (string, error) {
	// 定义签名算法

	// 定义密钥

	// 定义JWT声明
	claims := jwt.MapClaims{
		"iss":    issuer,                                // 发行者
		"sub":    userID,                                // 主题，这里使用用户ID
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // 过期时间
		"iat":    time.Now().Unix(),                     // 签发时间
		"userID": userID,                                // 自定义声明，用户ID
	}

	// 创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名并获得完整的编码后的token作为字符串
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func parseJWT(tokenString string) (jwt.MapClaims, error) {
	// 定义密钥
	// 解析token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	// 验证token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return claims, nil
	}
	return nil, err
}

func encryptPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

func UserRegister(username string, password string) (bool, string, error) {
	// 密码入库
	hash, err := passwordHash(password)
	if err != nil {
		Logger.Error("password hash error: " + err.Error())
		return false, "password hash error", err
	}
	// 入库，先查询用户名是否存在
	var user User
	err = DB.Where("username = ?", username).First(&user).Error
	if err == nil || username == adminAccount {
		return false, "username exists", nil
	}
	user = User{
		Username: username,
		Password: hash,
	}
	err = DB.Create(&user).Error
	if err != nil {
		return false, "db create error", err
	}
	return true, "ok", nil
}

// UserLogin 用户登录 username: 用户名 password: 明文密码
func UserLogin(username string, password string) (*User, string) {
	var user User
	if username == adminAccount && password == adminPasswd {
		return &User{
			ID:       0,
			Username: adminAccount,
		}, "ok"
	}

	err := DB.Where("username = ?", username).First(&user).Error
	if err != nil {

		return nil, "user not found"
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "password error"
	}
	return &user, "ok"
}

func passwordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func Authorize(tokenString string) (bool, string, *User) {
	claims, err := parseJWT(tokenString)
	if err != nil {
		return false, "parse jwt error", nil
	}
	userID := claims["userID"].(string)
	var user User
	err = DB.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return false, "user not found", nil
	}

	// 期限判断
	if int64(claims["exp"].(float64)) < time.Now().Unix() {
		return false, "token expired", nil
	}

	return true, "ok", &user
}

func IsAdmin(user *User) bool {
	return user.Username == adminAccount
}
