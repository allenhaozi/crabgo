package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/go-logr/logr"
	log "github.com/sirupsen/logrus"
	ctrl "sigs.k8s.io/controller-runtime"
)

var logR logr.Logger

// {path}/{coredumpfilename}.txt
var coredumpFileFormat = "%s/openaios-iam-%s.txt"
var coredumpTimeFormat = "20060102150405"

func init() {
	logR = ctrl.Log.WithName("dump")
}

func StackTrace(all bool) string {
	buf := make([]byte, 10240)

	for {
		size := runtime.Stack(buf, all)

		if size == len(buf) {
			buf = make([]byte, len(buf)<<1)
			continue
		}
		break

	}

	return string(buf)
}

func InstallgoroutineDumpGenerator() {

	logR.Info("Register goroutine dump generator")

	signals := make(chan os.Signal, 1)

	signal.Notify(signals, syscall.SIGQUIT)

	go func() {
		for {
			sig := <-signals

			switch sig {
			case syscall.SIGQUIT:
				logR.Info("User uses kill -3 to generate goroutine dump")
				f := fileName()
				log.Info(f)
				coredump(f)
			// case syscall.SIGTERM:
			// 	fmt.Println("User told me to exit")
			// 	os.Exit(0)
			default:
				continue
			}
		}

	}()
}

func fileName() string {
	t := time.Now()
	timestamp := fmt.Sprint(t.Format(coredumpTimeFormat))

	return fmt.Sprintf(coredumpFileFormat, "/tmp", timestamp)

}

func coredump(fileName string) {
	logR.Info("Dump stacktrace to file", "fileName", fileName)
	trace := StackTrace(true)
	err := ioutil.WriteFile(fileName, []byte(trace), 0644)
	if err != nil {
		logR.Error(err, "Failed to write coredump.")
	}
	stdout := fmt.Sprintf("=== received SIGQUIT ===\n*** goroutine dump...\n%s", trace)
	logR.Info(stdout)

}
