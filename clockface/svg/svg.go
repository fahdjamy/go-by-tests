package svg

import (
	"fmt"
	"go-by-tests/clockface"
	"io"
	"time"
)

const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 50
	clockCentreX     = 150
	clockCentreY     = 150
)

// Write writes an SVG representation of an analogue clock, showing the time t, to the writer w.
func Write(w io.Writer, tm time.Time) {
	_, _ = io.WriteString(w, svgStart)
	_, _ = io.WriteString(w, bezel)
	secondHand(w, tm)
	minuteHand(w, tm)
	hourHand(w, tm)
	_, _ = io.WriteString(w, svgEnd)
}

// secondHand is the unit vector of the second hand of an analogue clock at time `t`
// represented as a Point.
func secondHand(w io.Writer, t time.Time) {
	point := makeHand(clockface.SecondHandPoint(t), secondHandLength)
	_, _ = fmt.Fprintf(w,
		`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`,
		point.X, point.Y)
}

func minuteHand(w io.Writer, t time.Time) {
	point := makeHand(clockface.MinuteHandPoint(t), minuteHandLength)
	_, _ = fmt.Fprintf(w,
		`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`,
		point.X,
		point.Y)
}

func hourHand(w io.Writer, t time.Time) {
	point := makeHand(clockface.HourHandPoint(t), hourHandLength)
	_, _ = fmt.Fprintf(w,
		`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`,
		point.X,
		point.Y)
}

func makeHand(p clockface.Point, length float64) clockface.Point {
	p = clockface.Point{X: p.X * length, Y: p.Y * length}                // scale
	p = clockface.Point{X: p.X, Y: -p.Y}                                 // flip
	return clockface.Point{X: p.X + clockCentreX, Y: p.Y + clockCentreY} // translate
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
