package middleware

import (
	"Assignment2/model"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authentication(c *gin.Context) {
	db := model.GetDB()
	// get request cookie
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		fmt.Println("Error retrieving token from cookie")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token not found",
		})
		c.Abort()
		return
	}
	fmt.Println("Token:", tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("brianhabib252"), nil
	})
	if err != nil {
		fmt.Println("Error Validating JWT :", err)
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// check token exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// check user
		var user model.User
		db.First(&user, claims["sub"])
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User not found",
			})
			c.Abort()
			return
		}
		// Attach to req and continue
		c.Set("user", user)
		c.Next()
	} else {
		fmt.Println(err)
	}
}
