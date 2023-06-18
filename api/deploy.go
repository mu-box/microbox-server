package api

import (
	"net/http"

	"github.com/mu-box/microbox-server/jobs"
)

// CreateDeploy
func (api *API) CreateDeploy(rw http.ResponseWriter, req *http.Request) {

	//
	deploy := jobs.Deploy{
		ID:    newUUID(),
		Reset: (req.FormValue("reset") == "true"),
		Run:   (req.FormValue("run") == "true"),
	}

	//
	api.Worker.QueueAndProcess(&deploy)

	//
	rw.Write([]byte("{\"id\":\"" + deploy.ID + "\", \"status\":\"created\"}"))
}
