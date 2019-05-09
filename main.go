package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lonord/cmd-clock/weather"
)

var (
	appName    = "cmd-clock"
	appVersion = "dev"
	buildTime  = "unknow"
)

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
	w, err := weather.GetWeather()
	if err != nil {
		return err
	}
	fmt.Println(w)
	return nil
}
