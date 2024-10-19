package main

import (
	"testing"
)

func TestDirections(t *testing.T) {
    directions := ">";
    expected := 2;
    result := VisitHouses(directions);

    if expected != int(result) {
        t.Errorf("expected %d but got %d", expected, result);
    }
}

func TestDirectionsTwoHouses(t *testing.T) {
    directions := "^v^v^v^v^v";
    expected := 2;
    result := VisitHouses(directions);

    if expected != int(result) {
        t.Errorf("expected %d but got %d", expected, result);
    }
}
