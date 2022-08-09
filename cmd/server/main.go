package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-quilting-api/pkg/logging"
)

func health(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func main() {
	r := gin.New()

	r.UseRawPath = true
	r.Use(gin.Recovery())
	r.GET("/health", health)

	v1 := r.Group("/v1")
	{
		v1.Use(logging.LogJSON("quilting-api"))
	}

	err := r.Run()
	if err != nil {
		fmt.Println("failed with error:", err.Error())
	}
}
