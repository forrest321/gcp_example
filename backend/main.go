package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		log.Info("Root path visited")
		c.String(http.StatusOK, "Hello, World!!")
	})

	log.Info("Starting server on port 8080")
	r.Run(":8080") // Listen and serve on 0.0.0.0:8080
}
