package api

import (
	"net/http"

	"github.com/mu-box/microbox-server/jobs"
)

// CreateBuild
func (api *API) CreateBuild(rw http.ResponseWriter, req *http.Request) {

	//
	build := jobs.Build{
		ID:    newUUID(),
		Reset: (req.FormValue("reset") == "true"),
	}

	//
	api.Worker.QueueAndProcess(&build)

	//
	rw.Write([]byte("{\"id\":\"" + build.ID + "\", \"status\":\"created\"}"))
}
