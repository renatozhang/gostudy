package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32M)
	router.MaxMultipartMemory = 8 << 20 //8MiB
	router.POST("/upload", func(ctx *gin.Context) {
		form, _ := ctx.MultipartForm()
		files := form.File["file"]
		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("C:/tmp/%s_%d", file.Filename, index)
			ctx.SaveUploadedFile(file, dst)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d file uploaded", len(files)),
		})
	})

	router.Run() //listen and serve on 0.0.0.0:8080
}
