package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserID   string `json: "userid"`
	Login    string `json: "login"`
	Password string `json: "password"`
}

var userMap = make(map[string]*User, 0)

func GeyQuerryController(c *gin.Context) {
	query := c.Query("userID")

	c.JSON(200, gin.H{
		"fine": userMap[query],
	})
}

func PostBodyController(c *gin.Context) {
	u := User{}
	if err := c.ShouldBindBodyWithJSON(&u); err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}
	userID := strconv.Itoa(len(userMap))
	userMap[userID] = &u

	c.JSON(200, gin.H{
		"fine": gin.H{
			"userID":   userID,
			"login":    u.Login,
			"password": u.Password,
		},
	})
}

func PostBodyQueryController(c *gin.Context) {
	u := User{}
	userID := c.Query("userID")
	if err := c.ShouldBindBodyWithJSON(&u); err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}
	userMap[userID] = &u

	c.JSON(200, gin.H{
		"fine": gin.H{
			"userID":   userID,
			"login":    u.Login,
			"password": u.Password,
		},
	})
}

func PutBodyQueryController(c *gin.Context) {
	u := User{}
	if err := c.ShouldBindBodyWithJSON(&u); err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}

	userMap[u.UserID] = &u
	c.JSON(200, gin.H{
		"fine": gin.H{
			"userID":   u.UserID,
			"login":    u.Login,
			"password": u.Password,
		},
	})
}

func PatchBodyQueryController(c *gin.Context) {
	u := User{}
	if err := c.ShouldBindBodyWithJSON(&u); err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}
	currentUser := userMap[u.UserID]

	if u.Login != "" && u.Login != currentUser.Login {
		currentUser.Login = u.Login
	}

	if u.Password != "" && u.Password != currentUser.Password {
		currentUser.Password = u.Password
	}

	c.JSON(200, gin.H{
		"fine": gin.H{
			"userID":   u.UserID,
			"login":    currentUser.Login,
			"password": currentUser.Password,
		},
	})
}
