package bootstrap

import (
	"goproject/app/log"
	"os"
	"os/exec"
	"strings"
)

//辅助函数
func GetCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	log.LogError(err)
	i := strings.LastIndex(s, "\\")

	path := string(s[0 : i+1])
	return path
}

