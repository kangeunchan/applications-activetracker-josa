package main

import (
	"fmt"
	activeapplicationinfo "josa/functions/C/ActiveApplicationInfo/togo"
)

func main() {
	appInfo := activeapplicationinfo.NewActiveApplicationInfo()

	appName := appInfo.GetAppName()
	runningTime := appInfo.GetRunningTime()

	fmt.Printf("App Name: %s, Running Time: %s\n", appName, runningTime)
}
