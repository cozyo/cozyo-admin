package router

import (
	"github.com/gin-gonic/gin"
)

func ApiRouter(r *gin.RouterGroup)  {
	v1 := r.Group("/v1")
	{
		gUser := v1.Group("users")
		{
			gUser.GET("/test_api", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "test_api",
				})
			})
		}
	}
}

