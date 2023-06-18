package jobs

import (
	"encoding/json"
	"strconv"

	"github.com/mu-box/microbox-golang-stylish"
	"github.com/mu-box/microbox-server/config"
	"github.com/mu-box/microbox-server/util"
	"github.com/mu-box/microbox-server/util/docker"
	"github.com/mu-box/microbox-server/util/script"
)

type ServiceEnv struct {
	EVars     map[string]string
	UID       string
	Success   bool
	FirstTime bool
}

func (j *ServiceEnv) Process() {
	j.Success = false

	// run environment hook (blocking)
	if out, err := script.Exec("environment", j.UID, nil); err != nil {
		util.HandleError(stylish.ErrorHead("Failed to configure %v's environment variables", j.UID))
		util.HandleError(stylish.ErrorBody(err.Error()))
		return
	} else {
		config.Log.Info("getting port data: %s", out)
		if err := json.Unmarshal(out, &j.EVars); err != nil {
			util.HandleError(stylish.ErrorHead("Failed to configure %v's environment variables", j.UID))
			util.HandleError(stylish.ErrorBody(err.Error()))
			return
		}
	}
	config.Log.Debug("getting port data: %+v", j.EVars)
	// if a service doesnt have a port we cant continue
	if j.EVars["PORT"] == "" {
		util.HandleError(stylish.ErrorHead("Failed to configure %v's tunnel", j.UID))
		util.HandleError(stylish.ErrorBody("no port given in environment"))
		return
	}

	// now we need to set the host in the evars as well as create a tunnel port in the router
	container, err := docker.InspectContainer(j.UID)
	if err != nil {
		util.HandleError(stylish.ErrorHead("Failed to configure %v's tunnel", j.UID))
		util.HandleError(stylish.ErrorBody(err.Error()))
	}
	config.Log.Debug("container: %+v", container)

	if j.FirstTime {
		j.EVars["HOST"] = container.NetworkSettings.IPAddress
		err = util.AddForward(j.EVars["PORT"], j.EVars["HOST"], j.EVars["PORT"])
		if err != nil {
			port, _ := strconv.Atoi(j.EVars["PORT"])
			for i := 1; i <= 10; i++ {
				err = util.AddForward(strconv.Itoa(port+i), j.EVars["HOST"], j.EVars["PORT"])
				if err == nil {
					break
				}
			}
			if err != nil {
				util.HandleError(stylish.Error("Failed to setup forward for service", err.Error()))
			}
		}
	}

	j.Success = true
}
