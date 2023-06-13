//go:build darwin
// +build darwin

package util

// #cgo darwin CFLAGS:-mmacosx-version-min=11 -x objective-c
// #cgo darwin LDFLAGS: -lobjc -framework Foundation -framework Cocoa
//#include <CoreFoundation/CoreFoundation.h>
import "C"

type Ref C.CFTypeRef
