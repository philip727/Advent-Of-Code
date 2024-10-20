package main

import "testing"

func TestAbcdef(t *testing.T) {
    key := "abcdef";
    expected := 609043;
    result := FindLowestIterOfKey(key);

    if result != expected {
        t.Errorf("Expected result %d but got %d", expected, result);
    }
}

func TestPqrstuv(t *testing.T) {
    key := "pqrstuv";
    expected := 1048970;
    result := FindLowestIterOfKey(key);

    if result != expected {
        t.Errorf("Expected result %d but got %d", expected, result);
    }
}
