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
	var lastNow *time.Time
	defer resetCursor()
	sigChan, timeTimer := initializeChannels()
    firstNow := time.Now()
	lastNow = &firstNow
	err := printer.Print(firstNow.Add(time.Hour * 8))
	if err != nil {
		return err
	}
	for {
		select {
		case <-timeTimer.C:
			now := time.Now()
			// now := currentTime.Add(time.Hour * 8)
			if lastNow.Minute() != now.Minute() {
				err := printer.Print(now.Add(time.Hour * 8))
				if err != nil {
					fmt.Println(err)
				}
				lastNow = &now
			}
			timeTimer.Reset(timeInterval)
		case sig := <-sigChan:
			switch sig {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
				return nil
			}
		}
	}
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
