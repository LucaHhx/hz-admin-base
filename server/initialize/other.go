package initialize

import (
	"bufio"
	"os"
	"strings"

	"github.com/songzhibin97/gkit/cache/local_cache"

	"hab/global"
	"hab/utils"
)

func OtherInit() {
	dr, err := utils.ParseDuration(global.HAB_CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.HAB_CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)
	file, err := os.Open("go.mod")
	if err == nil && global.HAB_CONFIG.AutoCode.Module == "" {
		scanner := bufio.NewScanner(file)
		scanner.Scan()
		global.HAB_CONFIG.AutoCode.Module = strings.TrimPrefix(scanner.Text(), "module ")
	}
}
