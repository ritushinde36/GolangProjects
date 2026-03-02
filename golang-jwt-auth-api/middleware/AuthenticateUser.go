package middleware

import (
	"fmt"
	"jwt-auth-sql/models"
	"jwt-auth-sql/pkg"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var Cookie_user models.User

func AuthenticateUser(c *gin.Context) {
	pkg.LoadEnvVariables()

	//get the cookie from the request
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)

	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("PRIVATEKEY")), nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//check if cookie is expired. is it is, then the user has to login again
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//get the user id and look up user
		var found_user models.User
		pkg.DB.First(&found_user, claims["user_id"])
		if found_user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		Cookie_user = found_user

		//send the userdata of that user id
		c.Set("user", found_user)

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Next()
}
