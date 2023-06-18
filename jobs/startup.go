package jobs

//
import (
	"github.com/mu-box/microbox-server/config"
	"github.com/mu-box/microbox-server/util/docker"
	"github.com/mu-box/microbox-server/util/worker"
)

type Startup struct{}

// process on startup
func (j *Startup) Process() {
	config.Log.Info("starting startup job")

	docker.RemoveContainer("exec1")
	box := CombinedBoxfile(false)

	configureRoutes(*box)
	configurePorts(*box)

	// we also need to set up a ssh tunnel for each running docker container
	// this is easiest to do by creating a ServiceEnv job and working it
	worker := worker.New()
	worker.Blocking = true
	worker.Concurrent = true

	serviceContainers, _ := docker.ListContainers("service")
	for _, container := range serviceContainers {
		s := ServiceEnv{UID: container.Config.Labels["uid"], FirstTime: true}
		worker.Queue(&s)
	}

	worker.Process()
}
