package main

import (
	"github.com/gin-gonic/gin"
)

// This middleware sets whether the session was created or not
func setSessionStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if sessionId, err := c.Cookie("session_id"); err == nil || sessionId != "" {
			c.Set("session_created", true)
		} else {
			c.Set("session_created", false)
		}
	}
}
