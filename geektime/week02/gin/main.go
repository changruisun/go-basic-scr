package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

	server.Use(func(ctx *gin.Context) {
		println("this is the first middleware")
	}, func(ctx *gin.Context) {
		println("this is the second middleware")
	})
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello World")
	})
	server.GET("/user/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "hello "+name)
	})
	server.GET("/order", func(ctx *gin.Context) {
		id := ctx.Query("id")
		ctx.String(http.StatusOK, "hello "+id)
	})

	server.GET("/views/*.html", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello "+ctx.Param(".html"))

	})

	// go func() {
	//	  server.Run(":8081")
	// }()
	server.Run(":8080")

}
