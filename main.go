package main

import (
	"os"

	"github.com/morgulbrut/colorlog"

	"github.com/morgulbrut/color256"
	"github.com/morgulbrut/dotgo/converter"
	"github.com/morgulbrut/dotgo/converter/config"
)

var Version string
var Build string

func main() {
	logo()

	if os.Args[1] == "init" {
		config.Init()
	} else {
		colorlog.SetLogLevel(colorlog.TRACE)
		config := converter.ReadConfig()
		converter.Convert(config)
	}
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
