//go:build windows

package ux

import (
	"fmt"
	"strings"

	win32waf "github.com/adrianriobo/goax/pkg/os/windows/api/user-interface/windows-accesibility-features"
	"github.com/adrianriobo/goax/pkg/util/logging"
	wa "github.com/openstandia/w32uiautomation"
	"golang.org/x/exp/maps"
)

func Initialize() {
	// Initialize context
	win32waf.Initalize()
}

func Finalize() {
	// Finalize context
	win32waf.Finalize()
}

func GetUXElementRef(ref *wa.IUIAutomationElement, parent *UXElement) (*UXElement, error) {
	r := UXElement{}
	r.Ref = ref
	children, err := getChildren(ref, parent)
	if err != nil {
		return nil, err
	}
	r.children = children
	id, err := getID(ref)
	if err != nil {
		return nil, err
	}
	r.name = id
	// elementType, err := win32waf.GetCurrentControlType(ref)
	// if err != nil {
	// 	return nil, err
	// }
	// r.elementType = elementType
	if parent != nil {
		r.Parent = parent
	}
	return &r, nil
}

func (e *UXElement) PressElementByID(id string) error {
	element, err := e.FindElementByID(id)
	if err != nil {
		return err
	}
	return element.Click()
}

func (e *UXElement) FindElementByID(id string) (*UXElement, error) {
	// fmt.Println("element on headers", e.Role(), e.GetID(), len(e.Children()))
	if strings.Contains(e.name, id) {
		return e, nil
	}
	for _, child := range e.children {
		if element, err := child.FindElementByID(id); err == nil {
			return element, nil
		}
	}
	return nil, fmt.Errorf("element not found")
}

func (u *UXElement) Print() {
	logging.Debugf("type %s with name %s", u.elementType, u.name)
	for _, child := range u.children {
		child.Print()
	}
}

func getID(ref *wa.IUIAutomationElement) (string, error) {
	// if id, err := ref.Get_CurrentAutomationId(); err == nil {
	// 	return id, nil
	// }
	name, _ := ref.Get_CurrentName()
	if len(name) > 0 {
		return name, nil
	}
	text, err := win32waf.GetElementText(ref)
	if err != nil {
		return "", nil
	}
	return text, nil
}

func GetActiveElement(name string, elementType string) (*UXElement, error) {
	logging.Debugf("Get %s: %s", elementType, name)
	if elementTypeId, ok := elementTypes[elementType]; !ok {
		return nil, fmt.Errorf("Error elementType %s is not supported", elementType)
	} else {
		if element, err := win32waf.GetActiveElement(name, elementTypeId); err != nil || element == nil {
			return nil, fmt.Errorf("Error getting element %s, with error %v", name, err)
		} else {
			return &UXElement{
				name:        name,
				elementType: elementType,
				Ref:         element}, nil
		}
	}
}

func GetActiveElementByType(elementType string) (*UXElement, error) {
	if elementTypeId, ok := elementTypes[elementType]; !ok {
		return nil, fmt.Errorf("Error elementType %s is not supported", elementType)
	} else {
		if element, err := win32waf.GetActiveElementByType(elementTypeId); err != nil || element == nil {
			return nil, fmt.Errorf("Error getting element %s, with error %v", elementType, err)
		} else {
			return &UXElement{
				name:        "",
				elementType: elementType,
				Ref:         element}, nil
		}
	}
}

func (u UXElement) GetName() string {
	return u.name
}

func (u UXElement) GetFullName() string {
	return fmt.Sprintf("%s: %s", u.elementType, u.name)
}

func (u UXElement) Click() error {
	logging.Debug("Click on %s", u.GetFullName())
	return wa.Invoke(u.Ref)
	// if u.elementType == CHECKBOX {
	// 	return wa.Invoke(u.Ref)
	// } else {
	// 	position, err := win32waf.GetElementRect(u.Ref)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return interaction.ClickOnRect(*position)
	// }
}

func (u UXElement) SetValue(value string) error {
	logging.Debug("setting value %s on %s", value, u.GetFullName())
	if u.elementType != EDIT {
		return fmt.Errorf("Error elementType %s is not supported", u.elementType)
	}
	return win32waf.SetElementValue(u.Ref, value)

}

func (u UXElement) GetElement(name string, elementType string) (*UXElement, error) {
	if elementTypeId, ok := elementTypes[elementType]; !ok {
		return nil, fmt.Errorf("Error elementType %s is not supported", elementType)
	} else {
		if element, err := win32waf.GetElementFromParent(u.Ref, name, elementTypeId); err != nil || element == nil {
			return nil, fmt.Errorf("%s not found on parent %s", elementType, u.GetFullName())
		} else {
			logging.Debugf("Get first %s on parent %s", elementType, u.GetFullName())
			return &UXElement{
				name:        name,
				elementType: elementType,
				Ref:         element}, nil
		}
	}
}

func (u UXElement) GetElementByType(elementType string) (*UXElement, error) {
	if elementTypeId, ok := elementTypes[elementType]; !ok {
		return nil, fmt.Errorf("Error elementType %s is not supported", elementType)
	} else {
		if element, err := win32waf.GetElementFromParentByType(u.Ref, elementTypeId); err != nil || element == nil {
			return nil, fmt.Errorf("%s not found on parent %s", elementType, u.GetFullName())
		} else {
			name, err := element.Get_CurrentName()
			if err != nil {
				logging.Error(err)
			}
			logging.Debugf("Get first %s on parent %s", elementType, u.GetFullName())
			return &UXElement{
				name:        name,
				elementType: elementType,
				Ref:         element}, nil
		}
	}
}

func (u UXElement) GetAllChildren(elementType string) ([]*UXElement, error) {
	logging.Debugf("Get all %s on parent %s ", elementType, u.name)
	if elementTypeId, ok := elementTypes[elementType]; !ok {
		return nil, fmt.Errorf("Error elementType %s is not supported", elementType)
	} else {
		var children []*UXElement
		elements, err := win32waf.GetAllChildren(u.Ref, []int64{elementTypeId})
		if err != nil {
			return nil, fmt.Errorf("Error getting %s on parent %s with error %v", elementType, u.GetFullName(), err)
		}
		childrenCount, err := elements.Get_Length()
		if err != nil {
			return nil, fmt.Errorf("Error getting %s on parent %s with error %v", elementType, u.GetFullName(), err)
		}
		var i int32
		for i = 0; i < childrenCount; i++ {
			if element, err := elements.GetElement(i); err == nil {
				children = append(children, &UXElement{
					elementType: elementType,
					Ref:         element})
			}
		}
		return children, nil
	}
}

func getChildren(ref *wa.IUIAutomationElement, parent *UXElement) ([]*UXElement, error) {
	var children []*UXElement
	elements, err := win32waf.GetAllChildren(ref, maps.Values(elementTypes))
	if err != nil {
		return nil, fmt.Errorf("error getting all children elements", err)
	}
	childrenCount, err := elements.Get_Length()
	if err != nil {
		return nil, fmt.Errorf("error getting children count", err)
	}
	var i int32
	for i = 0; i < childrenCount; i++ {
		if element, err := elements.GetElement(i); err == nil {
			if child, err := GetUXElementRef(element, parent); err == nil {
				children = append(children, child)
			}
		}
	}
	return children, nil
}
