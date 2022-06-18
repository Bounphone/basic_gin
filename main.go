package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"basic_gin/handle"
)

func main() {
	r := gin.Default()
	m := handle.NewMember()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/member", func(c *gin.Context){
		c.JSON(http.StatusOK, m)
	})
	r.Run(":2000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}