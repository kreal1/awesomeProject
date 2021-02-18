package main

import "testing"

func TestSelect1(t *testing.T) {
	result := Select1(1)
	if result == "" {
		t.Errorf("Данные не получены")
	}
}

func TestSelectAll(t *testing.T) {
	result := SelectAll()
	if result == "" {
		t.Errorf("Данные не получены")
	}
}
