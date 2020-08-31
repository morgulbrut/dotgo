# Features
## Implemented:

md => html:
:  using **golang-commonmark**, so no additional Software needed.

md => pdf:
:  using html and **wkhtmltopdf** (smaller than **LaTeX**), wkhtmltopdf needs to be installed

md => pdf:
: using **pandoc**. Pandoc and the tools for the pdf engine (see below) needs to be installed.

md => LaTeX:
: using **pandoc**

md => beamer:
: using **pandoc** and **XeLaTeX**

md => reveal.js presentation:
: using **pandoc**

# Install and configuration
## Configuration
### Paths

Configuration is done with `settings.toml`. It basically has two parts.

**paths** defines the paths for the tools it uses

```toml
[paths]
pandoc=""      
xelatex=""     
wkhtmltopdf="" 
```

**files** defines the which file needs to processed in which way, as an example:

```toml
[[files]]
title="DoTGo"
subtitle="Document Translator in Go"
author="morgulbrut"
date="\\today"
keywords=["documents", "go", "xelatex"]
abstract="Document Translator in G"
type="pdf"
infile="README.md"
outfile="readme_pandoc_xelatex.pdf"
pdfengine="xelatex" 
converter="pandoc" 
toc=true
numbered=true
verbose=true
codestyle="pygments"
variables=[
    "documentclass=scrartcl",
    "fontsize=10pt",
    "papersize=a4",
    "lang=de",
    "mainfont=Calibri.ttf"
]
```

### Options

Options for the settings file
`title`
: title of the generated pdf (when using **pandoc**).

`subtitle`
: subtitle of the generated pdf (when using **pandoc**).

`author`
: author of the generated pdf (when using **pandoc**).

`date`
: date of the generated pdf (when using **pandoc**).

`keywords`
: keywords of the generated pdf (when using **pandoc**).

`abstract`
: abstract of the generated pdf (when using **pandoc**).

`type`
: type of conversion, can be one of: **md2html**, **md2pdf**, **md2latex**, **md2beamer**, **md2revealjs**.

`infile`
: file name of the input file.

`outfile`
: file name of the outputfile file.

`pdfengine`
: pdf engine for **pandoc** can be **pdflatex**, **lualatex**, **xelatex**, **latexmk**, **tectonic**, **wkhtmltopdf**, **weasyprint**, **prince**, **context** or **pdfroff**. Default is pdflatex.

`converter`
: sets the converter, can be one of: **wkhtmltopdf**, **pandoc**, **golang-commonmark**.

`toc`
: sets wether or not to render a table of contents (when using **pandoc**).

`numbered`
: sets wether or not to number the sections (when using **pandoc**).

`verbose`
: sets wether or not to output messages from pandoc (when using **pandoc**).

`codestyle`
: sets the style for codelistings (when using **pandoc**). Can be one of: **pygments**, **kate**, **monochrome**, **espesso**, **haddock**, **zenburn**, **tango**.

`variables`
: variables to pass on to pandoc (when using **pandoc**).

`header`
: path to a custom header file, will be passed to pandoc (when using **pandoc**).

`replace`
: a list of strings, multiples of two, where the first gets replaces with the second (when using **pandoc** to make a pdf).

## Examples

Check the settings.toml in this Repo

# Markdown primer
## Sections

```
# H1
## H2
### H3
#### H4
##### H5
###### H6
```

Alternatively, for H1 and H2, an underline-ish style:

```md
Alt-H1
======

Alt-H2
------
```

## Textformat

Emphasis, aka italics, with *asterisks* or _underscores_.

Strong emphasis, aka bold, with **asterisks** or __underscores__.

Combined emphasis with **asterisks and _underscores_**.

Strikethrough uses two tildes. ~~Scratch this.~~

```
*italics* _italics_

**bold** __bold__

***bold italics*** **_bold italics_** 
__*bold italics*__ _**bold italics**_

~~strikethrough~~
```

## Lists
### Numbered lists 

1. First ordered list item
2. Another item
   * Unordered sub-list. 
1. Actual numbers don't matter, just that it's a number
   1. Ordered sub-list
4. And another item.

   You can have properly indented paragraphs within 
   list items. Notice the blank line above, and the 
   leading spaces (at least one, but we'll use three 
   here to also align the raw Markdown).

   To have a line break without a paragraph, you will 
   need to use two trailing spaces.Note that this line 
   is separate, but within the same paragraph. (This is 
   contrary to the typical GFM line break behaviour, 
   where trailing spaces are not required.)

