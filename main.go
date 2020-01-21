package main

import (
	"github.com/morgulbrut/colorlog"
	"github.com/morgulbrut/dotgo/converter"

	"github.com/morgulbrut/color256"
)

func main() {
	logo()

	colorlog.SetLogLevel(colorlog.TRACE)
	config := converter.ReadConfig()
	converter.Convert(config)
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
	color256.PrintGreen("Document Translator in Go")
}
