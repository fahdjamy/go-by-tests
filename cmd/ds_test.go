package main

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {

	verifySliceResponse := func(expected, received []int) {
		if reflect.DeepEqual(expected, received) != true {
			t.Errorf("got %v, want %v", received, expected)
		}
	}

	t.Run("sum all", func(t *testing.T) {
		received := SumAll([]int{1, 2, 3}, []int{3, 3, 3})
		expected := []int{6, 9}

		if reflect.DeepEqual(received, expected) != true {
			t.Errorf("got %v, want %v", received, expected)
		}
		verifySliceResponse(expected, received)
	})

	t.Run("zero sum if empty slices", func(t *testing.T) {
		received := SumAll([]int{}, []int{0, 0})
		expected := []int{0, 0}

		if reflect.DeepEqual(received, expected) != true {
			t.Errorf("got %v, want %v", received, expected)
		}
		verifySliceResponse(expected, received)
	})
}

func TestPerimeter(t *testing.T) {
	t.Run("perimeter of a rectangle calling calculateRectanglePerimeter", func(t *testing.T) {
		rectangle := Rectangle{
			Width:  10,
			Height: 10,
		}
		verifyFloatResponse(t, calculateRectanglePerimeter(rectangle), 40.0)
	})
}

func TestArea(t *testing.T) {
	t.Run("Area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{
			Width:  10,
			Height: 10,
		}
		verifyFloatResponse(t, CalculateArea(rectangle), 100.0)
	})

	t.Run("Area of circle when calling the Area method on a circle object", func(t *testing.T) {
		circle := Circle{
			Radius: 10,
		}

		verifyFloatResponse(t, circle.Area(), 314.1592653589793)
	})
}

func TestDictionary(t *testing.T) {

	t.Run("search word in dictionary", func(t *testing.T) {
		dict := Dictionary{"go": "go is awesome"}
		word := "go"
		expected := "go is awesome"

		assertDictionaryDescription(t, dict, word, expected)
	})

	t.Run("search word not in dictionary", func(t *testing.T) {
		dict := Dictionary{"go": "go is awesome"}
		word := "python"
		expected := ""
		result, err := dict.Search(word)

		assertDictError(t, err, ErrMissingWord)
		if result != expected {
			t.Errorf("got %v, want %v", result, expected)
		}
	})

	t.Run("add new word to dictionary", func(t *testing.T) {
		dict := Dictionary{}
		word := "python"
		description := "python is old"
		err := dict.Add(word, description)

		assertNoError(t, err)
	})

	t.Run("add existing word to dictionary", func(t *testing.T) {
		dict := Dictionary{"python": "python is mid"}
		word := "python"
		description := "python is old"
		err := dict.Add(word, description)

		assertDictError(t, err, ErrWordExists)
	})

	t.Run("update existing word to dictionary", func(t *testing.T) {
		dict := Dictionary{"python": "python is mid"}
		word := "python"
		description := "python is updated"

		ok := dict.Update(word, description)
		if !ok {
			t.Errorf("got %v, want %v", ok, true)
		}
	})

	t.Run("delete existing word to dictionary", func(t *testing.T) {
		dict := Dictionary{"python": "python is mid"}
		word := "python"

		dict.Delete(word)
		_, err := dict.Search(word)
		assertDictError(t, err, ErrMissingWord)
	})
}

func assertDictionaryDescription(t testing.TB, dict Dictionary, word, expectedDesc string) {
	t.Helper()
	resp, err := dict.Search(word)
	if resp != expectedDesc {
		t.Errorf("got %v, want %v", resp, expectedDesc)
	}
	assertNoError(t, err)
}

func assertDictError(t testing.TB, expected, received error) {
	t.Helper()
	if expected == nil {
		t.Fatal("expected an error to be return but was nil")
	}

	if expected.Error() != received.Error() {
		t.Errorf("got %v, want %v", received, expected)
	}
}

func verifyFloatResponse(t *testing.T, response, expected float64) {
	if response != expected {
		t.Errorf("got %f, want %f", response, expected)
	}
}
