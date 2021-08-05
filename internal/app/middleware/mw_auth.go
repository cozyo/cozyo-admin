package middleware

import (
	"github.com/cozyo/internal/app/conf"
	"github.com/cozyo/internal/app/contextS"
	req "github.com/cozyo/internal/app/request"
	res "github.com/cozyo/internal/app/response"
	"github.com/cozyo/pkg/auth"
	"github.com/cozyo/pkg/errors"
	"github.com/cozyo/pkg/logger"
	"github.com/gin-gonic/gin"
)

func wrapUserAuthContext(c *gin.Context, userID string) {
	req.SetUserID(c, userID)
	ctx := contextS.NewUserID(c.Request.Context(), userID)
	ctx = logger.NewUserIDContext(ctx, userID)
	c.Request = c.Request.WithContext(ctx)
}

// UserAuthMiddleware 用户授权中间件
func UserAuthMiddleware(a auth.Auther, skippers ...SkipperFunc) gin.HandlerFunc {
	if !conf.C.JWTAuth.Enable {
		return func(c *gin.Context) {
			wrapUserAuthContext(c, conf.C.Root.UserName)
			c.Next()
		}
	}

	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		userID, err := a.ParseUserID(c.Request.Context(), req.GetToken(c))
		if err != nil {
			if err == auth.ErrInvalidToken {
				if conf.C.IsDebugMode() {
					wrapUserAuthContext(c, conf.C.Root.UserName)
					c.Next()
					return
				}
				res.Error(c, errors.ErrInvalidToken)
				return
			}
			res.Error(c, errors.WithStack(err))
			return
		}

		wrapUserAuthContext(c, userID)
		c.Next()
	}
}