```
1. First ordered list item
2. Another item
   * Unordered sub-list. 
1. Actual numbers don't matter, just that it's a number
   1. Ordered sub-list
4. And another item.

   You can have properly indented paragraphs within 
   list items. Notice the blank line above, and the 
   leading spaces (at least one, but we'll use three 
   here to also align the raw Markdown).

   To have a line break without a paragraph, you will 
   need to use two trailing spaces.Note that this line 
   is separate, but within the same paragraph. (This is 
   contrary to the typical GFM line break behaviour, 
   where trailing spaces are not required.)
```

### Bullet point lists

* Unordered list can use asterisks
- Or minuses
+ Or pluses

```
* Unordered list can use asterisks
- Or minuses
+ Or pluses
```

### Definition lists

def
: definition list

another def
: works in pandoc, therefor it works here.

```
def
: definition list

another def
: works in pandoc, therefor it works here.
```

## Links

[I'm an inline-style link](https://www.google.com)

[I'm an inline-style link with title](https://www.google.com "Google's Homepage")

[I'm a reference-style link][Arbitrary case-insensitive reference text]

[I'm a relative reference to a repository file](../blob/master/LICENSE)

[You can use numbers for reference-style link definitions][1]

Or leave it empty and use the [link text itself].

URLs and URLs in angle brackets will automatically get turned into links. 
http://www.example.com or <http://www.example.com> and sometimes 
example.com (but not on Github, for example).

Some text to show that the reference links can follow later.

[arbitrary case-insensitive reference text]: https://www.mozilla.org
[1]: http://slashdot.org
[link text itself]: http://www.reddit.com

```
[I'm an inline-style link](https://www.google.com)

[I'm an inline-style link with title](https://www.google.com "Google's Homepage")

[I'm a reference-style link][Arbitrary case-insensitive reference text]

[I'm a relative reference to a repository file](../blob/master/LICENSE)

[You can use numbers for reference-style link definitions][1]

Or leave it empty and use the [link text itself].

URLs and URLs in angle brackets will automatically get turned into links. 
http://www.example.com or <http://www.example.com> and sometimes 
example.com (but not on Github, for example).

Some text to show that the reference links can follow later.

[arbitrary case-insensitive reference text]: https://www.mozilla.org
[1]: http://slashdot.org
[link text itself]: http://www.reddit.com
```

## Images

Here's our logo (hover to see the title text):

Inline-style: 
![alt text](examplepics\640px-Bing_Hoele_Flower.jpg "Logo Title Text 1")

Reference-style: 
![alt text][logo]

[logo]:  examplepics\640px-Bing_Hoele_Flower.jpg "Logo Title Text 2"

```
Inline-style: 
![alt text](examplepics\640px-Bing_Hoele_Flower.jpg "Logo Title Text 1")

Reference-style: 
![alt text][logo]

[logo]: ./examplepics/204.jpg "Logo Title Text 2"
```

## Code

Inline `code` has `back-ticks around` it.

```javascript
var s = "JavaScript syntax highlighting";
alert(s);
```
 
```python
s = "Python syntax highlighting"
print s
```
 
```
No language indicated, so no syntax highlighting. 
But let's throw in a <b>tag</b>.
```

## Tables

Colons can be used to align columns.

| Tables        | Are           | Cool  |
| ------------- |:-------------:| -----:|
| col 3 is      | right-aligned | $1600 |
| col 2 is      | centered      |   $12 |
| zebra stripes | are neat      |    $1 |

There must be at least 3 dashes separating each header cell.
The outer pipes (|) are optional, and you don't need to make the 
raw Markdown line up prettily. You can also use inline Markdown.

Markdown | Less | Pretty
--- | --- | ---
*Still* | `renders` | **nicely**
1 | 2 | 3

## Quotes

> Blockquotes are very handy in email to emulate reply text.
> This line is part of the same quote.

Quote break.

> This is a very long line that will still be quoted properly when it wraps. Oh boy let's keep writing to make sure this is long enough to actually wrap for everyone. Oh, you can *put* **Markdown** into a blockquote. 


## HTML

<dl>
  <dt>Definition list</dt>
  <dd>Is something people use sometimes.</dd>

  <dt>Markdown in HTML</dt>
  <dd>Does *not* work **very** well. Use HTML <em>tags</em>.</dd>
</dl>


## LaTeX

**Pandoc** markdown supports inline LaTeX, that will be executed when building with any LaTeX engine
