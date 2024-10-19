package main

import "testing"

func TestDimension(t *testing.T) {
    dim := []uint8{
        2, 3, 4,
    };

    expected := 58;
    result := GetSurfaceAreaOfBox([3]uint8(dim));

    if result != uint8(expected) {
        t.Errorf("expected %d but got %d", expected, result);
    }
}

func TestDimensionTwo(t *testing.T) {
    dim := []uint8{
        1, 1, 10,
    };

    expected := 43;
    result := GetSurfaceAreaOfBox([3]uint8(dim));

    if result != uint8(expected) {
        t.Errorf("expected %d but got %d", expected, result);
    }
}
