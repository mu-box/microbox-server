// Automatically generated by MockGen. DO NOT EDIT!
// Source: docker.go

package mock_docker

import (
	go_dockerclient "github.com/fsouza/go-dockerclient"
	gomock "github.com/golang/mock/gomock"
	docker "github.com/mu-box/microbox-server/util/docker"
	io "io"
)

// Mock of ClientInterface interface
type MockClientInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockClientInterfaceRecorder
}

// Recorder for MockClientInterface (not exported)
type _MockClientInterfaceRecorder struct {
	mock *MockClientInterface
}

func NewMockClientInterface(ctrl *gomock.Controller) *MockClientInterface {
	mock := &MockClientInterface{ctrl: ctrl}
	mock.recorder = &_MockClientInterfaceRecorder{mock}
	return mock
}

func (_m *MockClientInterface) EXPECT() *_MockClientInterfaceRecorder {
	return _m.recorder
}

func (_m *MockClientInterface) ListImages(opts go_dockerclient.ListImagesOptions) ([]go_dockerclient.APIImages, error) {
	ret := _m.ctrl.Call(_m, "ListImages", opts)
	ret0, _ := ret[0].([]go_dockerclient.APIImages)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientInterfaceRecorder) ListImages(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListImages", arg0)
}

func (_m *MockClientInterface) PullImage(opts go_dockerclient.PullImageOptions, auth go_dockerclient.AuthConfiguration) error {
	ret := _m.ctrl.Call(_m, "PullImage", opts, auth)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientInterfaceRecorder) PullImage(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PullImage", arg0, arg1)
}

