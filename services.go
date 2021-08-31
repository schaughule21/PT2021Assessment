package main

import (
	"31Aug-Assessment/database"
	"31Aug-Assessment/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var TokensMap = map[string]*models.LoginToken{}

func Login(c *gin.Context) {
	var user models.Factory
	err := c.Bind(&user)
	database.CheckError(err)
	res := database.CheckValidation(user)
	if user.UName != res.UName || user.Password != res.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	token, err := CreateToken(res.SrNo)
	if err != nil {
		log.Fatal(err)
	}
	lt := models.LoginToken{
		ID:    res.SrNo,
		Token: token,
	}
	TokensMap[token] = &lt
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, bson.M{"token": token, "user": res})
}

func CreateToken(userid string) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "Siddhesh") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
