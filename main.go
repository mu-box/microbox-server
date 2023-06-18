package main

import (
	"os"
	"time"

	mistServer "github.com/mu-box/mist/server"
	logapi "github.com/mu-box/microbox-logtap/api"
	"github.com/mu-box/microbox-logtap/archive"
	"github.com/mu-box/microbox-logtap/collector"
	"github.com/mu-box/microbox-logtap/drain"
	"github.com/mu-box/microbox-server/api"
	"github.com/mu-box/microbox-server/config"
)

func main() {
	// dont start until the app is populated
	for config.App() == "" {
		<-time.After(time.Second)
		config.Log.Info("waiting on app mount")
	}

	// start a mist TCP server listening at 0.0.0.0:1445
	mistServer.Start([]string{"tcp://0.0.0.0:1445"}, "")

	setupLogtap()

	// create new router
	err := router.StartHTTP(":" + config.Ports["router"])
	if err != nil {
		config.Log.Error("error: %s\n", err.Error())
	}

	// initialize the api and set up routing
	api := api.Init()

	// start microbox
	if err := api.Start(config.Ports["api"]); err != nil {
		config.Log.Fatal("[microbox/main.go] Unable to start API, aborting...\n%v\n", err)
		os.Exit(1)
	}
}

func setupLogtap() {
	//
	console := drain.AdaptLogger(config.Log)
	config.Logtap.AddDrain("console", console)

	// define logtap collectors/drains; we don't need to defer Close() anything here,
	// because these want to live as long as the server
	if _, err := collector.SyslogUDPStart("app", config.Ports["logtap"], config.Logtap); err != nil {
		panic(err)
	}

	//
	if _, err := collector.SyslogTCPStart("app", config.Ports["logtap"], config.Logtap); err != nil {
		panic(err)
	}

	// we will be adding a 0 to the end of the logtap port because we cant have 2 tcp listeneres
	// on the same port
	if _, err := collector.StartHttpCollector("deploy", config.Ports["logtap"]+"0", config.Logtap); err != nil {
		panic(err)
	}

	//
	db, err := archive.NewBoltArchive("/tmp/bolt.db")
	if err != nil {
		panic(err)
	}
	config.LogHandler = logapi.GenerateArchiveEndpoint(db)

	//
	config.Logtap.AddDrain("historical", db.Write)
	config.Logtap.AddDrain("mist", drain.AdaptPublisher(&mist.Proxy{}))

}
