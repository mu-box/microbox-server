package config

import (
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/jcelliott/lumber"

	"github.com/mu-box/microbox-logtap"
)

var (
	app         string
	LogtapHost  string
	Ports       map[string]string
	IP          string
	MountFolder string
	DockerMount string
	CachedBox   string

	Log        lumber.Logger
	Logtap     *logtap.Logtap
	LogHandler http.HandlerFunc
)

func init() {
	MountFolder = "/vagrant/"
	DockerMount = "/mnt/"
	CachedBox = DockerMount + "sda/var/microbox/Boxfile.cache"
	// create an error object
	var err error
	levelEnv := os.Getenv("MICROBOX_LOGLEVEL")
	if levelEnv == "" {
		levelEnv = "INFO"
	}
	Log = lumber.NewConsoleLogger(lumber.LvlInt(levelEnv))

	//
	Ports = map[string]string{
		"api":    ":1757",
		"logtap": ":514",
		"router": "60000",
	}

	IP, err = externalIP()
	if err != nil {
		Log.Error("error: %s\n", err.Error())
	}

	LogtapHost = IP

	Logtap = logtap.New(Log)
}

func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			if strings.HasPrefix(ip.String(), "10") {
				continue
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

func App() string {
	if app != "" {
		return app
	}
	Log.Debug("appfolder: %s", MountFolder+"code/")

	files, err := ioutil.ReadDir(MountFolder + "code/")
	if err != nil {
		Log.Error(err.Error())
		return ""
	}

	if len(files) < 1 || !files[0].IsDir() {
		Log.Error("There is no code in your " + MountFolder + "code/ folder")
		return ""
	}
	app = files[0].Name()
	return app
}
