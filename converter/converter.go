package converter

import (
	"os"

	"github.com/morgulbrut/colorlog"
	"github.com/morgulbrut/dotgo/converter/config"
	"github.com/morgulbrut/dotgo/converter/md2html"
	"github.com/morgulbrut/dotgo/converter/pandoc"
	"github.com/morgulbrut/dotgo/converter/wkhtmlto"
	"github.com/morgulbrut/toml"
)

const line string = "----------------------------------"

// Convert actually kicks off the conversation of the different files
func Convert(c config.Config) {
	for _, f := range c.Files {
		colorlog.Info("Deleting old %s", f.Outfile)
		os.Remove(f.Outfile)
		colorlog.Info("Converting %s to %s using %s", f.Infile, f.Outfile, f.Converter)
		if f.Type == "pdf" && f.Converter == "wkhtmltopdf" {
			md2html.Html(f)
			wkhtmlto.Pdf(f, c)
		}
		if f.Type == "html" && f.Converter == "golang-commonmark" {
			md2html.Html(f)
		}
		if f.Converter == "pandoc" {

			if f.Type == "pdf" {
				pandoc.Pdf(f)
			} else if f.Type == "beamer" {
				pandoc.Beamer(f)
			} else {
				pandoc.Generic(f)
			}
		}
	}
}

// ReadConfig reads the settings.toml file
func ReadConfig(conffile string) config.Config {
	var conf config.Config
	_, err := toml.DecodeFile(conffile, &conf)
	if err != nil {
		colorlog.Fatal("Config file not found.")
		colorlog.Fatal("Run dotgo with the -i flag to initalize a new config file.")
		colorlog.Fatal("Or run dotgo with the -c flag an a path to a config file.")
		os.Exit(1)
	}

	colorlog.Debug("Configuration:")
	colorlog.Debug("[paths]")
	colorlog.Debug("    pandoc:      %s", conf.Paths.Pandoc)
	colorlog.Debug("    xelatex:     %s", conf.Paths.Xelatex)
	colorlog.Debug("    wkhtmltopdf: %s", conf.Paths.Wkhtmltopdf)
	colorlog.Debug(line)

	colorlog.Debug("[files]")
	for _, f := range conf.Files {
		colorlog.Debug("    Type:        %s", f.Type)
		colorlog.Debug("    Infile:      %s", f.Infile)
		colorlog.Debug("    Outfile:     %s", f.Outfile)
		colorlog.Debug("    Converter:   %s", f.Converter)
		colorlog.Debug("    %s", line)
	}
	return conf
}
