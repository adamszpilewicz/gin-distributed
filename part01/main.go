package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/:name", IndexHandler)
	router.Run(":8888")
}

func IndexHandler(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"message": "hello " + name,
	})
}
