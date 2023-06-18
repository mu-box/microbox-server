package jobs_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mu-box/microbox-boxfile"

	dc "github.com/fsouza/go-dockerclient"
	"github.com/mu-box/microbox-server/jobs"
	"github.com/mu-box/microbox-server/util/docker"
	"github.com/mu-box/microbox-server/util/docker/mock_docker"

	"github.com/mu-box/microbox-server/util/fs"
	"github.com/mu-box/microbox-server/util/fs/mock_fs"

	"github.com/mu-box/microbox-server/util/script"
)

func TestDeployRemoveOldContainers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mDocker := mock_docker.NewMockDockerDefault(ctrl)
	docker.Default = mDocker

	mDocker.EXPECT().ListContainers("code", "build", "bootstrap", "dev", "tcp", "udp").Return([]*dc.Container{&dc.Container{ID: "1234", NetworkSettings: &dc.NetworkSettings{IPAddress: "1.2.3.4"}}}, nil)
	mDocker.EXPECT().RemoveContainer("1234")

	deploy := jobs.Deploy{}
	deploy.RemoveOldContainers()

}

func TestDeploySetupFs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mFs := mock_fs.NewMockFsUtil(ctrl)
	fs.FsDefault = mFs

	mFs.EXPECT().CreateDirs()
	mFs.EXPECT().Clean()

	deploy := jobs.Deploy{Reset: true}
	deploy.SetupFS()

}

func TestCreateBuildContainer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mDocker := mock_docker.NewMockDockerDefault(ctrl)
	docker.Default = mDocker

	gomock.InOrder(
		mDocker.EXPECT().ImageExists("mubox/build").Return(false),
		mDocker.EXPECT().InstallImage("mubox/build"),
		mDocker.EXPECT().CreateContainer(docker.CreateConfig{Image: "mubox/build", Category: "build", UID: "build1"}),
	)

	deploy := jobs.Deploy{}
	deploy.CreateBuildContainer(boxfile.Boxfile{})

}

func TestSetupBuild(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mFs := mock_fs.NewMockFsUtil(ctrl)
	fs.FsDefault = mFs

	mFs.EXPECT().UserPayload()

	names := []string{}
	script.Exec = func(name, container string, payload map[string]interface{}) ([]byte, error) {
		names = append(names, name)
		return []byte{}, nil
	}
	deploy := jobs.Deploy{}
	deploy.SetupBuild()
	expectedNames := []string{
		"default-user",
		"default-configure",
		"default-detect",
		"default-sync",
		"default-setup",
	}
	if len(names) != len(expectedNames) {
		t.Errorf("calls dont match the expected list of calls (%+v)", names)
		return
	}
	for i, name := range expectedNames {
		if names[i] != name {
			t.Errorf("I was expecting %s but got %s", name, names[i])
		}
	}
}

func TestRunBuild(t *testing.T) {
	names := []string{}
	script.Exec = func(name, container string, payload map[string]interface{}) ([]byte, error) {
		names = append(names, name)
		return []byte{}, nil
	}
	deploy := jobs.Deploy{Run: true}
	deploy.RunBuild()
	expectedNames := []string{
		"default-prepare",
		"default-build",
		"default-publish",
		"default-cleanup",
	}
	if len(names) != len(expectedNames) {
		t.Errorf("calls dont match the expected list of calls (%+v)", names)
		return
	}
	for i, name := range expectedNames {
		if names[i] != name {
			t.Errorf("I was expecting %s but got %s", name, names[i])
		}
	}
}

func TestRunDeployScripts(t *testing.T) {
	names := []string{}
	script.Exec = func(name, container string, payload map[string]interface{}) ([]byte, error) {
		names = append(names, name)
		return []byte{}, nil
	}
	deploy := jobs.Deploy{Run: true}
	deploy.RunDeployScripts("before", boxfile.New([]byte(`---
web1:
  before_deploy:
    - "php artisan migrate"
  before_deploy_all:
    - "php scripts/clear_cache.php"
`)))
	expectedNames := []string{
		"default-before_deploy",
	}
	if len(names) != len(expectedNames) {
		t.Errorf("calls dont match the expected list of calls (%+v)", names)
		return
	}
	for i, name := range expectedNames {
		if names[i] != name {
			t.Errorf("I was expecting %s but got %s", name, names[i])
		}
	}
}

func TestDefaultEVars(t *testing.T) {
	box := boxfile.New([]byte(`---
env:
  PORT: 3000
build:
  stability: beta
  engine: '../../../microbox-engine-golang'
`))
	r := jobs.DefaultEVars(box)
	if r["PORT"] != "3000" {
		t.Error("Numeric ports are not being processed correctly")
	}
}
