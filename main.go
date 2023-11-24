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
	appName    = "cmd-clock"
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
	// var
	var lastNow *time.Time
	// recover cursor
	defer os.Stdout.WriteString("\033[?25h\033c")
	// chan and timer
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan)
	timeTimer := time.NewTimer(timeInterval)
	// first action
	firstNow := time.Now()
	lastNow = &firstNow
	err := printer.Print(firstNow)
	if err != nil {
		return err
	}
	for {
		select {
		case <-timeTimer.C:
			now := time.Now()
			if lastNow.Minute() != now.Minute() {
				err := printer.Print(now)
				if err != nil {
					fmt.Println(err)
				}
				lastNow = &now
			}
			timeTimer.Reset(timeInterval)
		case sig := <-signalChan:
			switch sig {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
				return nil
			}
		}
	}
}
