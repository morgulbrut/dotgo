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

// Convert actually kicks of the conversation of the different files
func Convert(c config.Config) {
	for _, f := range c.Files {
		colorlog.Info("Deleting old %s", f.Outfile)
		os.Remove(f.Outfile)
		colorlog.Info("Converting %s to %s using %s", f.Infile, f.Outfile, f.Converter)
		if f.Type == "md2pdf" && f.Converter == "wkhtmltopdf" {
			md2html.Html(f)
			wkhtmlto.Pdf(f, c)
		}
		if f.Type == "md2html" && f.Converter == "golang-commonmark" {
			md2html.Html(f)
		}
		if f.Type == "md2pdf" && f.Converter == "pandoc" {
			pandoc.Pdf(f, c)
		}
		if f.Type == "md2latex" && f.Converter == "pandoc" {
			pandoc.LaTeX(f, c)
		}
		if f.Type == "md2revealjs" && f.Converter == "pandoc" {
			pandoc.Revealjs(f, c)
		}
		if f.Type == "md2beamer" && f.Converter == "pandoc" {
			pandoc.Beamer(f, c)
		}
	}
}

// ReadConfig reads the settings.toml file
func ReadConfig() config.Config {
	var conf config.Config
	_, err := toml.DecodeFile("settings.toml", &conf)
	if err != nil {
		colorlog.Fatal(err.Error())
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
