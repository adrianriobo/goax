//go:build darwin
// +build darwin

package appkit

// #cgo darwin CFLAGS:-mmacosx-version-min=11 -x objective-c
// #cgo darwin LDFLAGS: -lobjc -framework Foundation -framework Cocoa
//#include "nsrunningapplication.h"
import "C"
import (
	"fmt"
	"time"
	"unsafe"

	"github.com/adrianriobo/goax/pkg/os/darwin/api/axuielement"
	"github.com/adrianriobo/goax/pkg/os/darwin/api/util"
)

var (
	defaultDelay      = 5 * time.Second
	longDelay         = 15 * time.Second
	clickableElements = []string{"AXButton", "AXStaticText", "AXLink"}
	checkableElements = []string{"AXCheckBox"}
)

// https://developer.apple.com/documentation/appkit/nsrunningapplication?language=objc
type NSRunningApplication struct {
	ref              unsafe.Pointer
	bundleIdentifier string
	axRef            util.Ref
	focusedWindow    *axuielement.AXUIElementRef
}

func GetApplication() (*NSRunningApplication, error) {
	// Get the fromtmost application
	app := NSRunningApplication{
		ref: C.FrontmostApplication()}
	// create an ax element to access the app
	app.createAX()
	time.Sleep(defaultDelay)
	// Load all AX elements for the app
	if err := app.LoadFocusedWindow(); err != nil {
		return nil, fmt.Errorf("error getting app")
	}
	return &app, nil
}

func (r *NSRunningApplication) createAX() {
	r.axRef = util.Ref(C.CreateApplicationAXRef(r.ref))
}

func (r *NSRunningApplication) LoadFocusedWindow() (err error) {
	// Get the ax ui ref for the focused window
	fwAXRef := C.GetAXFocusedWindow(C.CFTypeRef(r.axRef))
	// Greate hierachy of elements
	r.focusedWindow, err = axuielement.GetAXUIElementRef(util.Ref(fwAXRef), nil)
	return
}

func (r *NSRunningApplication) Click(id string) error {
	return r.pressElement(id, clickableElements)
}

func (r *NSRunningApplication) Check(id string) error {
	return r.pressElement(id, checkableElements)
}

func (r *NSRunningApplication) pressElement(id string, elementRoles []string) error {
	if err := r.LoadFocusedWindow(); err != nil {
		return err
	}
	element, err := r.getElementbyRoleAndID(id, clickableElements)
	if err != nil {
		return err
	}
	element.Press()
	time.Sleep(longDelay)
	return nil
}

func (r *NSRunningApplication) getElementbyRoleAndID(id string, elementTypes []string) (*axuielement.AXUIElementRef, error) {
	for _, ct := range elementTypes {
		clickable, err := r.focusedWindow.FindElementByRoleAndID(ct, id)
		if err == nil {
			return clickable, nil
		}
	}
	return nil, fmt.Errorf("not found any clickable element with id %s", id)
}
