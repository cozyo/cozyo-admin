// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/cozyo/internal/app/router"
	"github.com/google/wire"
)

// BuildInjector 生成注入器
func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		// mock.MockSet,
		InitGormDB,
		InitAuth,
		router.InitWebEngine,
		router.RouterSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}