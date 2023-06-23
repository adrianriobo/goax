//go:build darwin
// +build darwin

package appkit

// #cgo darwin CFLAGS:-mmacosx-version-min=11 -x objective-c
// #cgo darwin LDFLAGS: -lobjc -framework Foundation -framework Cocoa
//#include "nsrunningapplication.h"
import "C"
import (
	"unsafe"

	"github.com/adrianriobo/goax/pkg/os/darwin/api/util"
)

func GetFrontmostApplication() unsafe.Pointer {
	return C.FrontmostApplication()
}

func CreateAX(ref unsafe.Pointer) util.Ref {
	return util.Ref(C.CreateApplicationAXRef(ref))
}

func GetAXFocusedWindow(ref util.Ref) util.Ref {
	fwAXRef := C.GetAXFocusedWindow(C.CFTypeRef(ref))
	return util.Ref(fwAXRef)
}

// TODO windowTitle could be localized, how to handle that??
func GetAppByBundleAndWindow(bundleID, windowTitle string) unsafe.Pointer {
	cBundleID := C.CString(bundleID)
	defer C.free(unsafe.Pointer(cBundleID))
	cWindowTitle := C.CString(windowTitle)
	defer C.free(unsafe.Pointer(cWindowTitle))
	return C.FindRunningApplication(cBundleID, cWindowTitle)
}
