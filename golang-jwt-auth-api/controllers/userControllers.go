package controllers

import (
	"fmt"
	"jwt-auth-sql/middleware"
	"jwt-auth-sql/models"
	"jwt-auth-sql/pkg"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	// "github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	//user will send data and we need to save it in DB

	//get the data from the request
	var new_user models.User
	err := c.ShouldBindJSON(&new_user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//hash the password so its encrypted
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(new_user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash the password"})
		return
	}

	//save the user in the DB
	updated_user := models.User{
		Name:     new_user.Name,
		Password: string(hashed_password),
		Age:      new_user.Age,
	}

	result := pkg.DB.Create(&updated_user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not able to create user"})

		return
	}

	//respond
	c.JSON(http.StatusOK, gin.H{"success": "User added successfully"})

}

func Login(c *gin.Context) {
	//get the data from the request

	var login_user models.User
	var found_user models.User

	err := c.ShouldBindJSON(&login_user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//look up the requested user in the DB
	pkg.DB.First(&found_user, "name = ?", login_user.Name)
	if found_user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not find the user"})
		return
	}

	//compare the sent password with the password on the db
	err = bcrypt.CompareHashAndPassword([]byte(found_user.Password), []byte(login_user.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}
	//generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": found_user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	pkg.LoadEnvVariables()

	//sign the token with a private key
	tokenString, err := token.SignedString([]byte(os.Getenv("PRIVATEKEY")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	//send the token back using a cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func GetUserDataByID(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{"message": user})

}

func GetAllUsers(c *gin.Context) {
	var all_users []models.User
	result := pkg.DB.Find(&all_users)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to retrive users"})
	}
	fmt.Print(all_users)
	c.JSON(http.StatusOK, gin.H{"message": all_users})

}

func UpdateUser(c *gin.Context) {

	// user, _ := c.Get("user")

	//get the data from the request
	var new_user models.User
	err := c.ShouldBindJSON(&new_user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	middleware.Cookie_user.Age = new_user.Age
	pkg.DB.Save(&middleware.Cookie_user)
	c.JSON(http.StatusOK, gin.H{"message": "user is updated"})

}

func DeleteUser(c *gin.Context) {
	user, _ := c.Get("user")
	pkg.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "user is deleted"})

}
