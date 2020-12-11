package bootstrap

import (
	"encoding/json"
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


func JSONToMap(str string) map[string]interface{} {

	var tempMap map[string]interface{}

	err := json.Unmarshal([]byte(str), &tempMap)

	if err != nil {
		panic(err)
	}

	return tempMap
}

