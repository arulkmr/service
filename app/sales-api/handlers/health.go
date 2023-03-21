package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// struct for logger
type check struct {
log *log.Logger
shutdown chan os.Signal

}

func (c check)health(ctx context.Context,w http.ResponseWriter, r *http.Request) error{

		status := struct{
			Status string
		}{
			Status: "Ok",
		}
		c.log.Println("logger Inside health func()")
		json.NewEncoder(w).Encode(status)
		return nil

	}