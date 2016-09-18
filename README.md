# mindown
[![Build Status](https://travis-ci.org/jutkko/mindown.svg?branch=master)](https://travis-ci.org/jutkko/mindown)

<img src="/figures/mindown.png" width="750">

## What is it?
mindown is the tool that maps what is in your mind to files! It is simple to
use and easy to extend for your needs!

## How to use it?
It is both an app and a library. Follow the instructions below to get your
hands on.

### App
Running

```
go install github.com/jutkko/mindown
mindown --input-file input.opml --output-file output.md
```

And boom! You get your markdown skeleton from your mindmap.

For further options use `mindown -h` to get you further.

### Library
See `main.go`. It should be pretty straight forward.

## How to extend it?
The core idea of this project is to make mindmaps more programmable, editable
and approachable. The direction of this can be from mindmaps, so we are able
to export them to various formats. The other direction of this can be to
mindmaps, so we can visualise different formatted documents. It can be the
table of contents of an article, a book or whatever you can think of.

To add a new type of input/output for mindown, it's possible to only implement
one direction: i.e., you don't have add them in pairs. This is achieved by
providing a centric interface [graph](https:github.com/jutkko/mindown/utils/parse.go).

Right now we only support the following format(s) to input:

- [OPML](http://dev.opml.org/)

And for output:

- GitHub style [markdown](http://daringfireball.net/projects/markdown/)

### Input

### Output
## Todo
I use [Mindnote](https://mindnode.com/) for visualising the mindmaps. There are
a few things I'd like to be included in the graph interface.
### Notes
### Checkbox
### Photos

## Project52
This is a project from my [Project52](https://github.com/jutkko/project52).
