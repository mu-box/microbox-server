package docker_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	dc "github.com/fsouza/go-dockerclient"
	"github.com/golang/mock/gomock"
	"github.com/jcelliott/lumber"

	"github.com/mu-box/microbox-server/config"
	"github.com/mu-box/microbox-server/util/docker"
	"github.com/mu-box/microbox-server/util/docker/mock_docker"
)

type createMatcher struct {
}

func (c createMatcher) Matches(x interface{}) bool {
	createConfig, ok := x.(dc.CreateContainerOptions)
	if !ok {
		return false
	}
	binds := []string{
		"/mnt/sda/var/microbox/cache/:/mnt/cache/",
		"/mnt/sda/var/microbox/deploy/:/mnt/deploy/",
		"/mnt/sda/var/microbox/build/:/mnt/build/",
		config.MountFolder + "code/app/:/share/code/:ro", // the app name cannot be grabbed outside the environment
		config.MountFolder + "engines/:/share/engines/:ro",
	}
	for i, bind := range createConfig.HostConfig.Binds {
		if binds[i] != bind {
			return false
		}
	}
	return true
}

func (c createMatcher) String() string {
	return "is a CreateContainerOptions"
}

func TestMain(m *testing.M) {
	config.Log = lumber.NewConsoleLogger(lumber.ERROR)
	if testing.Verbose() {
		config.Log = lumber.NewConsoleLogger(lumber.DEBUG)
	}

	curDir, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}
	dir := strings.Replace(curDir, "/util/docker", "/test/", 1)
	config.MountFolder = dir
	config.DockerMount = dir
	fmt.Println("appname: ", config.App())
	os.Exit(m.Run())
}

func TestCreatContainer(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mClient := mock_docker.NewMockClientInterface(ctrl)
	docker.Client = mClient

	gomock.InOrder(
		mClient.EXPECT().ListImages(dc.ListImagesOptions{}).Return([]dc.APIImages{dc.APIImages{RepoTags: []string{"mubox/build:latest"}}}, nil),
		mClient.EXPECT().CreateContainer(createMatcher{}).Return(&dc.Container{ID: "1234"}, nil),
		mClient.EXPECT().StartContainer("1234", nil),
		mClient.EXPECT().InspectContainer("1234"),
	)

	cc := docker.CreateConfig{
		Category: "build",
		UID:      "build1",
		Name:     "build",
		Cmd:      []string{"ls"},
		Image:    "mubox/build",
	}
	docker.CreateContainer(cc)

}

func TestExecInContainer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mClient := mock_docker.NewMockClientInterface(ctrl)
	docker.Client = mClient

	opts := dc.CreateExecOptions{
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          []string{"ls", "-la"},
		Container:    "exec1",
		User:         "root",
	}
	gomock.InOrder(
		mClient.EXPECT().CreateExec(opts).Return(&dc.Exec{ID: "1234"}, nil),
		mClient.EXPECT().StartExec("1234", gomock.Any()),
		mClient.EXPECT().InspectExec("1234").Return(&dc.ExecInspect{ExitCode: 0}, nil),
	)
	docker.ExecInContainer("exec1", "ls", "-la")

}

func TestListContainers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mClient := mock_docker.NewMockClientInterface(ctrl)
	docker.Client = mClient

	web := dc.APIContainers{ID: "1234", Labels: map[string]string{"web1": "true"}}
	db := dc.APIContainers{ID: "4321", Labels: map[string]string{"mysql1": "true"}}

	mClient.EXPECT().ListContainers(dc.ListContainersOptions{All: true, Size: false}).Times(2).Return([]dc.APIContainers{web, db}, nil)
	mClient.EXPECT().InspectContainer("1234").Times(2).Return(&dc.Container{ID: "1234"}, nil)
	mClient.EXPECT().InspectContainer("4321").Return(&dc.Container{ID: "4321"}, nil)

	results, err := docker.ListContainers()
	if err != nil || len(results) != 2 || results[0].ID != "1234" || results[1].ID != "4321" {
		t.Errorf("bad result from list containers")
	}
	results, err = docker.ListContainers("web1")

	if err != nil || len(results) != 1 || results[0].ID != "1234" {
		t.Errorf("bad result from list containers")
	}
}

func TestGetContainer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mClient := mock_docker.NewMockClientInterface(ctrl)
	docker.Client = mClient

	web := dc.APIContainers{ID: "1234", Labels: map[string]string{"web1": "true"}}
	db := dc.APIContainers{ID: "4321", Labels: map[string]string{"mysql1": "true"}}

	mClient.EXPECT().ListContainers(dc.ListContainersOptions{All: true, Size: false}).Return([]dc.APIContainers{web, db}, nil)
	mClient.EXPECT().InspectContainer("1234").Return(&dc.Container{ID: "1234"}, nil)
	mClient.EXPECT().InspectContainer("4321").Return(&dc.Container{ID: "4321"}, nil)

	result, err := docker.GetContainer("1234")
	if err != nil || result.ID != "1234" {
		t.Errorf("failed to retrieve container")
	}
}

func TestImageExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mClient := mock_docker.NewMockClientInterface(ctrl)
	docker.Client = mClient

	type APIImages struct {
		ID          string            `json:"Id" yaml:"Id"`
		RepoTags    []string          `json:"RepoTags,omitempty" yaml:"RepoTags,omitempty"`
		Created     int64             `json:"Created,omitempty" yaml:"Created,omitempty"`
		Size        int64             `json:"Size,omitempty" yaml:"Size,omitempty"`
		VirtualSize int64             `json:"VirtualSize,omitempty" yaml:"VirtualSize,omitempty"`
		ParentID    string            `json:"ParentId,omitempty" yaml:"ParentId,omitempty"`
		RepoDigests []string          `json:"RepoDigests,omitempty" yaml:"RepoDigests,omitempty"`
		Labels      map[string]string `json:"Labels,omitempty" yaml:"Labels,omitempty"`
	}
	base := dc.APIImages{RepoTags: []string{"mubox/base:alpha", "mubox/base:latest", "mubox/base:beta"}}
	redis := dc.APIImages{RepoTags: []string{"mubox/redis:3.4", "mubox/redis:latest", "mubox/base:3.4-stable"}}
	code := dc.APIImages{RepoTags: []string{"mubox/code:alpha", "mubox/code:latest", "mubox/code:beta"}}
	mClient.EXPECT().ListImages(dc.ListImagesOptions{}).AnyTimes().Return([]dc.APIImages{base, redis, code}, nil)

	working := []string{
		"mubox/base",
		"mubox/base:alpha",
		"mubox/base:latest",
		"mubox/base:beta",
		"mubox/redis:3.4",
		"mubox/redis:latest",
		"mubox/base:3.4-stable",
		"mubox/code:alpha",
		"mubox/code:latest",
		"mubox/code:beta",
	}
	for _, work := range working {
		if !docker.ImageExists(work) {
			t.Errorf("ImageExists couldnt find %s", work)
		}
	}

	notWorking := []string{
		"mubox/base:alphacentari",
		"mubox/bass:latest",
		"mubox/basestable",
		"mubox/redis:3.7",
		"mubox/redis:3.6-latest",
		"mubox/base:3.4-alpha",
		"cancer",
		"mubox/code:",
	}
	for _, noWork := range notWorking {
		if docker.ImageExists(noWork) {
			t.Errorf("ImageExists found %s", noWork)
		}
	}

}
