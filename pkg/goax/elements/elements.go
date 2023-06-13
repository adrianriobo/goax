package elements

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

// cast element type and check if it is supported
func GetElementType(elementType string) (*ElementType, error) {
	if len(elementType) == 0 {
		return nil, nil
	}
	t := ElementType(strings.TrimSpace(strings.ToLower(elementType)))
	if !slices.Contains(SupporterElementTypes, t) {
		return nil, fmt.Errorf("%s is not a supported element type", elementType)
	}
	return &t, nil
}
