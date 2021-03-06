package pandoc

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

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
	if cf.Header != "" {
		arguments = append(arguments, "-H")
		arguments = append(arguments, cf.Header)
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
	for _, f := range cf.Filter {
		arguments = append(arguments, "--filter")
		arguments = append(arguments, f)
	}
	for _, v := range cf.Variables {
		arguments = append(arguments, "-V")
		arguments = append(arguments, v)
	}
	return arguments
}

func Beamer(cf config.File) {
	colorlog.Info("[pandoc]: converting")
	arguments := []string{"-s", "-t", "beamer", cf.Infile, "-o", cf.Outfile}
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

func Generic(cf config.File) {
	colorlog.Info("[pandoc]: converting")
	arguments := []string{"-s", "-t", cf.Type, cf.Infile, "-o", cf.Outfile}
	arguments = append(arguments, addArguments(cf)...)

	colorlog.Info("[pandoc]: Arguments: %q", arguments)
	cmd := exec.Command("pandoc", arguments...)
	_, err := cmd.Output()
	if err != nil {
		colorlog.Fatal(err.Error())
	}
}

func Pdf(cf config.File) {
	colorlog.Info("[pandoc]: converting")
	infile := cf.Infile
	if len(cf.Replace) > 0 {
		colorlog.Info("Replacing arrows")
		symReplace(cf)
		infile = "temp.md"
	}
	arguments := []string{"-s", infile, "-o", cf.Outfile}
	arguments = append(arguments, "--pdf-engine="+cf.PdfEngine)
	arguments = append(arguments, addArguments(cf)...)
	colorlog.Info("[pandoc]: Arguments: %q", arguments)
	cmd := exec.Command("pandoc", arguments...)

	st, err := cmd.Output()
	if err != nil {
		colorlog.Fatal(err.Error())
	}
	colorlog.Trace(string(st))
	if len(cf.Replace) > 0 {
		os.Remove("temp.md")
	}
}

func symReplace(cf config.File) {
	input, err := ioutil.ReadFile(cf.Infile)
	if err != nil {
		colorlog.Fatal(err.Error())
	}
	text := string(input)

	replacer := strings.NewReplacer(cf.Replace...)

	output := replacer.Replace(text)

	err = ioutil.WriteFile("temp.md", []byte(output), 0644)
	if err != nil {
		colorlog.Fatal(err.Error())
	}
}
