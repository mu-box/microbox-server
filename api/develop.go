// Copyright (c) 2014 Pagoda Box Inc.
//
// This Source Code Form is subject to the terms of the Mozilla Public License,
// v. 2.0. If a copy of the MPL was not distributed with this file, You can
// obtain one at http://mozilla.org/MPL/2.0/.

package api

import (
	"net/http"
	"sync"

	"github.com/nanobox-io/nanobox-boxfile"
	"github.com/nanobox-io/nanobox-server/config"
	"github.com/nanobox-io/nanobox-server/util/docker"
	"github.com/nanobox-io/nanobox-server/util/fs"
	"github.com/nanobox-io/nanobox-server/util/script"
)

var developTex = sync.Mutex{}

func (api *API) Develop(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseMultipartForm(32 << 20)
	if err != nil {
		config.Log.Debug("form parsing error: \n %s", err.Error())
	}
	// force the develop route to go into a dev1 container
	req.Form["container"] = []string{"dev1"}

	box := combinedBox()

	containerControl := false

	developTex.Lock()
	// if there is no dev1 it needs to be created and this thread needs to remember
	// to shut it down when its done conatinerControl is used for that purpose
	container, err := docker.GetContainer("dev1")
	if err != nil || !container.State.Running {
		if container != nil && !container.State.Running {
			docker.RemoveContainer(container.ID)
		}
		containerControl = true
		cmd := []string{"/bin/sleep", "365d"}

		image := "nanobox/build"
		if stab := box.Node("build").StringValue("stability"); stab != "" {
			image = image + ":" + stab
		}

		container, err = docker.CreateContainer(docker.CreateConfig{Image: image, Category: "dev", UID: "dev1", Cmd: cmd})
		if err != nil {
			config.Log.Debug("develop create containter: %s", err.Error())
			rw.Write([]byte(err.Error()))
			return
		}

		// run the default-user hook to get ssh keys setup
		out, err := script.Exec("default-user", "dev1", fs.UserPayload())
		if err != nil {
			config.Log.Debug("Failed script output: \n %s", out)
			config.Log.Debug("out: %s", string(out))
		}
	}

	developTex.Unlock()

	api.Exec(rw, req)

	if containerControl {
		execWait.Wait()
		docker.RemoveContainer(container.ID)
	}
}

func combinedBox() boxfile.Boxfile {
	box := boxfile.NewFromPath(config.MountFolder + "code/" + config.App + "/Boxfile")

	if !box.Node("build").BoolValue("disable_engine_boxfile") {
		if out, err := script.Exec("default-boxfile", "build1", nil); err == nil {
			box.Merge(boxfile.New([]byte(out)))
		}
	}
	return box
}
