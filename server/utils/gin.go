package utils

import "github.com/gin-gonic/gin"

func GetIp(c *gin.Context) string {
	return c.ClientIP()
}
