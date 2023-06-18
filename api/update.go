package api

import (
	"net/http"

	"github.com/mu-box/microbox-server/jobs"
)

// UpdateImages
func (api *API) UpdateImages(rw http.ResponseWriter, req *http.Request) {

	//
	api.Worker.QueueAndProcess(&jobs.ImageUpdate{})

	//
	rw.Write([]byte("{\"id\":\"1\", \"status\":\"created\"}"))
}
