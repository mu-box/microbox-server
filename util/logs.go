package util

import (
	"fmt"

	"github.com/mu-box/microbox-router"
	"github.com/mu-box/microbox-server/config"
)

// LogDebug
func LogDebug(f string, v ...interface{}) {
	config.Logtap.Publish("deploy", 1, fmt.Sprintf(f, v...))
}

// LogInfo
func LogInfo(f string, v ...interface{}) {
	config.Logtap.Publish("deploy", 2, fmt.Sprintf(f, v...))
}

// LogWarn
func LogWarn(f string, v ...interface{}) {
	config.Logtap.Publish("deploy", 3, fmt.Sprintf(f, v...))
}

// LogError
func LogError(f string, v ...interface{}) {
	config.Logtap.Publish("deploy", 4, fmt.Sprintf(f, v...))
}

// LogFatal
func LogFatal(f string, v ...interface{}) {
	config.Logtap.Publish("deploy", 5, fmt.Sprintf(f, v...))
}

// HandleError
func HandleError(msg string) {
	LogDebug(msg)
	router.ErrorHandler = router.FailedDeploy{}
}
