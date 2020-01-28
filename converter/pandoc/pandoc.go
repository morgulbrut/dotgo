package pandoc

import (
	"os/exec"

	"github.com/morgulbrut/colorlog"
	"github.com/morgulbrut/dotgo/converter/config"
)

func addArguments(cf config.File) []string {
	var arguments []string
	if cf.Toc {
		arguments = append(arguments, "--toc")
	}
	if cf.Codestyle != "" {
		arguments = append(arguments, "--highlight-style")
		arguments = append(arguments, cf.Codestyle)
	}
	if cf.Numbered {
		arguments = append(arguments, "-N")
	}
	if cf.Verbose {
		arguments = append(arguments, "--verbose")
	}
	if cf.Author != "" {
		arguments = append(arguments, "-V")
		arguments = append(arguments, "author="+cf.Author)
	}
	if cf.Title != "" {
		arguments = append(arguments, "-V")
		arguments = append(arguments, "title="+cf.Title)
	}
	if cf.Subtitle != "" {
		arguments = append(arguments, "-V")
		arguments = append(arguments, "subtitle="+cf.Subtitle)
	}
	if cf.Date != "" {
		arguments = append(arguments, "-V")
		arguments = append(arguments, "date="+cf.Date)
	}
	if cf.Abstract != "" {
		arguments = append(arguments, "-V")
		arguments = append(arguments, "abstract="+cf.Abstract)
	}
	for _, v := range cf.Variables {
		arguments = append(arguments, "-V")
		arguments = append(arguments, v)
	}
	return arguments
}

func Beamer(cf config.File, c config.Config) {
	colorlog.Info("[pandoc]: converting")
	arguments := []string{"-t", "beamer", cf.Infile, "-o", cf.Outfile}
	arguments = append(arguments, addArguments(cf)...)

	colorlog.Info("[pandoc]: Arguments: %q", arguments)
	cmd := exec.Command("pandoc", arguments...)
	_, err := cmd.Output()
	if err != nil {
		colorlog.Fatal(err.Error())
	}
}

func Revealjs(cf config.File, c config.Config) {
	colorlog.Info("[pandoc]: converting")
	arguments := []string{"-s", "-t", "revealjs", cf.Infile, "-o", cf.Outfile}
	arguments = append(arguments, addArguments(cf)...)

	colorlog.Info("[pandoc]: Arguments: %q", arguments)
	cmd := exec.Command("pandoc", arguments...)
	_, err := cmd.Output()
	if err != nil {
		colorlog.Fatal(err.Error())
	}
}

func LaTeX(cf config.File, c config.Config) {
	colorlog.Info("[pandoc]: converting")
	arguments := []string{"-s", cf.Infile, "-o", cf.Outfile}
	arguments = append(arguments, addArguments(cf)...)

	colorlog.Info("[pandoc]: Arguments: %q", arguments)
	cmd := exec.Command("pandoc", arguments...)
	_, err := cmd.Output()
	if err != nil {
		colorlog.Fatal(err.Error())
	}
}

func Pdf(cf config.File, c config.Config) {
	colorlog.Info("[pandoc]: converting")
	arguments := []string{cf.Infile, "-o", cf.Outfile}
	arguments = append(arguments, "--pdf-engine="+cf.PdfEngine)
	arguments = append(arguments, addArguments(cf)...)
	colorlog.Info("[pandoc]: Arguments: %q", arguments)
	cmd := exec.Command("pandoc", arguments...)

	st, err := cmd.Output()
	if err != nil {
		colorlog.Fatal(err.Error())
	}
	colorlog.Trace(string(st))
}
