package api

//
import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/pat"
	"github.com/pborman/uuid"

	"github.com/mu-box/microbox-server/config"
	"github.com/mu-box/microbox-server/jobs"
	"github.com/mu-box/microbox-server/util/worker"
)

// structs
type (
	API struct {
		Worker *worker.Worker
	}
)

func Init() *API {
	return &API{
		Worker: worker.New(),
	}
}

// Start
func (api *API) Start(port string) error {
	config.Log.Info("[microbox/api] Starting server...\n")

	//
	api.Worker.QueueAndProcess(&jobs.Startup{})

	//
	routes, err := api.registerRoutes()
	if err != nil {
		return err
	}

	//
	config.Log.Info("[microbox/api] Listening on port %v\n", port)

	// blocking...
	if err := http.ListenAndServe("0.0.0.0"+port, routes); err != nil {
		return err
	}

	return nil
}

// registerRoutes
func (api *API) registerRoutes() (*pat.Router, error) {
	config.Log.Info("[microbox/api] Registering routes...\n")

	//
	router := pat.New()

	//
	router.Get("/ping", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("pong"))
	})

	router.Get("/logs", config.LogHandler)

	router.Put("/suspend", api.handleRequest(api.Suspend))
	router.Put("/lock", api.handleRequest(api.Lock))
	router.Get("/lock-count", api.handleRequest(api.LockCount))

	router.Post("/develop", api.handleRequest(api.Develop))
	router.Post("/exec", api.handleRequest(api.Exec))

	router.Post("/console", api.handleRequest(api.Exec))
	router.Post("/resizeexec", api.handleRequest(api.ResizeExec))

	router.Get("/libdirs", api.handleRequest(api.LibDirs))
	router.Post("/file-change", api.handleRequest(api.FileChange))

	router.Post("/bootstrap", api.handleRequest(api.CreateBootstrap))
	router.Post("/builds", api.handleRequest(api.CreateBuild))
	router.Post("/deploys", api.handleRequest(api.CreateDeploy))
	router.Post("/image-update", api.handleRequest(api.UpdateImages))

	router.Get("/services", api.handleRequest(api.ListServices))
	router.Get("/routes", api.handleRequest(api.ListRoutes))
	router.Get("/vips", api.handleRequest(api.ListVips))
	return router, nil
}

// handleRequest
func (api *API) handleRequest(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {

		config.Log.Debug(`
Request:
--------------------------------------------------------------------------------
%+v

`, req)

		//
		fn(rw, req)

		config.Log.Debug(`
Response:
--------------------------------------------------------------------------------
%+v

`, rw)
	}
}

// newUUID
func newUUID() string {
	return uuid.New()
}

// parseBody
func parseBody(req *http.Request, v interface{}) error {

	//
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	defer req.Body.Close()

	//
	if err := json.Unmarshal(b, v); err != nil {
		return err
	}

	return nil
}

// writeBody
func writeBody(v interface{}, rw http.ResponseWriter, status int) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(b)

	return nil
}
