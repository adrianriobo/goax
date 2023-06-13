package ax

import (
	"fmt"

	axAPI "github.com/adrianriobo/goax/pkg/goax/axapp/api"
	appElements "github.com/adrianriobo/goax/pkg/goax/elements"
	win32waf "github.com/adrianriobo/goax/pkg/os/windows/api/user-interface/windows-accesibility-features"
	win32wam "github.com/adrianriobo/goax/pkg/os/windows/api/user-interface/windows-and-messages"
	"github.com/adrianriobo/goax/pkg/os/windows/util"
	wa "github.com/openstandia/w32uiautomation"
	"golang.org/x/exp/maps"
)

var elementTypes = map[int64]appElements.ElementType{
	wa.UIA_WindowControlTypeId:    appElements.WINDOW,
	wa.UIA_ButtonControlTypeId:    appElements.BUTTON,
	wa.UIA_ListControlTypeId:      appElements.LIST,
	wa.UIA_ListItemControlTypeId:  appElements.LISTITEM,
	wa.UIA_GroupControlTypeId:     appElements.GROUP,
	wa.UIA_TextControlTypeId:      appElements.TEXT,
	wa.UIA_MenuControlTypeId:      appElements.MENU,
	wa.UIA_MenuItemControlTypeId:  appElements.MENUITEM,
	wa.UIA_CheckBoxControlTypeId:  appElements.CHECKBOX,
	wa.UIA_EditControlTypeId:      appElements.EDIT,
	wa.UIA_PaneControlTypeId:      appElements.PANE,
	wa.UIA_ComboBoxControlTypeId:  appElements.COMBOBOX,
	wa.UIA_DocumentControlTypeId:  appElements.DOCUMENT,
	wa.UIA_HyperlinkControlTypeId: appElements.HYPERLINK,
}

type AXElement struct {
	ref *wa.IUIAutomationElement
}

func GetForegroundRootAXElement() (axAPI.OSAXElement, error) {
	// Move this to cmd or functions
	util.Initialize()
	// // Initialize context
	// win32waf.Initalize()
	hwnd, err := win32wam.GetForegroundWindow()
	if err != nil {
		return nil, err
	}
	parentRef, err := win32waf.ElementFromHandle(hwnd)
	if err != nil {
		return nil, err
	}
	return AXElement{ref: parentRef}, nil
}

func (a AXElement) Press() error {
	err := wa.Invoke(a.ref)
	util.Finalize()
	// // Finalize context
	// win32waf.Finalize()
	return err
}

func (a AXElement) SetValue(value string) error {
	err := win32waf.SetElementValue(a.ref, value)
	util.Finalize()
	// // Finalize context
	// win32waf.Finalize()
	return err
}

func (a AXElement) SetValueOnFocus(value string) error {
	return fmt.Errorf("not implemented yet")
}

func (a AXElement) GetAllChildren() ([]axAPI.OSAXElement, error) {
	var children []axAPI.OSAXElement
	elements, err := win32waf.GetAllChildren(a.ref, maps.Keys(elementTypes))
	if err != nil {
		return nil, fmt.Errorf("error getting all children for element: %v", err)
	}
	childrenCount, err := elements.Get_Length()
	if err != nil {
		return nil, fmt.Errorf("error getting children count", err)
	}
	// logging.Debugf("we got %d children", childrenCount)
	var i int32
	for i = 0; i < childrenCount; i++ {
		if element, err := elements.GetElement(i); err == nil {
			children = append(children, AXElement{ref: element})
		}
	}
	return children, nil
}

func (a AXElement) GetType() (*appElements.ElementType, error) {
	iType, err := win32waf.GetCurrentControlType(a.ref)
	if err != nil {
		return nil, fmt.Errorf("error getting the type for an element: %v", err)
	}
	t, ok := elementTypes[int64(iType)]
	if ok {
		return &t, nil
	}
	return nil, fmt.Errorf("can not cast value to existing AX types, current internal value is %d", iType)
}

func (a AXElement) GetID() (string, error) {
	name, _ := a.ref.Get_CurrentName()
	if len(name) > 0 {
		return name, nil
	}
	return win32waf.GetElementText(a.ref)
}

func (a AXElement) GetValue() (string, error) {
	return "", fmt.Errorf("not implemented yet")
}
