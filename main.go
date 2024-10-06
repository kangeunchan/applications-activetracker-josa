package main

import (
	"fmt"
	activeapplicationinfo "josa/functions/c/ActiveApplicationInfo/service"
)

func main() {
	appInfo := activeapplicationinfo.NewActiveApplicationInfo()

	appName := appInfo.GetAppName()
	runningTime := appInfo.GetRunningTime()

	fmt.Printf("App Name: %s, Running Time: %s\n", appName, runningTime)
}
