package middlewares

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateToken(userId int, userName string, userEmail string) (string, error) {
	secretKey := []byte("your-secret-key")
	token := jwt.New(jwt.SigningMethodHS256)

	//Claim
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = userId
	claims["userName"] = userName
	claims["userEmail"] = userEmail

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(c *gin.Context) (string, *jwt.MapClaims) {
	// Request Token
	tokenString := c.GetHeader("Authorization")

	// Authorization Token nil
	if tokenString == "" {
		return "Token is missing", nil
	}

	// Parse Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 在这里提供令牌的签名密钥
		return []byte("your-secret-key"), nil
	})

	// 输出一些调试信息
	fmt.Println("Token String:", tokenString)
	fmt.Printf("Error Details: %+v\n", err)
	// Token err
	if err != nil {
		return err.Error(), nil
	}

	// Authorization Token Vaild
	if token.Valid {
		// Token Cliams
		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			return "Failed to extract claims", nil
		} else {
			return "", &claims
		}

	} else {
		return "Token is invalid", nil
	}
}
