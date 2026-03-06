package ast

import (
	"hab/global"
	"path/filepath"
)

func init() {
	global.HAB_CONFIG.AutoCode.Root, _ = filepath.Abs("../../../")
	global.HAB_CONFIG.AutoCode.Server = "server"
}
