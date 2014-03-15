package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/nsf/termbox-go"
)

type Point struct {
	X, Y int
}

var (
	x, y, w, h int
	autoRun    = false
)

const (
	empty = 0x0020
)

func setDimensions(width, height int) {
	w = width
	h = height
}

func print_tb(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func printf_tb(x, y int, fg, bg termbox.Attribute, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	print_tb(x, y, fg, bg, s)
}

func showCursor() {
	termbox.SetCursor(x, y)
}

func pressLeft() {
	if x > 0 {
		x = x - 1
	}

	showCursor()
}

func pressUp() {
	if y > 0 {
		y = y - 1
	}

	showCursor()
}

func pressRight() {
	if x < w {
		x = x + 1
	}

	showCursor()
}

func pressDown() {
	if y < h {
		y = y + 1
	}

	showCursor()
}

func getCell(x, y int) (termbox.Cell, error) {
	w, h := termbox.Size()

	if x < 0 || x >= w {
		return termbox.Cell{}, fmt.Errorf("ERROR")
	}

	if y < 0 || y >= h {
		return termbox.Cell{}, fmt.Errorf("ERROR")
	}

	return termbox.CellBuffer()[y*w+x], nil
}

func spawnCell(x, y int) {
	termbox.SetCell(x, y, empty, termbox.ColorDefault, termbox.ColorGreen)
}

func killCell(x, y int) {
	termbox.SetCell(x, y, empty, termbox.ColorDefault, termbox.ColorBlack)
}

func clearCell(x, y int) {
	termbox.SetCell(x, y, empty, termbox.ColorDefault, termbox.ColorBlack)
}

func toggleCell(x, y int) {
	showCursor()

	if isAlive(x, y) {
		killCell(x, y)
	} else {
		spawnCell(x, y)
	}

	termbox.HideCursor()
}

func main() {
	// initialize termbox
	err := termbox.Init()
	if err != nil {
		os.Exit(1)
	}
	defer termbox.Close()

	// setup termbox
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)

	// set initial values
	w, h = termbox.Size()

	// put cursor in the middle
	x = w / 2
	y = h / 2
	showCursor()

	// set fps
	fpsSleepTime := time.Duration(1000000/25) * time.Microsecond

	go func() {
		for {
			time.Sleep(fpsSleepTime)
			termbox.Flush()

			if autoRun {
				tick()
			}
		}
	}()

	eventChan := make(chan termbox.Event)
	go func() {
		for {
			event := termbox.PollEvent()
			eventChan <- event
		}
	}()

	// register signals to channel
	sigChan := make(chan os.Signal)

	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	func() {
		for {
			select {
			case event := <-eventChan:
				switch event.Type {
				case termbox.EventKey:
					switch event.Key {
					case termbox.KeyCtrlZ, termbox.KeyCtrlC, termbox.KeyEsc:
						return
					case termbox.KeySpace: // Step forward one generation
						autoRun = false
						tick()
					case termbox.KeyEnter: // Toggle auto run
						toggleAutoRun()
					case termbox.KeyArrowLeft:
						pressLeft()
					case termbox.KeyArrowUp:
						pressUp()
					case termbox.KeyArrowRight:
						pressRight()
					case termbox.KeyArrowDown:
						pressDown()
					}

					// Keyboard shortcuts
					switch event.Ch {
					case 'a':
						drawPattern(&acorn)
					case 'd':
						drawPattern(&dieHard)
					case 'g':
						drawPattern(&glider)
					case 'G':
						drawPattern(&gosperGliderGun)
					case 'h':
						pressLeft()
					case 'j':
						pressDown()
					case 'k':
						pressUp()
					case 'l':
						pressRight()
					case 'c':
						clearGrid()
					case 'x':
						toggleCell(x, y)
					case 'q':
						return
					}

				case termbox.EventResize: // set new terminal dimensions
					setDimensions(event.Width, event.Height)

				case termbox.EventMouse:
					x = event.MouseX
					y = event.MouseY

					toggleCell(x, y)

				case termbox.EventError: // quit
					log.Fatalf("Quitting because of termbox error: \n%s\n", event.Err)
				}

			case <-sigChan:
				return
			}
		}
	}()
}

func isAlive(x, y int) bool {
	if x < 0 || y < 0 || x > w || y > h {
		return false
	}

	c, err := getCell(x, y)

	if err != nil {
		return false
	}

	return c.Bg == termbox.ColorGreen
}

func aliveNeighbours(x, y int) int {
	a := 0

	for ay := y - 1; ay <= y+1; ay++ {
		for ax := x - 1; ax <= x+1; ax++ {
			if !(ay == y && ax == x) {
				if isAlive(ax, ay) {
					a++
				}
			}
		}
	}

	return a
}

func tick() {
	aliveCells := 0

	var killList = make([]Point, 0)
	var spawnList = make([]Point, 0)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			n := aliveNeighbours(x, y)

			if isAlive(x, y) {
				aliveCells++

				if n < 2 || n > 3 {
					killList = append(killList, Point{x, y})
				}
			} else {
				if n == 3 {
					spawnList = append(spawnList, Point{x, y})
				}
			}
		}
	}

	// Stop the simulation when all cells are dead
	if aliveCells == 0 {
		autoRun = false
	}

	for _, pt := range killList {
		killCell(pt.X, pt.Y)
	}

	for _, pt := range spawnList {
		spawnCell(pt.X, pt.Y)
	}
}

func toggleAutoRun() {
	autoRun = !autoRun
}

func clearGrid() {
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
}

func drawPattern(pattern *[][]int) {
	termbox.HideCursor()

	for i, row := range *pattern {
		for j, pixel := range row {
			if pixel == 1 {
				spawnCell(x+j, y+i)
			} else {
				clearCell(x+j, y+i)
			}
		}
	}
}