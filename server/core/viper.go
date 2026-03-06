package core

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"hab/global"
)

// Viper //
// 优先级: 命令行 > 环境变量 > 默认值
// Author [SliverHorn](https://github.com/SliverHorn)
func Viper() *viper.Viper {
	config := os.Getenv("HAB_CONFIG")
	if config == "" {
		config = "config.yaml"
	}
	fmt.Printf("config的路径为%s\n", config)

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.HAB_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.HAB_CONFIG); err != nil {
		panic(err)
	}

	// root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	global.HAB_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return v
}
