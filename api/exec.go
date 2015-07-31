// Copyright (c) 2014 Pagoda Box Inc.
//
// This Source Code Form is subject to the terms of the Mozilla Public License,
// v. 2.0. If a copy of the MPL was not distributed with this file, You can
// obtain one at http://mozilla.org/MPL/2.0/.

package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/pagodabox/nanobox-boxfile"
	"github.com/pagodabox/nanobox-server/config"
	"github.com/pagodabox/nanobox-server/util"
)

func (api *API) Exec(rw http.ResponseWriter, req *http.Request) {
	util.RemoveContainer("exec1")
	conn, _, err := rw.(http.Hijacker).Hijack()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	defer conn.Close()

	cmd := []string{"/bin/bash"}
	if additionalCmd := req.FormValue("cmd"); additionalCmd != "" {
		cmd = append(cmd, "-c", additionalCmd)
	}

	container, err := util.CreateExecContainer("exec1", cmd)
	if err != nil {
		conn.Write([]byte(err.Error()))
		return
	}

	forwards := []string{}
	if req.FormValue("forward") != "" {
	  forwards = append(forwards, strings.Split(req.FormValue("forward"), ",")...)
	}

	box := mergedBox()
	if boxForwards, ok := box.Node("console").Value("forwards").([]interface{}); ok {
		for _, boxFInterface := range boxForwards {
			if boxForward, ok := boxFInterface.(string); ok {
				forwards = append(forwards, boxForward)
			}
		}
	}
	// maybe add a forward port mapping
	for _, rule := range forwards {
		strSlice := strings.Split(rule, ":")
		if len(strSlice) == 2 {
			portInt, _ := strconv.Atoi(strSlice[0])
			config.Router.AddForward("enter-"+rule, portInt, container.NetworkSettings.IPAddress+":"+strSlice[1])
			defer config.Router.RemoveForward("enter-" + rule)
		}
	}

	// Flush the options to make sure the client sets the raw mode
	conn.Write([]byte{})

	util.AttachToContainer(container.ID, conn, conn, conn)
	util.RemoveContainer(container.ID)
}

func (api *API) KillExec(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("signal recieved: %s\n", req.FormValue("signal"))
	err := util.KillContainer("exec1", req.FormValue("signal"))
	fmt.Println(err)
}

func (api *API) ResizeExec(rw http.ResponseWriter, req *http.Request) {
	h, _ := strconv.Atoi(req.FormValue("h"))
	w, _ := strconv.Atoi(req.FormValue("w"))
	if h == 0 || w == 0 {
		return
	}
	err := util.ResizeContainerTTY("exec1", h, w)
	fmt.Println(err)
}

func mergedBox() (box boxfile.Boxfile) {
	box = boxfile.NewFromPath("/vagrant/code/" + config.App + "/Boxfile")
	if out, err := util.ExecHook("boxfile", "build1", map[string]interface{}{}); err == nil {
		box.Merge(boxfile.New([]byte(out)))
	}
	return
}
