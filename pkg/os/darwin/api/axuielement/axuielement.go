//go:build darwin
// +build darwin

package axuielement

// #cgo darwin CFLAGS:-mmacosx-version-min=11 -x objective-c
// #cgo darwin LDFLAGS: -lobjc -framework Foundation -framework Cocoa
//#include "axuielement.h"
import "C"
import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"

	"github.com/adrianriobo/goax/pkg/os/darwin/api/util"
)

func Press(ref util.Ref) {
	C.Press(C.CFTypeRef(ref))
}

func SetValue(ref util.Ref, value string) {
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	C.SetValue(C.CFTypeRef(ref), cValue)
}

func GetAllChildren(ref util.Ref) []util.Ref {
	var children []util.Ref
	has, err := strconv.ParseBool(C.GoString(C.HasChildren(C.CFTypeRef(ref))))
	if err == nil && has {
		childrenASCFArray := C.GetChildren(C.CFTypeRef(ref))
		count := C.CFArrayGetCount(childrenASCFArray)
		for i := 0; i < int(count); i++ {
			children = append(children, util.Ref(C.GetChild(childrenASCFArray, C.CFIndex(i))))
		}
	}
	return children
}

// TODO id should be only title or description
func GetID(ref util.Ref) (string, error) {
	id := C.GoString(C.GetTitle(C.CFTypeRef(ref)))
	if len(id) == 0 {
		id = C.GoString(C.GetValue(C.CFTypeRef(ref)))
	}
	if len(id) == 0 {
		id = C.GoString(C.GetDescription(C.CFTypeRef(ref)))
	}
	if len(id) == 0 {
		return "", fmt.Errorf("object has no id")
	}
	return strings.TrimSpace(id), nil
}

func GetRole(ref util.Ref) string {
	return C.GoString(C.GetRole(C.CFTypeRef(ref)))
}

func ShowActions(ref util.Ref) {
	C.ShowActions(C.CFTypeRef(ref))
}
