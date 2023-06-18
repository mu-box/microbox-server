package jobs

import (
	"regexp"
	"strings"

	"github.com/mu-box/microbox-boxfile"
	"github.com/mu-box/microbox-golang-stylish"
	"github.com/mu-box/microbox-server/config"
	"github.com/mu-box/microbox-server/util"
	"github.com/mu-box/microbox-server/util/docker"
	"github.com/mu-box/microbox-server/util/script"
)

type ServiceStart struct {
	deploy Deploy

	Boxfile boxfile.Boxfile
	EVars   map[string]string
	Success bool
	UID     string
}

func (j *ServiceStart) Process() {
	// var ci *docker.Container
	var err error

	j.Success = false

	createConfig := docker.CreateConfig{UID: j.UID, Name: j.Boxfile.StringValue("name")}

	image := regexp.MustCompile(`\d+`).ReplaceAllString(j.UID, "")
	if image == "web" || image == "worker" || image == "tcp" || image == "udp" {
		createConfig.Category = "code"
		image = "code"
	} else {
		createConfig.Category = "service"
	}

	if j.Boxfile.StringValue("image") != "" {
		image = j.Boxfile.StringValue("image")
	} else {
		image = "mubox/" + image
	}

	extra := strings.Trim(strings.Join([]string{j.Boxfile.VersionValue("version"), j.Boxfile.StringValue("stability")}, "-"), "-")
	if extra == "" {
		image = image + ":latest"
	} else {
		image = image + ":" + extra
	}

	createConfig.Image = image

	if !docker.ImageExists(createConfig.Image) {
		util.LogInfo(stylish.SubBullet("- Pulling the %s image (this may take awhile)... ", createConfig.Image))
		docker.InstallImage(createConfig.Image)
	}

	util.LogDebug(stylish.SubBullet("- Image name: %v", createConfig.Image))

	util.LogInfo(stylish.SubBullet("- Creating %v container", j.UID))

	// start the container
	if _, err = docker.CreateContainer(createConfig); err != nil {
		util.HandleError(stylish.ErrorHead("Failed to create %v container", j.UID))
		util.HandleError(stylish.ErrorBody(err.Error()))
		util.UpdateStatus(&j.deploy, "errored")
		return
	}

	// payload
	payload := map[string]interface{}{
		"platform":    "local",
		"boxfile":     j.Boxfile.Parsed,
		"logtap_host": config.LogtapHost,
		"uid":         j.UID,
	}

	// adds to the payload storage information if storage is required
	needsStorage := false
	storage := map[string]map[string]string{}
	for key, val := range j.EVars {
		matched, _ := regexp.MatchString(`NFS\d+_HOST`, key)
		if matched {
			needsStorage = true
			nfsUid := strings.ToLower(regexp.MustCompile(`_HOST`).ReplaceAllString(key, ""))
			host := map[string]string{"host": val}
			storage[nfsUid] = host
		}
	}

	if needsStorage {
		payload["storage"] = storage
	}

	// run configure hook (blocking)
	if data, err := script.Exec("default-configure", j.UID, payload); err != nil {
		util.LogDebug("Failed Script Output:\n%s\n", data)
		util.HandleError(stylish.Error("Configure hook failed", err.Error()))
		util.UpdateStatus(&j.deploy, "errored")
		return
	}

	util.LogInfo(stylish.SubBullet("- Starting %v service", j.UID))

	// run start hook (blocking)
	if data, err := script.Exec("default-start", j.UID, payload); err != nil {
		util.LogDebug("Failed Script Output:\n%s\n", data)
		util.HandleError(stylish.Error("Start hook failed", err.Error()))
		util.UpdateStatus(&j.deploy, "errored")
		return
	}

	// if we make it to the end it was a success!
	j.Success = true

	util.LogDebug("   [√] SUCCESS\n")
}
