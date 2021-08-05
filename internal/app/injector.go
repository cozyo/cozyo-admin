package app

import (
	"github.com/cozyo/pkg/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

// InjectorSet 注入Injector
var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Engine         *gin.Engine
	Auth           auth.Auther
	gorm           *gorm.DB
}