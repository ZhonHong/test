package middlewares

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//				    加密/解析token
// DB -> Localhost <--------------> client , 減少呼叫DB資源，透過token解析獲得用戶資訊，後續就不須與DB請求用戶資訊

func GenerateToken(userId int, userName string, userEmail string) (string, error) {
	secretKey := []byte("your-secret-key")
	token := jwt.New(jwt.SigningMethodHS256)

	//Claim
	claims := token.Claims.(jwt.MapClaims) //宣告claims為jwt.MapClaims類型，用來儲存user
	claims["userId"] = userId
	claims["userName"] = userName
	claims["userEmail"] = userEmail

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Parse Token
func ParseToken(c *gin.Context) (string, *jwt.MapClaims) { //(string=>err , claims=>*jwt.MapClaims)對應getUserData
	// Request Token
	tokenString := c.GetHeader("Authorization") //請求HTTP Header字段，用來傳遞驗證token

	// Authorization Token nil
	if tokenString == "" {
		return "Token is missing", nil //請求Header中找不到token，回傳字串及jwt.MapClaims nil
	}

	// Parse Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// tokenString用於解析token，token *jwt.Token為回調函數用於返還密鑰(用於驗證token是否有效)，使interface{}可以驗證
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
		claims, ok := token.Claims.(jwt.MapClaims) //返回接口宣告為jwt.MapClaims類型，此物件包含用戶訊息

		if !ok {
			return "Failed to extract claims", nil
		} else {
			return "", &claims //ture則回傳""表示無錯誤及回傳指向claims的指標使得可調用claims
		}

	} else {
		return "Token is invalid", nil
	}
}
