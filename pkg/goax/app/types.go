package app

import (
	"github.com/adrianriobo/goax/pkg/goax/app/api"
)

// representation on an app with a handler
// which is able interact within the app to run
// the operations defined by the interface
type App struct {
	handler api.AppHandler
}
