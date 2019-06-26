package service

import (
	"github.com/ha666/logs"
	"os"
	"runtime"
	"strings"
)

// 验证调用来源
func validCallPath() {
	pc, _, _, _ := runtime.Caller(2)
	name := runtime.FuncForPC(pc).Name()
	if !strings.HasPrefix(name, "github.com/hardbornman/garglewool-service/controller/") &&
		!strings.HasPrefix(name, "github.com/hardbornman/garglewool-service/worker/") &&
		!strings.HasPrefix(name, "github.com/hardbornman/garglewool-service/service.") {
		logs.Emergency("【validCallPath】引用错误,%s", name)
		os.Exit(0)
	}
}
