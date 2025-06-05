package main

import (
	"testing"
)

func TestAdd(t *testing.T){
	total := add(5,6)
	want := 11
	if total != want{
		t.Fatalf("error - want %d and got %d ",want,total)
	}
}

func TestMinus(t *testing.T){
	got := subtract(5,3)
	want := 2
	if got != want{
		t.Fatalf("error - want %d and got %d ",want,got)
	}
}