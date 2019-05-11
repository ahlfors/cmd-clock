package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lonord/cmd-clock/printer"
	"github.com/lonord/cmd-clock/weather"
)

var (
	appName    = "cmd-clock"
	appVersion = "dev"
	buildTime  = "unknow"
)

const timeInterval = time.Second
const weaInterval = time.Minute * 10

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
	var wea *weather.Weather
	var lastNow *time.Time
	// recover cursor
	defer os.Stdout.WriteString("\033[?25h\033c")
	// chan and timer
	weaCh := make(chan *weather.Weather)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan)
	timeTimer := time.NewTimer(timeInterval)
	weaTimer := time.NewTimer(weaInterval)
	// first action
	firstNow := time.Now()
	lastNow = &firstNow
	err := printer.Print(firstNow, wea)
	if err != nil {
		return err
	}
	updateWeather(weaCh)
	for {
		select {
		case w := <-weaCh:
			wea = w
			err := printer.Print(*lastNow, wea)
			if err != nil {
				fmt.Println(err)
			}
		case <-timeTimer.C:
			now := time.Now()
			if lastNow.Minute() != now.Minute() {
				err := printer.Print(now, wea)
				if err != nil {
					fmt.Println(err)
				}
				lastNow = &now
			}
			timeTimer.Reset(timeInterval)
		case <-weaTimer.C:
			updateWeather(weaCh)
			weaTimer.Reset(weaInterval)
		case sig := <-signalChan:
			switch sig {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
				return nil
			}
		}
	}
}

func updateWeather(ch chan *weather.Weather) {
	go func() {
		w, _ := weather.GetWeather()
		ch <- w
	}()
}
