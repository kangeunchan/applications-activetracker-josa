package main

import (
	"fmt"
)

func main() {
	// ActiveApplicationInfo 객체 생성
	appInfo := NewActiveApplicationInfo

	// 객체 메서드를 사용하여 정보 가져오기
	appName := appInfo.GetAppName()
	runningTime := appInfo.GetRunningTime()

	// 출력
	fmt.Printf("App Name: %s, Running Time: %s\n", appName, runningTime)

	// 또는 PrintInfo 메서드를 사용하여 모든 정보 출력
	appInfo.PrintInfo()
}
