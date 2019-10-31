package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//HeaderSet return header setting middleware function
func HeaderSet() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*
		 */
		//ENV PASS

		OriginHost := os.Getenv("ORIGIN_HOST")
		if OriginHost == "" {
			log.Fatalf("Origin env set: Empty host")
		}

		OriginPort := os.Getenv("ORIGIN_PORT")
		if OriginPort == "" {
			log.Fatalf("Origin env set: Empty port")
		}
		url := "http://" + OriginHost + ":" + OriginPort
		fmt.Println("origin", url)

		c.Writer.Header().Set("Access-Control-Allow-Origin", url)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, X-Requested-With, Content-Type, Accept")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Content-Type", "application/json")

		if c.Request.Method != "OPTIONS" {
			//CORS header
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
			c.Abort()
		}
		return

	}
}
