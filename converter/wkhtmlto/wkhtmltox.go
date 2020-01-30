package wkhtmlto

import (
	"os"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/morgulbrut/colorlog"
	"github.com/morgulbrut/dotgo/converter/config"
)

func Pdf(cf config.File, c config.Config) {
	colorlog.Info("[wkhtmltopdf]: converting")
	wkhtmltopdf.SetPath(c.Paths.Wkhtmltopdf)
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		colorlog.Fatal("[%15s]: wkhtmltopdf.NewPDFGenerator(): %s", "Html2pdfWk", err.Error())
		os.Exit(1)
	}

	// Set global options
	pdfg.Dpi.Set(300)
	// TODO: expose settings
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.Grayscale.Set(false)
	pdfg.Title.Set(cf.Title)

	// Create a new input page from an URL

	var infile string
	if cf.Type == "pdf" && cf.Converter == "wkhtmltopdf" {
		infile = "temp.html"
	} else {
		infile = cf.Infile
	}

	page := wkhtmltopdf.NewPage(infile)

	// Set options for this page
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	page.Zoom.Set(0.95)
	page.UserStyleSheet.Set("res/css/github.css")
	// Add to document
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		colorlog.Fatal("[Html2pdfWk] pdfg.Create(): %s", err.Error())
		os.Exit(1)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile(cf.Outfile)
	if err != nil {
		colorlog.Fatal("[Html2pdfWk] pdfg.WriteFile(): %s", err.Error())
		os.Exit(1)
	}
}
