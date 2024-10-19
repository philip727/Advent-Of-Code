package main

import "testing"

func TestFloorZero(t *testing.T) {
    floor := 0;
    directions := "()()()";
    GetNewFloorFromDirections(&floor, directions);

    expected := 0;
    if floor != expected {
        t.Errorf("expected %d but got %d", expected, floor);
    }
}

func TestFloorThree(t *testing.T) {
    floor := 0;
    directions := "((()(";
    GetNewFloorFromDirections(&floor, directions);

    expected := 3;
    if floor != expected {
        t.Errorf("expected %d but got %d", expected, floor);
    }
}

func TestFloorMinusThree(t *testing.T) {
    floor := 0;
    directions := ")))";
    GetNewFloorFromDirections(&floor, directions);

    expected := -3;
    if floor != expected {
        t.Errorf("expected %d but got %d", expected, floor);
    }
}
