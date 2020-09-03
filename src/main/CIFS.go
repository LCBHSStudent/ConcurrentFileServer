package main

import (
	"ConcurrentImageFileServer/src/config"
	"ConcurrentImageFileServer/src/service"
	"ConcurrentImageFileServer/src/utils/dateUtil"
	"ConcurrentImageFileServer/src/utils/logUtil"
	"fmt"
	"strings"
)

const utilName = "main"

func main() {
	config.InitConfig()
	printUtilMessage("Initialization of server config data has been finished")
	service.StartImageFileService()
	printUtilMessage("Succeed to run an image file visit server instance...")
	service.StartOperationServer()
	printUtilMessage("Succeed to run an CIFS operation server instance...")

	select {}
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