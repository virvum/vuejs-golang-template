package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

var (
	// TODO export the build* variables to a separate module named "build".
	buildName      string
	buildVersion   string
	buildDate      string
	buildUserAgent string
	buildHomePage  string
	log            Log
	app            struct {
		start  time.Time
		status string
		config Config
	}
)

func main() {
	app.start = time.Now()
	app.status = "startup"
	log = LogInit(Trace, 32, true)

	configFile := flag.String("c", "config.yaml", "path to the configuration file")
	foreground := flag.Bool("f", false, "run in foreground (don't fork)")
	showVersion := flag.Bool("v", false, "show version and exit")
	flag.Parse()

	if *showVersion {
		fmt.Fprintf(os.Stderr, "%s %s compiled on %s with %v for %v/%v\n",
			buildName, buildVersion, buildDate, runtime.Version(), runtime.GOOS, runtime.GOARCH)
		return
	}

	if !*foreground {
		// TODO
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGHUP)

	for {
		switch <-sigchan {
		case os.Interrupt:
			log.Fatal("caught interrupt, shutting down")
		case syscall.SIGHUP:
			log.Info("caught SIGHUP, reloading configuration")

			if err := ConfigLoad(*configFile); err != nil {
				log.Error("ConfigLoad: %s", err)
				continue
			}

			// TODO reconnect to database etc.
		}
	}
}
