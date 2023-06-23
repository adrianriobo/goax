//go:build darwin
// +build darwin

package ax

import (
	"fmt"

	axAPI "github.com/adrianriobo/goax/pkg/goax/axapp/api"
	appElements "github.com/adrianriobo/goax/pkg/goax/elements"
	"github.com/adrianriobo/goax/pkg/os/darwin/api/appkit"
	"github.com/adrianriobo/goax/pkg/os/darwin/api/axuielement"
	"github.com/adrianriobo/goax/pkg/os/darwin/api/util"
	"github.com/adrianriobo/goax/pkg/os/darwin/applescript"
)

var elementTypes = map[string]appElements.ElementType{
	"AXWindow": appElements.WINDOW,
	"AXButton": appElements.BUTTON,
	"AXList":   appElements.LIST,
	"AXImage":  appElements.IMAGE,
	// wa.UIA_ListItemControlTypeId:  axAPI.LISTITEM,
	"AXGroup":      appElements.GROUP,
	"AXStaticText": appElements.TEXT,
	// wa.UIA_MenuControlTypeId:      axAPI.MENU,
	"AXMenuButton": appElements.MENUITEM,
	"AXCheckBox":   appElements.CHECKBOX,
	"AXTextField":  appElements.EDIT,
	// wa.UIA_PaneControlTypeId:      axAPI.PANE,
	// wa.UIA_ComboBoxControlTypeId:  axAPI.COMBOBOX,
	// wa.UIA_DocumentControlTypeId:  axAPI.DOCUMENT,
	"AXLink":    appElements.HYPERLINK,
	"AXWebArea": appElements.AREA,
	"AXHeading": appElements.TITLE,
}

type AXElement struct {
	ref util.Ref
}

func GetForegroundRootAXElement() (axAPI.OSAXElement, error) {
	// Get the pointer to the frontmost application
	appPointer := appkit.GetFrontmostApplication()
	// Create ax element to access the app and get its reference
	axRef := appkit.CreateAX(appPointer)
	// Get the ax ref of the focused window
	parentRef := appkit.GetAXFocusedWindow(axRef)
	return AXElement{ref: parentRef}, nil
}

// Get an app based on the bundle id and the title
func GetAppByBundleAndTitleRootAXElement(bundleID, title string) (axAPI.OSAXElement, error) {
	// Get the pointer to the frontmost application
	appPointer := appkit.GetAppByBundleAndWindow(bundleID, title)
	// Create ax element to access the app and get its reference
	axRef := appkit.CreateAX(appPointer)
	// Get the ax ref of the focused window
	parentRef := appkit.GetAXFocusedWindow(axRef)
	return AXElement{ref: parentRef}, nil
}

func (a AXElement) Press() error {
	axuielement.Press(a.ref)
	return nil
}

func (a AXElement) SetValue(value string) error {
	axuielement.SetValue(a.ref, value)
	return nil
}

func (a AXElement) SetValueOnFocus(value string) error {
	applescript.Keystroke(value)
	return nil
}

func (a AXElement) GetAllChildren() ([]axAPI.OSAXElement, error) {
	var children []axAPI.OSAXElement
	for _, child := range axuielement.GetAllChildren(a.ref) {
		children = append(children, AXElement{child})
	}
	return children, nil
}

func (a AXElement) GetType() (*appElements.ElementType, error) {
	role := axuielement.GetRole(a.ref)
	t, ok := elementTypes[role]
	if ok {
		return &t, nil
	}
	return nil, fmt.Errorf("can not cast value to existing AX types, current internal value is %s", role)
}

func (a AXElement) GetID() (string, error) {
	return axuielement.GetID(a.ref)
}

func (a AXElement) GetValue() (string, error) {
	return "", fmt.Errorf("not implemented yet")
}
