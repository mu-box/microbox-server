package jobs

import (
	"github.com/mu-box/microbox-golang-stylish"
	"github.com/mu-box/microbox-server/config"
	"github.com/mu-box/microbox-server/util"
	"github.com/mu-box/microbox-server/util/docker"
	"github.com/mu-box/microbox-server/util/fs"
	"github.com/mu-box/microbox-server/util/script"
)

type Bootstrap struct {
	ID     string
	Engine string
}

// Bootstrap the code according to the engine provided
func (j *Bootstrap) Process() {

	// Make sure we have the directories
	util.LogDebug(stylish.Bullet("Ensure directories exist on host..."))
	if err := fs.CreateDirs(); err != nil {
		util.HandleError(stylish.Error("Failed to create dirs", err.Error()))
		util.UpdateStatus(j, "errored")
		return
	}

	// if the build image doesn't exist it needs to be downloaded
	if !docker.ImageExists("mubox/build") {
		util.LogInfo(stylish.Bullet("Pulling the latest build image (this will take awhile)... "))
		docker.InstallImage("mubox/build")
	}

	// create a build container
	util.LogInfo(stylish.Bullet("Creating build container..."))
	_, err := docker.CreateContainer(docker.CreateConfig{Image: "mubox/build", Category: "bootstrap", UID: "bootstrap1"})
	if err != nil {
		util.HandleError(stylish.Error("Failed to create build container", err.Error()))
		util.UpdateStatus(j, "errored")
		return
	}

	// define the deploy payload
	payload := map[string]interface{}{
		"platform":    "local",
		"engine":      j.Engine,
		"logtap_host": config.LogtapHost,
	}

	// run configure hook (blocking)
	if _, err := script.Exec("default-bootstrap", "bootstrap1", payload); err != nil {
		util.HandleError(stylish.Error("Failed to run bootstrap hook", err.Error()))
		util.UpdateStatus(j, "errored")
	}

	docker.RemoveContainer("bootstrap1")

	util.UpdateStatus(j, "complete")
}
