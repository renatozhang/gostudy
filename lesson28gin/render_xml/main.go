package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/moreXML", func(ctx *gin.Context) {
		type MessageRecord struct {
			Name    string
			Message string
			Numeber int
		}
		var msg MessageRecord
		msg.Name = "lena"
		msg.Message = "hey"
		msg.Numeber = 123
		ctx.XML(http.StatusOK, msg)
	})
	router.Run()
}