func (_m *MockClientInterface) CreateContainer(opts go_dockerclient.CreateContainerOptions) (*go_dockerclient.Container, error) {
	ret := _m.ctrl.Call(_m, "CreateContainer", opts)
	ret0, _ := ret[0].(*go_dockerclient.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientInterfaceRecorder) CreateContainer(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateContainer", arg0)
}

func (_m *MockClientInterface) StartContainer(id string, hostConfig *go_dockerclient.HostConfig) error {
	ret := _m.ctrl.Call(_m, "StartContainer", id, hostConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientInterfaceRecorder) StartContainer(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartContainer", arg0, arg1)
}

func (_m *MockClientInterface) KillContainer(opts go_dockerclient.KillContainerOptions) error {
	ret := _m.ctrl.Call(_m, "KillContainer", opts)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientInterfaceRecorder) KillContainer(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "KillContainer", arg0)
}

func (_m *MockClientInterface) ResizeContainerTTY(id string, height int, width int) error {
	ret := _m.ctrl.Call(_m, "ResizeContainerTTY", id, height, width)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientInterfaceRecorder) ResizeContainerTTY(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ResizeContainerTTY", arg0, arg1, arg2)
}

func (_m *MockClientInterface) StopContainer(id string, timeout uint) error {
	ret := _m.ctrl.Call(_m, "StopContainer", id, timeout)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientInterfaceRecorder) StopContainer(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StopContainer", arg0, arg1)
}

func (_m *MockClientInterface) RemoveContainer(opts go_dockerclient.RemoveContainerOptions) error {
	ret := _m.ctrl.Call(_m, "RemoveContainer", opts)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientInterfaceRecorder) RemoveContainer(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveContainer", arg0)
}

func (_m *MockClientInterface) WaitContainer(id string) (int, error) {
	ret := _m.ctrl.Call(_m, "WaitContainer", id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientInterfaceRecorder) WaitContainer(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WaitContainer", arg0)
}

func (_m *MockClientInterface) InspectContainer(id string) (*go_dockerclient.Container, error) {
	ret := _m.ctrl.Call(_m, "InspectContainer", id)
	ret0, _ := ret[0].(*go_dockerclient.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientInterfaceRecorder) InspectContainer(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InspectContainer", arg0)
}

func (_m *MockClientInterface) ListContainers(opts go_dockerclient.ListContainersOptions) ([]go_dockerclient.APIContainers, error) {
	ret := _m.ctrl.Call(_m, "ListContainers", opts)
	ret0, _ := ret[0].([]go_dockerclient.APIContainers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientInterfaceRecorder) ListContainers(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListContainers", arg0)
}

func (_m *MockClientInterface) CreateExec(opts go_dockerclient.CreateExecOptions) (*go_dockerclient.Exec, error) {
	ret := _m.ctrl.Call(_m, "CreateExec", opts)
	ret0, _ := ret[0].(*go_dockerclient.Exec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientInterfaceRecorder) CreateExec(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateExec", arg0)
}

func (_m *MockClientInterface) ResizeExecTTY(id string, height int, width int) error {
	ret := _m.ctrl.Call(_m, "ResizeExecTTY", id, height, width)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientInterfaceRecorder) ResizeExecTTY(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ResizeExecTTY", arg0, arg1, arg2)
}

func (_m *MockClientInterface) StartExec(id string, opts go_dockerclient.StartExecOptions) error {
	ret := _m.ctrl.Call(_m, "StartExec", id, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientInterfaceRecorder) StartExec(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartExec", arg0, arg1)
}

func (_m *MockClientInterface) InspectExec(id string) (*go_dockerclient.ExecInspect, error) {
	ret := _m.ctrl.Call(_m, "InspectExec", id)
	ret0, _ := ret[0].(*go_dockerclient.ExecInspect)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientInterfaceRecorder) InspectExec(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InspectExec", arg0)
}

// Mock of DockerDefault interface
type MockDockerDefault struct {
	ctrl     *gomock.Controller
	recorder *_MockDockerDefaultRecorder
}

// Recorder for MockDockerDefault (not exported)
type _MockDockerDefaultRecorder struct {
	mock *MockDockerDefault
}

func NewMockDockerDefault(ctrl *gomock.Controller) *MockDockerDefault {
	mock := &MockDockerDefault{ctrl: ctrl}
	mock.recorder = &_MockDockerDefaultRecorder{mock}
	return mock
}

func (_m *MockDockerDefault) EXPECT() *_MockDockerDefaultRecorder {
	return _m.recorder
}

func (_m *MockDockerDefault) CreateContainer(conf docker.CreateConfig) (*go_dockerclient.Container, error) {
	ret := _m.ctrl.Call(_m, "CreateContainer", conf)
	ret0, _ := ret[0].(*go_dockerclient.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDockerDefaultRecorder) CreateContainer(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateContainer", arg0)
}

func (_m *MockDockerDefault) StartContainer(id string) error {
	ret := _m.ctrl.Call(_m, "StartContainer", id)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerDefaultRecorder) StartContainer(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartContainer", arg0)
}

func (_m *MockDockerDefault) KillContainer(id string, sig string) error {
	ret := _m.ctrl.Call(_m, "KillContainer", id, sig)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerDefaultRecorder) KillContainer(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "KillContainer", arg0, arg1)
}

func (_m *MockDockerDefault) ResizeContainerTTY(id string, height int, width int) error {
	ret := _m.ctrl.Call(_m, "ResizeContainerTTY", id, height, width)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerDefaultRecorder) ResizeContainerTTY(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ResizeContainerTTY", arg0, arg1, arg2)
}

func (_m *MockDockerDefault) WaitContainer(id string) (int, error) {
	ret := _m.ctrl.Call(_m, "WaitContainer", id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDockerDefaultRecorder) WaitContainer(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WaitContainer", arg0)
}

func (_m *MockDockerDefault) RemoveContainer(id string) error {
	ret := _m.ctrl.Call(_m, "RemoveContainer", id)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerDefaultRecorder) RemoveContainer(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveContainer", arg0)
}

func (_m *MockDockerDefault) InspectContainer(id string) (*go_dockerclient.Container, error) {
	ret := _m.ctrl.Call(_m, "InspectContainer", id)
	ret0, _ := ret[0].(*go_dockerclient.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDockerDefaultRecorder) InspectContainer(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InspectContainer", arg0)
}

func (_m *MockDockerDefault) GetContainer(id string) (*go_dockerclient.Container, error) {
	ret := _m.ctrl.Call(_m, "GetContainer", id)
	ret0, _ := ret[0].(*go_dockerclient.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDockerDefaultRecorder) GetContainer(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetContainer", arg0)
}

func (_m *MockDockerDefault) ListContainers(labels ...string) ([]*go_dockerclient.Container, error) {
	_s := []interface{}{}
	for _, _x := range labels {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "ListContainers", _s...)
	ret0, _ := ret[0].([]*go_dockerclient.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDockerDefaultRecorder) ListContainers(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListContainers", arg0...)
}

func (_m *MockDockerDefault) InstallImage(image string) error {
	ret := _m.ctrl.Call(_m, "InstallImage", image)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerDefaultRecorder) InstallImage(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InstallImage", arg0)
}

func (_m *MockDockerDefault) ListImages() ([]go_dockerclient.APIImages, error) {
	ret := _m.ctrl.Call(_m, "ListImages")
	ret0, _ := ret[0].([]go_dockerclient.APIImages)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDockerDefaultRecorder) ListImages() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListImages")
}

func (_m *MockDockerDefault) ImageExists(name string) bool {
	ret := _m.ctrl.Call(_m, "ImageExists", name)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockDockerDefaultRecorder) ImageExists(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ImageExists", arg0)
}

func (_m *MockDockerDefault) ExecInContainer(container string, args ...string) ([]byte, error) {
	_s := []interface{}{container}
	for _, _x := range args {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "ExecInContainer", _s...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDockerDefaultRecorder) ExecInContainer(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ExecInContainer", _s...)
}

func (_m *MockDockerDefault) CreateExec(id string, cmd []string, in bool, out bool, err bool) (*go_dockerclient.Exec, error) {
	ret := _m.ctrl.Call(_m, "CreateExec", id, cmd, in, out, err)
	ret0, _ := ret[0].(*go_dockerclient.Exec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDockerDefaultRecorder) CreateExec(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateExec", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockDockerDefault) ResizeExecTTY(id string, height int, width int) error {
	ret := _m.ctrl.Call(_m, "ResizeExecTTY", id, height, width)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerDefaultRecorder) ResizeExecTTY(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ResizeExecTTY", arg0, arg1, arg2)
}

func (_m *MockDockerDefault) RunExec(exec *go_dockerclient.Exec, in io.Reader, out io.Writer, err io.Writer) (*go_dockerclient.ExecInspect, error) {
	ret := _m.ctrl.Call(_m, "RunExec", exec, in, out, err)
	ret0, _ := ret[0].(*go_dockerclient.ExecInspect)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDockerDefaultRecorder) RunExec(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RunExec", arg0, arg1, arg2, arg3)
}
