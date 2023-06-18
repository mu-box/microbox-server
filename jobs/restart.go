package jobs

import (
	"github.com/mu-box/microbox-boxfile"
	"github.com/mu-box/microbox-golang-stylish"
	"github.com/mu-box/microbox-server/config"
	"github.com/mu-box/microbox-server/util"
	"github.com/mu-box/microbox-server/util/script"
)

type Restart struct {
	UID     string
	Success bool
	Boxfile boxfile.Boxfile
}

// Proccess syncronies your docker containers with the boxfile specification
func (j *Restart) Process() {
	// add a lock so the service wont go down whil im running
	util.Lock()
	defer util.Unlock()

	j.Success = false

	util.LogInfo(stylish.Bullet("Restarting app in %s container...", j.UID))
	box := CombinedBoxfile(false)
	// restart payload
	payload := map[string]interface{}{
		"platform":    "local",
		"boxfile":     box.Node(j.UID).Parsed,
		"logtap_host": config.LogtapHost,
		"uid":         j.UID,
	}

	// run restart hook (blocking)
	if _, err := script.Exec("default-restart", j.UID, payload); err != nil {
		util.LogInfo("ERROR %v\n", err)
		return
	}

	j.Success = true
}
