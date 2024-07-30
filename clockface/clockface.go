package clockface

import (
	"math"
	"time"
)

/**
So we'll say that
- every clock has a centre of (150, 150)
- the hour hand is 50 long
- the minute hand is 80 long
- the second hand is 90 long.
*/

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

// A Point represents a two-dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

func secondsInRadians(t time.Time) float64 {
	// second is (2π / 60) radians
	// cancel out the 2, and we get π/30 radians
	// Multiply that by the number of seconds (as a float64)
	// return (math.Pi / 30) * float64(t.Second())

	//Floating point arithmetic is notoriously inaccurate.
	//Computers can only really handle integers, and rational numbers to some extent.
	//Decimal numbers start to become inaccurate,
	//especially when we factor them up and down as we are in the secondsInRadians function
	return math.Pi / (secondsInHalfClock / (float64(t.Second())))
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / minutesInClock) +
		(math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / hoursInClock) +
		(math.Pi / (hoursInHalfClock / float64(t.Hour()%hoursInClock)))
}

func SecondHandPoint(t time.Time) Point {
	// we want to measure the angle from 12 o'clock which is the Y axis
	// instead of from the X axis which we would like measuring the angle between the second hand and 3 o'clock.
	return angleToPoint(secondsInRadians(t))
}

func MinuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func HourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
