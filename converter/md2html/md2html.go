package md2html

// Copyright 2015 The Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// An example command-line tool that uses opennota/markdown to process markdown input.

import (
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/morgulbrut/colorlog"
	"github.com/morgulbrut/dotgo/converter/config"

	"gitlab.com/golang-commonmark/markdown"
)

var title string

func readFromFile(fn string) ([]byte, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}

func extractText(tok markdown.Token) string {
	switch tok := tok.(type) {
	case *markdown.Text:
		return tok.Content
	case *markdown.Inline:
		text := ""
		for _, tok := range tok.Children {
			text += extractText(tok)
		}
		return text
	}
	return ""
}

func writePreamble(w io.Writer) error {
	var opening string
	var ending string
	opening = `<!DOCTYPE html>
<html>`
	_, err := fmt.Fprintf(w, `%s
<head>
<meta charset="utf-8"%s>
<title>%s</title>
</head>
<body>
`, opening, ending, html.EscapeString(title))

	return err
}

func writePostamble(w io.Writer) error {
	_, err := fmt.Fprint(w, `</body>
</html>
`)
	return err
}

func Html(cf config.File) {
	colorlog.Info("[Md2html]: converting")
	var rendererOutput string
	data, err := readFromFile(cf.Infile)
	if err != nil {
		colorlog.Fatal("[Md2html] readFromFile(): %s", err.Error())
		os.Exit(1)
	}

	md := markdown.New(
		markdown.HTML(true),
		markdown.Tables(true),
		markdown.Linkify(true),
		markdown.Typographer(true),
		markdown.XHTMLOutput(false),
	)

	tokens := md.Parse(data)
	if len(tokens) > 0 {
		if heading, ok := tokens[0].(*markdown.HeadingOpen); ok {
			for i := 1; i < len(tokens); i++ {
				if tok, ok := tokens[i].(*markdown.HeadingClose); ok && tok.Lvl == heading.Lvl {
					break
				}
				title += extractText(tokens[i])
			}
			title = strings.TrimSpace(title)
		}
	}

	rendererOutput = md.RenderTokensToString(tokens)
	var outfile string
	if cf.Type == "md2pdf" && cf.Converter == "wkhtmltopdf" {
		outfile = "temp.html"
	} else {
		outfile = cf.Outfile
	}
	f, err := os.OpenFile(outfile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		colorlog.Fatal("[Md2html] OpenFile(): %s", err.Error())
		os.Exit(1)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			colorlog.Fatal("[Md2html] Close(): %s", err.Error())
			os.Exit(1)
		}
	}()

	err = writePreamble(f)
	if err != nil {
		colorlog.Fatal("[Md2html] writePreamble(): %s", err.Error())
		os.Exit(1)
	}

	_, err = f.WriteString(rendererOutput)
	if err != nil {
		colorlog.Fatal("[Md2html] WriteString(): %s", err.Error())
		os.Exit(1)
	}

	err = writePostamble(f)
	if err != nil {
		colorlog.Fatal("[Md2html] writePostamble(): %s", err.Error())
		os.Exit(1)
	}

}
