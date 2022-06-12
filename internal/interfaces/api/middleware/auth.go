package middleware

import "github.com/gin-gonic/gin"


func Authorization(c *gin.Context) {
	// TODO: Add auth logic
	c.Next()
}
