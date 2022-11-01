package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// gin.H is a shortCut for map[string]interface{}]
	router.GET("/someJson", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	router.GET("moreJson", func(ctx *gin.Context) {
		// you also can use a struct
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "lena"
		msg.Message = "hey"
		msg.Number = 123
		//note that msg name become "user" in the Json
		ctx.JSON(http.StatusOK, msg)
	})
	router.Run()
}
