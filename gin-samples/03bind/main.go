package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	ID    int    `uri:"id" json:"id" binding:"required"`
	Title string `form:"title" json:"title"`
	Score int    `form:"score" json:"score"`
}

func sampleHandler(c *gin.Context) {
	request := Request{}
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, request)
}

func main() {
	r := gin.New()
	r.Use(gin.Recovery())

	r.POST("/sample/:id", sampleHandler)

	r.Run(":8080")
}
