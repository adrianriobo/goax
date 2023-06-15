package ax

import (
	"fmt"
	"strings"

	axAPI "github.com/adrianriobo/goax/pkg/goax/axapp/api"
	"github.com/adrianriobo/goax/pkg/goax/elements"
	"github.com/adrianriobo/goax/pkg/util/logging"
	"golang.org/x/exp/slices"
)

type AXElement struct {
	Ref         axAPI.OSAXElement
	ID          string
	Value       string
	ElementType *elements.ElementType
	Parent      *AXElement
	Children    []*AXElement
}

func GetAXElement(element axAPI.OSAXElement, parent *AXElement) (*AXElement, error) {
	out := AXElement{
		Ref:    element,
		Parent: parent,
	}
	children, err := getChildren(element, &out)
	if err != nil {
		logging.Errorf("error getting children for ax element: %v", err)
	}
	out.Children = children
	elementType, err := element.GetType()
	if err != nil {
		// logging.Errorf("error getting type for ax element: %v", err)
		nonSupported := elements.NONSUPPORTED
		elementType = &nonSupported
	}
	out.ElementType = elementType
	id, err := element.GetID()
	if err != nil {
		// logging.Errorf("error getting id for ax element: %v", err)
		id = "unknown"
	}
	out.ID = id
	// value, err := element.GetValue()
	// if err != nil {
	// 	return nil, err
	// }
	// out.Value = value
	return &out, nil
}

// Prints the axelement, if hierarchy is set to true
// it will print the full hierarchy using the element as the root node
func (a *AXElement) Print(hierarchy bool, idFilter string, strict bool) {
	a.print(0, hierarchy, idFilter, strict)
}

// Get the elements by type and id
// if id should match extacly strict should be true
// the output is a map with the element and the order on the app
func (a *AXElement) FindElements(id string, elementTypes []elements.ElementType, strict bool) (out []*AXElement) {
	// Condition to match
	match := (strict && strings.TrimSpace(a.ID) == id) ||
		(!strict && strings.Contains(strings.TrimSpace(a.ID), id))
	if match &&
		(len(elementTypes) > 0 && slices.Contains(elementTypes, *a.ElementType) ||
			len(elementTypes) == 0) {
		out = append(out, a)
	}
	for _, child := range a.Children {
		out = append(out, child.FindElements(id, elementTypes, strict)...)
	}
	return
}

func (a *AXElement) print(nodeLevel int, hierarchy bool, idFilter string, strict bool) {
	if len(a.ID) > 0 {
		elementType := "unknown"
		if a.ElementType != nil {
			elementType = string(*a.ElementType)
		}
		if a.ID != "unknown" && (len(idFilter) == 0 ||
			(len(idFilter) > 0 && strict && strings.TrimSpace(a.ID) == idFilter) ||
			(len(idFilter) > 0 && !strict && strings.Contains(strings.TrimSpace(a.ID), idFilter))) {
			logging.Debugf("element at node %d is %s with id %s", nodeLevel, elementType, a.ID)
		}
	}
	if hierarchy {
		for _, child := range a.Children {
			child.print(nodeLevel+1, hierarchy, idFilter, strict)
		}
	}
}

func getChildren(element axAPI.OSAXElement, parent *AXElement) ([]*AXElement, error) {
	var children []*AXElement
	childElements, err := element.GetAllChildren()
	if err != nil {
		return nil, fmt.Errorf("error getting all children elements: %v", err)
	}
	for _, child := range childElements {
		if axElementchild, err := GetAXElement(child, parent); err == nil {
			children = append(children, axElementchild)
		}
	}
	return children, nil
}
