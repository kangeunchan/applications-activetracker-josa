package activeapplicationinfo

/*
#cgo CFLAGS: -x objective-c -I../interface
#cgo LDFLAGS: -framework Cocoa
#import "./../interface/ActiveApplicationInfo.h"
*/
import "C"
import "fmt"

type ActiveApplicationInfo struct {
	AppName     string
	RunningTime string
}

func NewActiveApplicationInfo() *ActiveApplicationInfo {
	appName := C.GoString(C.getActiveAppName())
	runningTime := C.GoString(C.getActiveAppRunningTime())

	return &ActiveApplicationInfo{
		AppName:     appName,
		RunningTime: runningTime,
	}
}

func (info *ActiveApplicationInfo) GetAppName() string {
	return info.AppName
}

func (info *ActiveApplicationInfo) GetRunningTime() string {
	return info.RunningTime
}

func (info *ActiveApplicationInfo) PrintInfo() {
	fmt.Printf("App Name: %s, Running Time: %s\n", info.AppName, info.RunningTime)
}
