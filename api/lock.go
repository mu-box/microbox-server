package api

import (
	"fmt"
	"net/http"

	"github.com/mu-box/microbox-server/util"
)

func (api *API) Suspend(rw http.ResponseWriter, req *http.Request) {
	if util.LockCount() <= 0 {
		return
	}

	writeBody(map[string]string{"error": fmt.Sprintf("Current lock count: %d", util.LockCount())}, rw, http.StatusNotAcceptable)
}

func (api *API) LockCount(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "%d", util.LockCount())
}

// keeps a lock open as long as the connection is established with
// my service
func (api *API) Lock(rw http.ResponseWriter, req *http.Request) {
	util.Lock()
	defer util.Unlock()

	cNotify := rw.(http.CloseNotifier)
	<-cNotify.CloseNotify()
}
