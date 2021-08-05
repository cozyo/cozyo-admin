package router

import (
	"github.com/cozyo/internal/app/middleware"
	"github.com/cozyo/pkg/auth"
	"github.com/cozyo/router"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var _ IRouter = (*Router)(nil)
// RouterSet 注入router
var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

type IRouter interface {
	Register(engine *gin.Engine) error
	Prefixes() []string
}

// Router 路由管理器
type Router struct {
	Auth           auth.Auther
	//CasbinEnforcer *casbin.SyncedEnforcer
	//DemoAPI        *api.Demo
	//LoginAPI       *api.Login
	//MenuAPI        *api.Menu
	//RoleAPI        *api.Role
	//UserAPI        *api.User
}

// Register 注册路由
func (a *Router) Register(app *gin.Engine) error {
	group := app.Group("/api")

	group.Use(middleware.UserAuthMiddleware(a.Auth,
		middleware.AllowPathPrefixSkipper("/api/v1/pub/login"),
	))
	group.Use(middleware.RateLimiterMiddleware())
	router.ApiRouter(group)
	return nil
}

// Prefixes 路由前缀列表
func (a *Router) Prefixes() []string {
	return []string{
		"/api/",
	}
}