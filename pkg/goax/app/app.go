package app

import (
	"fmt"

	"github.com/adrianriobo/goax/pkg/goax/app/api"
	ax "github.com/adrianriobo/goax/pkg/goax/axapp"
	"github.com/adrianriobo/goax/pkg/goax/elements"
)

func Open(appPath string) error {
	if err := osOpen(appPath); err != nil {
		return fmt.Errorf("error opening te application at path %s: %v", appPath, err)
	}
	return nil
}

func LoadForefrontApp() (*App, error) {
	handler, err := osLoad()
	if err != nil {
		return nil, err
	}
	return &App{handler: handler}, nil
}

func LoadAppByTypeAndTitle(appType, appTytle string) (*App, error) {
	handler, err := osLoadByAppTypeAndTitle(appType, appTytle)
	if err != nil {
		return nil, err
	}
	return &App{handler: handler}, nil
}

func (a *App) Reload() (*App, error) {
	handler, err := a.handler.Reload()
	if err != nil {
		return nil, err
	}
	return &App{handler: handler}, nil
}

func (a *App) Click(element, elementType string, strict bool) error {
	return a.ClickWithOrder(element, elementType, strict, 0)
}

func (a *App) ClickWithOrder(element, elementType string, strict bool, order int) error {
	et, err := elements.GetElementType(elementType)
	if err != nil {
		return fmt.Errorf("error running click function: %v", err)
	}
	return a.handler.ClickWithOrder(element, et, strict, int8(order))
}

func (a *App) SetValue(element, elementType string, strict bool, value string) error {
	return a.SetValueWithOrder(element, elementType, strict, 0, value)
}

func (a *App) SetValueWithOrder(element, elementType string, strict bool, order int, value string) error {
	et, err := elements.GetElementType(elementType)
	if err != nil {
		return fmt.Errorf("error running set value function: %v", err)
	}
	return a.handler.SetValueWithOrder(element, et, strict, int8(order), value)
}

func (a *App) SetValueOnFocus(value string) error {
	return a.handler.SetValueOnFocus(value)
}

// Check if an element exists within the app with the element id/value
func (a *App) Exists(element, elementType string, strict bool) (bool, error) {
	et, err := elements.GetElementType(elementType)
	if err != nil {
		return false, fmt.Errorf("esrror running exists function: %v", err)
	}
	return a.handler.Exists(element, et, strict)
}

func (a *App) Print(idFilter string, strict bool) {
	a.handler.Print(idFilter, strict)
}

func osLoad() (api.AppHandler, error) {
	//need to create an inspect a windows app
	return ax.GetAXApp()
}

func osLoadByAppTypeAndTitle(appType, appTytle string) (api.AppHandler, error) {
	//need to create an inspect a windows app
	return ax.GetAXAppByTypeAndTitle(appType, appTytle)
}
