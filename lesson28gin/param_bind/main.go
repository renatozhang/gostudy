package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	// example for binding JSON ("user":"manu","password":"123")
	router.POST("/loginJson", func(ctx *gin.Context) {
		var login Login
		if err := ctx.ShouldBindJSON(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)
			ctx.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// example for binding form ("user":"manu","password":"123")
	router.POST("/loginForm", func(ctx *gin.Context) {
		var login Login
		if err := ctx.ShouldBind(&login); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	// example for binding HTML queryString ("user":"manu","password":"123")
	router.GET("/loginForm", func(ctx *gin.Context) {
		var login Login
		if err := ctx.ShouldBind(&login); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		}
	})
	router.Run()
}
