package tracer

import (
	"fmt"
	"strconv"
	"strings"
)

// PPMFormat is the default PPM format of a canvas.
const PPMFormat = "P3"

// PPMMaxColorValue is the default PPM max color value of a canvas.
const PPMMaxColorValue = 255

// PPMMaxCharacterCount is the default PPM max character count of a row of a canvas.
const PPMMaxCharacterCount = 70

// canvas represents a rectangular grid of pixels.
type Canvas [][]Tuple

// NewCanvas creates a new canvas of a specified length and height,
// initialized with black pixels.
func NewCanvas(width, height int) Canvas {
	out := make([][]Tuple, height)
	for y := range out {
		out[y] = make([]Tuple, width)

		for x := range out[y] {
			out[y][x] = Color(0, 0, 0)
		}
	}
	return out
}

// width is the width of a canvas.
func (c Canvas) width() int {
	return len(c[0])
}

// height is the height of a canvas.
func (c Canvas) height() int {
	return len(c)
}

// WritePixel writes a pixel at a specfied x, y coordinate on the canvas.
func (c Canvas) WritePixel(x, y int, t Tuple) {
	c[y][x] = t
}

// ToPPM converts a canvas to a PPM formatted strng.
func (c Canvas) ToPPM() string {
	// create first three lines
	var b strings.Builder
	fmt.Fprintf(&b, "%s\n%d %d\n%d\n", PPMFormat, c.width(), c.height(), PPMMaxColorValue)

	count := 0
	for i, row := range c {
		for j, color := range row {
			// add color values to string
			for k, v := range color {
				// scale and cutoff values
				v *= PPMMaxColorValue
				if v < 0 {
					v = 0
				}
				if v > PPMMaxColorValue {
					v = PPMMaxColorValue
				}

				// stringify and add to builder
				s := strconv.Itoa(round(v))
				l := count + len(s)

				if l > PPMMaxCharacterCount {
					fmt.Fprintf(&b, "\n%s", s)
					count = len(s)

				} else if l == PPMMaxCharacterCount {
					fmt.Fprintf(&b, "%s\n", s)
					count = 0

				} else {
					fmt.Fprintf(&b, "%s", s)
					count += len(s)
				}

				// add new line or space if necessary
				if count >= PPMMaxCharacterCount-3 {
					fmt.Fprint(&b, "\n")
					count = 0

				} else if count != 0 && !(j == len(row)-1 && k == 2) {
					fmt.Fprint(&b, " ")
					count += 1
				}
			}
		}

		// start new row
		if i != len(c)-1 {
			fmt.Fprint(&b, "\n")
			count = 0
		}
	}

	return b.String()
}

// round rounds a float to the nearest integer value.
func round(f float64) int {
	i := int(f)
	if f-float64(i) >= 0.5 {
		return i + 1
	}

	return i
}
