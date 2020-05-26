package main

import (
	"flag"
	"os"

	"github.com/morgulbrut/dotgo/converter/config"

	"github.com/morgulbrut/colorlog"

	"github.com/morgulbrut/color256"
	"github.com/morgulbrut/dotgo/converter"
)

var Version string
var Build string

func main() {
	logo()

	configfile := flag.String("c", "settings.toml", "config file")
	init := flag.Bool("i", false, "generate new settings.toml")
	flag.Parse()
	if *init {
		config.Init()
		os.Exit(0)
	}
	colorlog.SetLogLevel(colorlog.TRACE)
	conf := converter.ReadConfig(*configfile)
	converter.Convert(conf)
}

func logo() {
	logo := `
▓█████▄  ▒█████  ▄▄▄█████▓  ▄████  ▒█████
▒██▀ ██▌▒██▒  ██▒▓  ██▒ ▓▒ ██▒ ▀█▒▒██▒  ██▒
░██   █▌▒██░  ██▒▒ ▓██░ ▒░▒██░▄▄▄░▒██░  ██▒
░▓█▄   ▌▒██   ██░░ ▓██▓ ░ ░▓█  ██▓▒██   ██░
░▒████▓ ░ ████▓▒░  ▒██▒ ░ ░▒▓███▀▒░ ████▓▒░
▒▒▓  ▒ ░ ▒░▒░▒░   ▒ ░░    ░▒   ▒ ░ ▒░▒░▒░
 ░ ▒  ▒   ░ ▒ ▒░     ░      ░   ░   ░ ▒ ▒░
 ░ ░  ░ ░ ░ ░ ▒    ░      ░ ░   ░ ░ ░ ░ ▒
   ░        ░ ░                 ░     ░ ░
 ░                                         `

	color256.PrintHiRed("%s", logo)
	color256.PrintGreen("Document Translator in Go. Ver: %s-%s ", Version, Build)
}
