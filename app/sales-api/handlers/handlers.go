package handlers

import (
	"log"
	"os"

	"github.com/arulkmr/service/foundation/web"
)
func API(build string, shutdown chan os.Signal, log *log.Logger)  *web.App {
	
	//tm :=httptreemux.NewContextMux()
	//Call from foundation 

	app:= web.NewApp(shutdown)
  
	// to pass as struct 
	c:= check{
		log:log,
		shutdown:shutdown,
	}

	// tm.Handle("GET","/health",c.health)
	app.Handle("GET","/health",c.health)
	return app
}
