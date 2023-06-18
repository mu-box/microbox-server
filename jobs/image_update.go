package jobs

import (
	"strings"

	stylish "github.com/mu-box/microbox-golang-stylish"
	"github.com/mu-box/microbox-server/util"
	"github.com/mu-box/microbox-server/util/docker"
)

type ImageUpdate struct{}

func (j *ImageUpdate) Process() {

	//
	images, err := docker.ListImages()
	if err != nil {
		util.HandleError("Unable to pull images:" + err.Error())
		util.UpdateStatus(j, "errored")
		return
	}

	//
	if len(images) == 0 {
		util.LogInfo(stylish.SubBullet("- No images available for update..."))
	}

	//
	for _, image := range images {
		for _, tag := range image.RepoTags {

			//
			if strings.HasPrefix(tag, "mubox") {
				util.LogInfo(stylish.SubBullet("- Updating image: %s", tag))
				if err := docker.InstallImage(tag); err != nil {
					util.HandleError("Unable to update image:" + err.Error())
					util.UpdateStatus(j, "errored")
					return
				}
			}
		}
	}

	util.LogInfo(stylish.SubBullet("- Update complete"))
	util.UpdateStatus(j, "complete")
}
