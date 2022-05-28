package golang_united_school_homework

import (
	"math"
	"reflect"
	"strings"
	"testing"
)

const tooMuchError = "too much"
const notFoundError = "not found"

func TestAddShape(t *testing.T) {
	box := NewBox(2)

	shape := Circle{5}
	err := box.AddShape(shape)
	if !reflect.DeepEqual(err, nil) {
		t.Errorf("error, man 1")
	}

	shape = Circle{10}
	err = box.AddShape(shape)
	if !reflect.DeepEqual(err, nil) {
		t.Errorf("error, man 2")
	}

	shape = Circle{15}
	err = box.AddShape(shape)
	if err == nil || !strings.Contains(err.Error(), tooMuchError) {
		t.Errorf("error, man 3")
	}
}

func TestGetByIndex(t *testing.T) {
	box := NewBox(2)

	circle := Circle{5}
	err := box.AddShape(circle)
	if !reflect.DeepEqual(err, nil) {
		t.Errorf("error, man 1")
	}

	rectangle := Rectangle{2, 3}
	err = box.AddShape(rectangle)
	if !reflect.DeepEqual(err, nil) {
		t.Errorf("error, man 2")
	}

	shape, err := box.GetByIndex(0)
	if !reflect.DeepEqual(shape, circle) {
		t.Errorf("error, man 3")
	}

	shape, err = box.GetByIndex(1)
	if !reflect.DeepEqual(shape, rectangle) {
		t.Errorf("error, man 4")
	}

	shape, err = box.GetByIndex(2)
	if shape != nil || err == nil || !strings.Contains(err.Error(), notFoundError) {
		t.Errorf("error, man 5")
	}
}

func TestExtractByIndex(t *testing.T) {
	box := NewBox(3)

	circle := Circle{5}
	_ = box.AddShape(circle)

	rectangle := Rectangle{2, 3}
	_ = box.AddShape(rectangle)

	triangle := Triangle{4}
	_ = box.AddShape(triangle)

	shape, err := box.ExtractByIndex(1)
	if !reflect.DeepEqual(shape, rectangle) || !reflect.DeepEqual(err, nil) {
		t.Errorf("error, man 1")
	}

	shape, err = box.GetByIndex(0)
	if !reflect.DeepEqual(shape, circle) || !reflect.DeepEqual(err, nil) {
		t.Errorf("error, man 2")
	}

	shape, err = box.GetByIndex(1)
	if !reflect.DeepEqual(shape, triangle) || !reflect.DeepEqual(err, nil) {
		t.Errorf("error, man 3")
	}

	shape, err = box.GetByIndex(2)
	if shape != nil || err == nil || !strings.Contains(err.Error(), notFoundError) {
		t.Errorf("error, man 4")
	}
}

func TestReplaceByIndex(t *testing.T) {
	box := NewBox(2)

	circle := Circle{5}
	_ = box.AddShape(circle)

	triangle := Triangle{4}
	_ = box.AddShape(triangle)

	rectangle := Rectangle{2, 3}
	shape, err := box.ReplaceByIndex(1, rectangle)

	if !reflect.DeepEqual(shape, triangle) || !reflect.DeepEqual(err, nil) {
		t.Errorf("error, man 1")
	}

	shape, err = box.ReplaceByIndex(3, triangle)
	if shape != nil || err == nil || !strings.Contains(err.Error(), notFoundError) {
		t.Errorf("error, man 2")
	}
}

func TestSumPerimeter(t *testing.T) {
	box := NewBox(3)

	circle := Circle{1}
	_ = box.AddShape(circle)

	rectangle := Rectangle{2, 3}
	_ = box.AddShape(rectangle)

	triangle := Triangle{3}
	_ = box.AddShape(triangle)

	expected := 2*math.Pi + 10 + 9
	actual := box.SumPerimeter()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("error, man! expected: %b, actual: %b", expected, actual)
	}
}

func TestSumArea(t *testing.T) {
	box := NewBox(3)

	circle := Circle{1}
	_ = box.AddShape(circle)

	rectangle := Rectangle{2, 3}
	_ = box.AddShape(rectangle)

	triangle := Triangle{2}
	_ = box.AddShape(triangle)

	expected := math.Pi + 6 + math.Sqrt(3)
	actual := box.SumArea()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("error, man! expected: %b, actual: %b", expected, actual)
	}
}

func TestRemoveAllCircles(t *testing.T) {
	box := NewBox(4)

	err := box.RemoveAllCircles()
	if reflect.DeepEqual(err, nil) || !strings.Contains(err.Error(), notFoundError) {
		t.Errorf("error, man 1")
	}

	circle := &Circle{1}
	_ = box.AddShape(circle)

	circle = &Circle{2}
	_ = box.AddShape(circle)

	rectangle := &Rectangle{2, 3}
	_ = box.AddShape(rectangle)

	circle = &Circle{3}
	_ = box.AddShape(circle)

	err = box.RemoveAllCircles()

	if !reflect.DeepEqual(len(box.shapes), 1) {
		t.Errorf("error, man 2.1")
	}

	if !reflect.DeepEqual(err, nil) {
		t.Errorf("error, man 2.2")
	}

	shape, err := box.GetByIndex(0)
	if !reflect.DeepEqual(err, nil) || !reflect.DeepEqual(shape, rectangle) {
		t.Errorf("error, man 3")
	}

	shape, err = box.GetByIndex(1)
	if reflect.DeepEqual(err, nil) || !strings.Contains(err.Error(), notFoundError) || !reflect.DeepEqual(shape, nil) {
		t.Errorf("error, man 4")
	}
}
