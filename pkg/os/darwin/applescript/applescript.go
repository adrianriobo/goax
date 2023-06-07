//go:build darwin
// +build darwin

package applescript

// #cgo darwin CFLAGS:-mmacosx-version-min=11 -x objective-c
// #cgo darwin LDFLAGS: -lobjc -framework Foundation -framework Cocoa
//#include "applescript.h"
import "C"
import (
	"time"
	"unsafe"
)

var (
	longDelay = 30 * time.Second
)

func Keystroke(keystroke string) {
	cKeystroke := C.CString(keystroke)
	defer C.free(unsafe.Pointer(cKeystroke))
	C.Keystroke(cKeystroke)
	time.Sleep(longDelay)
}
