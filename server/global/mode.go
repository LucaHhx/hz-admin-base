package global

import (
	"flag"
	"os"
	"strings"
)

type Mode string

const (
	ModeApi     Mode = "api"
	ModeBackend Mode = "backend"
	ModeTask    Mode = "task"
	ModeAll     Mode = "all"
)

var SysMode Mode = ModeAll

var ServiceList = []Mode{ModeApi, ModeBackend, ModeAll}

func InitSysMode() {
	var mode string
	flag.StringVar(&mode, "type", "all", "the server type")
	flag.Parse()
	SysMode = Mode(mode)

	// 去除type参数
	var args []string
	for _, arg := range os.Args {
		if !strings.HasPrefix(arg, "-type=") && !strings.HasPrefix(arg, "--type=") {
			args = append(args, arg)
		}
	}
	os.Args = args
}
