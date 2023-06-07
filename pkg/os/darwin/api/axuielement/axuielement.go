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
	"github.com/adrianriobo/goax/pkg/util/logging"
)

type AXUIElementRef struct {
	ref      util.Ref
	id       string
	value    string
	role     string
	Parent   *AXUIElementRef
	children []*AXUIElementRef
}

func (e *AXUIElementRef) GetID() string {
	return e.id
}

func (e *AXUIElementRef) GetRef() util.Ref {
	return e.ref
}

func (e *AXUIElementRef) Role() string {
	return e.role
}

func (e *AXUIElementRef) Children() []*AXUIElementRef {
	return e.children
}

func (a *AXUIElementRef) Print() {
	logging.Debugf("type %s with name %s", a.role, a.id)
	for _, child := range a.children {
		child.Print()
	}
}

func GetAXUIElementRef(ref util.Ref, parent *AXUIElementRef) (*AXUIElementRef, error) {
	r := AXUIElementRef{}
	r.ref = ref
	hasChildren := hasChildren(ref)
	if hasChildren {
		r.children = getChildren(ref, &r)
	}
	id, _ := getID(ref)
	r.id = id
	r.role = getRole(ref)
	if parent != nil {
		r.Parent = parent
	}
	return &r, nil
}

func (e *AXUIElementRef) FindElementByRoleAndID(role, id string) (*AXUIElementRef, error) {
	// fmt.Println("element on headers", e.Role(), e.GetID(), len(e.Children()))
	if e.role == role && strings.Contains(e.id, id) {
		return e, nil
	}
	for _, child := range e.children {
		if element, err := child.FindElementByRoleAndID(role, id); err == nil {
			return element, nil
		}
	}
	return nil, fmt.Errorf("element not found")
}

func (e *AXUIElementRef) FindElementsByRoleAndID(role, id string) ([]*AXUIElementRef, error) {
	var elements []*AXUIElementRef
	// fmt.Println("element", e.Role(), e.GetID(), len(e.Children()))
	if e.role == role && strings.Contains(e.id, id) {
		elements = append(elements, e)
	}
	for _, child := range e.children {
		if childElements, err := child.FindElementsByRoleAndID(role, id); err == nil {
			elements = append(elements, childElements...)
		}
	}
	return elements, nil
}

func (e *AXUIElementRef) FindElementByID(id string) (*AXUIElementRef, error) {
	if len(e.id) > 0 && strings.Contains(e.id, id) {
		return e, nil
	}
	for _, child := range e.children {
		if element, err := child.FindElementByID(id); err == nil {
			return element, nil
		}
	}
	return nil, fmt.Errorf("element not found")
}

func (e *AXUIElementRef) FindElementByRole(role string) (*AXUIElementRef, error) {
	if e.role == role {
		return e, nil
	}
	for _, child := range e.children {
		if element, err := child.FindElementByRole(role); err == nil {
			return element, nil
		}
	}
	return nil, fmt.Errorf("element not found")
}

func (e *AXUIElementRef) SetValue(value string) {
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	C.SetValue(C.CFTypeRef(e.ref), cValue)
}

func (e *AXUIElementRef) ShowElements() {
	// elements with empty id are parents holding children with ids or with children below
	if len(e.children) > 0 {
		fmt.Printf("Parent type %s with id %s and ref %v\n", e.role, e.id, e.ref)
		for _, child := range e.children {
			child.ShowElements()
		}
	} else {
		fmt.Printf("Element %s id %s \n", e.role, e.id)
	}
}

func (e *AXUIElementRef) Press() {
	// println(e.ref)
	// println(getRole(e.ref))
	C.Press(C.CFTypeRef(e.ref))
}

func hasChildren(ref util.Ref) bool {
	has, err := strconv.ParseBool(C.GoString(C.HasChildren(C.CFTypeRef(ref))))
	if err != nil {
		return false
	}
	return has
}

func getChildren(ref util.Ref, parent *AXUIElementRef) []*AXUIElementRef {
	var children []*AXUIElementRef
	childrenASCFArray := C.GetChildren(C.CFTypeRef(ref))
	count := C.CFArrayGetCount(childrenASCFArray)
	for i := 0; i < int(count); i++ {
		if child, err := GetAXUIElementRef(util.Ref(C.GetChild(childrenASCFArray, C.CFIndex(i))), parent); err == nil {
			children = append(children, child)
		}
	}
	return children
}

// TODO id should be only title or description
func getID(ref util.Ref) (string, error) {
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

func getRole(ref util.Ref) string {
	return C.GoString(C.GetRole(C.CFTypeRef(ref)))
}

func showActions(ref util.Ref) {
	C.ShowActions(C.CFTypeRef(ref))
}
