package config

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/pelletier/go-toml"
	"github.com/zileyuan/go-working-calendar/model"
	"github.com/zileyuan/go-working-calendar/util"
)

var CalendarData model.CalendarModel

// Runtime 运行时配置实例
var Runtime RuntimeConfig

// RuntimeConfig 运行时配置
type RuntimeConfig struct {
	Common CommonConfig
	RunEnv RunEnvConfig
}

// CommonConfig 通用配置
type CommonConfig struct {
	// AppName 应用名称
	AppName string
	// RunMode 运行模式
	RunMode string
}

// RunEnvConfig 运行环境配置
type RunEnvConfig struct {
	// HTTPPort 启动端口
	HTTPPort int
}

// Configs 配置文件配置
type Configs struct {
	Common CommonConfig
	Dev    RunEnvConfig
	Prod   RunEnvConfig
}

func ReadConfig(pathFile string) {
	cfg, err := toml.LoadFile("./config.toml")
	if err != nil {
		util.Errorf("load config fail: %v\n", err)
	} else {
		configs := Configs{}
		err = cfg.Unmarshal(&configs)
		if err != nil {
			util.Errorf("analyse config fail: %v\n", err)
		} else {
			//get runtime args
			for i := range os.Args {
				arg := strings.ToLower(os.Args[i])
				if arg == "prod" || arg == "produce" {
					configs.Common.RunMode = "Prod"
				}
				if arg == "dev" || arg == "development" {
					configs.Common.RunMode = "Dev"
				}
			}
			Runtime.Common = configs.Common
			if Runtime.Common.RunMode == "Prod" {
				Runtime.RunEnv = configs.Prod
			} else {
				Runtime.RunEnv = configs.Dev
			}
		}
	}
}

func ReadCalendar(pathFile string) {
	bs, err := os.ReadFile(pathFile)
	if err == nil {
		err = json.Unmarshal(bs, &CalendarData)
		if err != nil {
			util.Infof("read calendar data fail: %v\n", err)
		}
	} else {
		util.Infof("read calendar data fail: %v\n", err)
	}
}

func init() {
	ReadConfig("./config.toml")
	ReadCalendar("./calendar.json")
}
