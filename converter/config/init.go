package config

import (
	"io/ioutil"

	"github.com/morgulbrut/color256"
)

func Init() {
	template := `[[files]]
title="title"
subtitle="subtitle"
author="author"
date="\\today"
keywords=[""]
abstract=""
type="pdf"
infile="README.md"
outfile="Readme.pdf"
pdfengine="xelatex" # pdflatex, lualatex, xelatex, latexmk, tectonic, wkhtmltopdf, weasyprint, prince, context, and pdfroff
converter="pandoc"
toc=false
numbered=true
verbose=true
codestyle="pygments" # pygments, kate, monochrome, espesso, haddock, zenburn, tango
variables=[
	"documentclass=scrartcl",
	"fontsize=10pt",
	"papersize=a4",
	"lang=de",
	"mainfont=Calibri.ttf"
]
replace= [ # If you want to replace some substrings in your pdf output
		"<--", "$\\lefttarrow$",
		"-->", "$\\rightarrow$",
		"<->", "$\\leftrightarrow$",
		"<==", "$\\Lefttarrow$",
		"==>", "$\\Rightarrow$",
		"<=>", "$\\leftrightarrow$",
]
filter = [
	"pandoc-xnos"
]
`

	color256.PrintGreen("Generating new settings.toml")
	ioutil.WriteFile("settings.toml", []byte(template), 0644)
}
