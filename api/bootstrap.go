package api

import (
	"net/http"

	"github.com/mu-box/microbox-server/jobs"
)

// CreateDeploy
func (api *API) CreateBootstrap(rw http.ResponseWriter, req *http.Request) {

	//
	bootstrap := jobs.Bootstrap{
		ID:     newUUID(),
		Engine: req.FormValue("engine"),
	}

	//
	api.Worker.QueueAndProcess(&bootstrap)

	//
	rw.Write([]byte("{\"id\":\"" + bootstrap.ID + "\", \"status\":\"created\"}"))
}
