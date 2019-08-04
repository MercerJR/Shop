package main

import (
	"github.com/gin-gonic/gin"
	"transaction/rabbit"
)
func main() {
	router := gin.Default()
	router.POST("/pay", func(context *gin.Context) {
		uID := context.PostForm("uID")
		gID := context.PostForm("gID")
		sID := context.PostForm("sID")
		number := context.PostForm("number")

		rabbit.Rabbit(uID,gID,sID,number)

		context.JSON(200, gin.H{
			"uID":    uID,
			"gID":    gID,
			"sID":    sID,
			"number": number,
		})
	})
	router.Run()
}