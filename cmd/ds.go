package main

import (
	"math"
)

type DictionaryError string
type Dictionary map[string]string

const (
	ErrWordExists  = DictionaryError("word already exist")
	ErrMissingWord = DictionaryError("word does not exist")
)

func (e DictionaryError) Error() string {
	return string(e)
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return calculateRectanglePerimeter(r)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func CalculateArea(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

func calculateRectanglePerimeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Height) * 2
}

func (d Dictionary) Search(word string) (string, error) {
	description, ok := d[word]
	if !ok {
		return "", ErrMissingWord
	}
	return description, nil
}

func (d Dictionary) Add(word, description string) error {
	if _, ok := d[word]; ok {
		return ErrWordExists
	}
	d[word] = description
	return nil
}

func (d Dictionary) Update(word, description string) bool {
	if _, err := d.Search(word); err == nil {
		d[word] = description
		return true
	}
	return false
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func SumAll(slices ...[]int) []int {
	var allSum []int
	for _, arr := range slices {
		var sum int
		for _, val := range arr {
			sum += val
		}
		allSum = append(allSum, sum)
	}

	return allSum
}
