package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lonord/cmd-clock/printer"
)

var (
	appVersion = "dev"
	buildTime  = "unknow"
)

const timeInterval = time.Second

func main() {
	ver := flag.Bool("version", false, "show version")
	flag.Parse()
	if *ver {
		fmt.Println("version", appVersion)
		fmt.Println("build time", buildTime)
		os.Exit(1)
	}
	err := run()
	if err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}
}

func run() error {
	defer resetCursor()

	sigChan, timeTimer := initializeChannels()

	currentTime := time.Now()

    if err := printer.Print(currentTime); err != nil {
        return err
    }

	lastTime := currentTime

	return manageEvents(sigChan, timeTimer, &lastTime)
}


// Reset cursor to default state
func resetCursor() {
    os.Stdout.WriteString("\033[?25h\033c")
}

// Initialize the channels for signals and timer
func initializeChannels() (chan os.Signal, *time.Timer) {
    signalChan := make(chan os.Signal, 1)
    signal.Notify(signalChan)

    return signalChan, time.NewTimer(timeInterval)
}

// Manage the incoming events from timer and signal channels
func manageEvents(sigChan chan os.Signal, timeTimer *time.Timer, lastTime *time.Time) error {
    for {
        select {
        case <-timeTimer.C:
            currentTime := time.Now()

            if lastTime.Minute() != currentTime.Minute() {
                if err := printer.Print(currentTime); err != nil {
                    fmt.Println(err)
                }

                lastTime = &currentTime
            }

            timeTimer.Reset(timeInterval)
        case sig := <-sigChan:
            if sig == syscall.SIGHUP || sig == syscall.SIGINT || sig == syscall.SIGQUIT || sig == syscall.SIGTERM {
                return nil
            }
        }
    }
}
