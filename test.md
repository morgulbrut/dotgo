---
title:  'This is the title: it contains a colon'
author:
- Author One
- Author Two
keywords: [nothing, nothingness]
abstract: |
  This is the abstract.

  It consists of two paragraphs.

documentclass: scrreprt
fontsize: 10pt
papersize: a4
---


Section 1
=========

Text can be *italicized* or **bolded** as well as `monospaced`. You can
\*escape certain\* special characters.

Subsection 1 (Level 2)
----------------------

Some section 2 text

### Sub-subsection 1 (level 3)

Some more text.

Examples
========

Comments
--------

Images
------

Add an image with:


Lists
-----

-   Bullet are made like this
-   Point levels must be consistent
    -   Sub-bullets
        -   Sub-sub-bullets
-   Lists

Term
:   Definition for term
Term2
:   Definition for term 2
:List of Things:
:   item1 - these are \'field lists\' not bulleted lists item2 item 3

Something
:   single item

Someitem
:   single item

Preformatted text
-----------------

A code example prefix must always end with double colon like it\'s
presenting something:

    Anything indented is part of the preformatted block

## Quote

> Until

> It gets back to

> Allll the way left

Code blocks
-----------

```py
import os
print(help(os))
```

Small Caps
---------

To write [Small caps]{.smallcaps} use the class `{.smallcaps}`

Links
-----

Web addresses by themselves will auto link, like this:
<https://www.devdungeon.com>

You can also inline custom links: [Google search
engine](https://www.google.com)

This is a simple [link](https://www.google.com) to Google with the link
defined separately.

This is a link to the [Python website](http://www.python.org/).

This is a link back to [Section 1](#section-1). You can link based off
of the heading name within a document.

Footnotes
---------

Footnote Reference[^1]

[^1]: This is footnote number one that would go at the bottom of the
    document.

Tables
------

  --------------------------
  Time     Number   Value
  -------- -------- --------
  12:00    42       2

  23:00    23       4
  --------------------------

Preserving line breaks
----------------------

Normally you can break the line in the middle of a paragraph and it will
ignore the newline. If you want to preserve the newlines, use the `|`
prefix on the lines. For example:

| These lines will
| break exactly
| where we told them to.

## Custom \LaTeX\ Commands

```Latex
\newcommand{\tuple}[1]{\langle #1 \rangle}

$\tuple{a, b, c}$
```

\newcommand{\tuple}[1]{\langle #1 \rangle}

$\tuple{a, b, c}$
