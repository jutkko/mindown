# mindown
[![Build Status](https://travis-ci.org/jutkko/mindown.svg?branch=master)](https://travis-ci.org/jutkko/mindown)

<img src="/figures/mindown.png" width="550">

## What is it?
mindown is the tool that maps what is in your mind to files! It is simple to
use and easy to extend for your needs!

## How to use it?
It is both an app and a library. Follow the instructions below to get your hands
on.

### App
Run

```
go install github.com/jutkko/mindown
mindown --input-file input.opml --output-file output.md
```

And boom! You get your markdown skeleton from your mindmap.

For further options use `mindown -h` to get you further.

### Library
See `main.go`. It should be pretty straight forward.

## How to extend it?
Right now we only support the following formats to input:

- [OPML](http://dev.opml.org/)

And for output:

- GitHub style [markdown](http://daringfireball.net/projects/markdown/)

### Input
### Output
## Todo
### Notes
### Checkbox
### Photos
## Project52
This is a project from my [Project52](https://github.com/jutkko/project52).
