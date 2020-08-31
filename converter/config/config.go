package config

type Config struct {
	Paths Paths
	Files []File
}

type Paths struct {
	Pandoc      string
	Xelatex     string
	Wkhtmltopdf string
}

type File struct {
	Type      string
	Infile    string
	Outfile   string
	Converter string
	Title     string
	Subtitle  string
	PdfEngine string
	Abstract  string
	Keywords  []string
	Author    string
	Date      string
	Toc       bool
	Numbered  bool
	Variables []string
	Verbose   bool
	Codestyle string
	Latex     bool
	Header    string
	Replace   []string
}
