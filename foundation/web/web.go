// Package web contains a small web framework extension.
package web

import (
	"context"
	"net/http"
	"os"
	"syscall"

	"github.com/dimfeld/httptreemux/v5"
)


type App struct {
	*httptreemux.ContextMux
	shutdown chan os.Signal
}
// A Handler is a type that handles an http request within our own little mini
// framework.
 type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error
// type Handler func(w http.ResponseWriter, r *http.Request) error


// NewApp creates an App value that handle a set of routes for the application.
func NewApp(shutdown chan os.Signal) *App {
	app := App{
		ContextMux: httptreemux.NewContextMux(),
		shutdown:   shutdown,
	}
	
	return &app
}

// Handle is our mechanism for mounting Handlers for a given HTTP verb and path
// pair, this makes for really easy, convenient routing.


//func (a *App) Handle(method string, path string, handler http.Handler) {

func (a *App) Handle(method string, path string, handler Handler) {


	// The function to execute for each request.
	h := func(w http.ResponseWriter, r *http.Request) {

		if err := handler(r.Context(),w,r); err != nil {
	      //  a.SignalShutdown()
			return
		}
	}

	// Add this handler for the specified verb and route.
	a.ContextMux.Handle(method, path, h)
}

// SignalShutdown is used to gracefully shutdown the app when an integrity
// issue is identified.
func (a *App) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}
