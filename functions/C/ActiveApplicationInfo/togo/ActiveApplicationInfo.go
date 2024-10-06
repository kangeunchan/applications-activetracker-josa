package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include "./../interface/ActiveApplicationInfo.h"
*/
import "C"
import "fmt"

// ActiveApplicationInfo 구조체 정의
type ActiveApplicationInfo struct {
	AppName     string
	RunningTime string
}

// NewActiveApplicationInfo 함수로 새로운 객체 생성
func NewActiveApplicationInfo() *ActiveApplicationInfo {
	// C 함수 호출을 통해 Objective-C에서 데이터를 받아옴
	appName := C.GoString(C.getActiveAppName())
	runningTime := C.GoString(C.getActiveAppRunningTime())

	// 새로운 ActiveApplicationInfo 객체를 반환
	return &ActiveApplicationInfo{
		AppName:     appName,
		RunningTime: runningTime,
	}
}

// 메서드로 앱 이름 반환
func (info *ActiveApplicationInfo) GetAppName() string {
	return info.AppName
}

// 메서드로 실행 시간 반환
func (info *ActiveApplicationInfo) GetRunningTime() string {
	return info.RunningTime
}

// 디버깅용 메서드: 모든 정보 출력
func (info *ActiveApplicationInfo) PrintInfo() {
	fmt.Printf("App Name: %s, Running Time: %s\n", info.AppName, info.RunningTime)
}
