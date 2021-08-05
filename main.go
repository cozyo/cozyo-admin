package main

import (
	"context"
	"fmt"
	"github.com/cozyo/internal/app"
	"github.com/cozyo/pkg/logger"
	"github.com/go-ini/ini"
)

var (
	version string
	rootPath string
	modelFile string
	menuFile string
	configFile string
)

type EnvError struct {
	Name string
	Msg     string
}

func (e *EnvError) Error() string { return e.Msg }


func init() {
	cfg, err := ini.Load(".env")
	if err != nil {
		logger.Fatalf("Fail to read file: %s", err.Error())
	}
	if err := checkEnvFileValue(cfg);err != nil {
		logger.Fatalf("Invalid to read file: %s", err.Error())
	}
}

func main()  {
	logger.SetVersion(version)
	ctx := logger.NewTagContext(context.Background(), "__main__")

	if err := app.Run(ctx,
		app.SetConfigFile(configFile),
		app.SetModelFile(modelFile),
		app.SetRootDir(rootPath),
		app.SetMenuFile(menuFile),
		app.SetVersion(version));
	err != nil {
		logger.WithContext(ctx).Errorf(err.Error())
	}
}

// 检查env配置文件是否正确
func checkEnvFileValue(cfg *ini.File) error {
	var e = new(EnvError)
	e.Name = "env"
	fmt.Println(e)
	if version = cfg.Section("").Key("Version").String();version == "" {
		e.Msg = "Version is Not defined"
		return e
	}
	if rootPath = cfg.Section("").Key("RootPath").String();rootPath == "" {
		e.Msg = "RootPath is Not defined"
		return e
	}
	if modelFile = cfg.Section("").Key("ModePath").String();modelFile == "" {
		e.Msg = "ModePath is Not defined"
		return e
	}
	if menuFile = cfg.Section("").Key("MenuPath").String();menuFile == "" {
		e.Msg = "MenuPath is Not defined"
		return e
	}
	if configFile = cfg.Section("").Key("ConfigPath").String();configFile == "" {
		e.Msg = "ConfigPath is Not defined"
		return e
	}
	return nil
}
