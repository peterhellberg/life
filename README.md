life
====

Conway's Game of Life written in Go using [termbox-go](https://github.com/nsf/termbox-go)

![Gosper Glider Gun](http://assets.c7.se/skitch/glider_gun-20140315-004349.png)

## Installation

Just go get it:

		go get -u github.com/peterhellberg/life

## Usage

You can toggle cells (alive/dead) using the mouse and/or keyboard.
Press `ENTER` to start the simulation or `SPACE` to step forward one generation.

## Keyboard shortcuts

| Key     | Alternate key | Action               |
| -------:|:--------------|:-------------------- |
|  a      |               | draw acorn           |
|  d      |               | draw dieHard         |
|  g      |               | draw glider          |
|  G      |               | draw gosperGliderGun |
|  p      |               | draw rPentomino      |
|  h      | `LEFT ARROW`  | move left            |
|  j      | `DOWN ARROW`  | move down            |
|  k      | `UP ARROW`    | move up              |
|  l      | `RIGHT ARROW` | move right           |
|  c      |               | clear grid           |
|  x      | `LEFT CLICK`  | toggle cell          |
|  q      | `ESC`         | quit                 |
|  r      | `ENTER`       | toggle auto run      |
|  s      | `SPACE`       | step forward         |

## License

**The MIT License (MIT)**

Copyright (c) 2014 [Peter Hellberg](http://c7.se/)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
