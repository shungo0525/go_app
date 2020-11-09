package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	data := "Hello Go/Gin!!"
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"data": data})
	})
	router.GET("/show", func(ctx *gin.Context) {
		ctx.HTML(200, "show.html", gin.H{})
	})

	router.Run(":3000")
}
