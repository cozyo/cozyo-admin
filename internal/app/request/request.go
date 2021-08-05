package request

import (
	res "github.com/cozyo/internal/app/response"
	"github.com/gin-gonic/gin"
	"strings"
)


// GetToken 获取用户令牌
func GetToken(c *gin.Context) string {
	var token string
	auth := c.GetHeader("Authorization")
	prefix := "Bearer "
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	return token
}
// GetUserID 获取用户ID
func GetUserID(c *gin.Context) string {
	return c.GetString(res.UserIDKey)
}

// SetUserID 设定用户ID
func SetUserID(c *gin.Context, userID string) {
	c.Set(res.UserIDKey, userID)
}