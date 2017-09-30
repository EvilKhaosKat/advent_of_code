package main

import (
	"testing"
)

func TestLineIntersection(t *testing.T) {
	horizontalLine := Line{Coor{-10, 0}, Coor{10, 0}}
	verticalLine := Line{Coor{0, -10}, Coor{0, 10}}

	point := verticalLine.isIntersected(horizontalLine)

	if point == nil {
		t.Error("point not found")
	}

	if point.x != 0 || point.y != 0 {
		t.Error("Coordinate is wrong, must be 0:0, but:", point)
	}
}

func TestLineIntersectionFail(t *testing.T) {
	horizontalLine := Line{Coor{-10, 0}, Coor{10, 0}}
	verticalLine := Line{Coor{20, -10}, Coor{20, 10}}

	point := verticalLine.isIntersected(horizontalLine)

	if point != nil {
		t.Error("point found, but should not:", point)
	}
}
