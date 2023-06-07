package app

type app struct {
	appPath string
	handler appHandler
}

type appHandler interface {
	Click(elementID string) error
	Check(elementID string) error
}

// Create app from app path
func New(appPath string) *app {
	return &app{
		appPath: appPath,
	}
}

// Open the app and load if needed
func (a *app) Open(load bool) error {
	err := osOpen(a.appPath)
	if !load {
		return err
	}
	handler, err := osLoad()
	a.handler = handler
	return err
}

func (a *app) Click(buttonID string) error {
	return a.handler.Click(buttonID)
}

func (a *app) Check(buttonID string) error {
	return a.handler.Check(buttonID)
}
