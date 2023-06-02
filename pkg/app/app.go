package app

type app struct {
	appPath string
}

// Create app from app path
func New(appPath string) *app {
	return &app{
		appPath: appPath,
	}
}

func (a *app) Open() error {
	return osOpenApp(a.appPath)
}
