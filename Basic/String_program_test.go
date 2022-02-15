package main

import "testing"

func TestCharacter(t *testing.T) {
	if Concat("Vivek", 2) != "Vi" {
		t.Errorf("Expected : 2 character but got different")
	}
}
