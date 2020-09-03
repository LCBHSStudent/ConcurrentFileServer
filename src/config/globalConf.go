package config

import (
	"ConcurrentImageFileServer/src/utils/logUtil"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	dateUtil "ConcurrentImageFileServer/src/utils/dateUtil"
)

const utilName = "config"

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)


var PWD string
var ProjectName = "Default-Project"

func initGlobalConfig(projName string) {
	ProjectName = projName

	PWD, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		printUtilMessage(err.Error())
	} else {
		if runtime.GOOS == "windows" {
			PWD = strings.ReplaceAll(PWD, `\`, "/")
		}
		PWD += "/"
		printUtilMessage("Current project working dir: " + PWD)
	}

}

func printUtilMessage(msg string) {
	logMsg := strings.Join(
		[]string {
			logUtil.Red(dateUtil.GetCurFormatTime()),
			" [Module ", logUtil.Blue(utilName),  "] ", msg,
		}, "",
	)
	fmt.Println(logMsg)
}