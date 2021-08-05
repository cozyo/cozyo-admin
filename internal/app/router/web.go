package router

import (
	"github.com/LyricTian/gzip"
	"github.com/cozyo/internal/app/conf"
	"github.com/cozyo/internal/app/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitWebEngine(r IRouter) *gin.Engine {
	gin.SetMode(conf.C.RunMode)

	app := gin.New()
	app.NoMethod()
	app.NoRoute()

	prefixes := r.Prefixes()
	// Trace ID
	app.Use(middleware.TraceMiddleware(middleware.AllowPathPrefixNoSkipper(prefixes...)))

	// Copy body
	app.Use(middleware.CopyBodyMiddleware(middleware.AllowPathPrefixNoSkipper(prefixes...)))

	// Access logger
	app.Use(middleware.LoggerMiddleware(middleware.AllowPathPrefixNoSkipper(prefixes...)))

	// Recover
	app.Use(middleware.RecoveryMiddleware())

	// CORS
	if conf.C.CORS.Enable {
		app.Use(middleware.CORSMiddleware())
	}

	// GZIP
	if conf.C.GZIP.Enable {
		app.Use(gzip.Gzip(gzip.BestCompression,
			gzip.WithExcludedExtensions(conf.C.GZIP.ExcludedExtentions),
			gzip.WithExcludedPaths(conf.C.GZIP.ExcludedPaths),
		))
	}

	// Router register
	if err := r.Register(app); err != nil {
		return nil
	}

	// Swagger
	if conf.C.Swagger {
		app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// Website
	if dir := conf.C.WWW; dir != "" {
		app.Use(middleware.RootMiddleware(dir, middleware.AllowPathPrefixSkipper(prefixes...)))
	}

	return app
}