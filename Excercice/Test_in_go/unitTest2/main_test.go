package main

import (
	"log"
	"testing"
)

func TestText(T *testing.T) {
	got := paradise("palm")
	want := "Hello palm"
	if got != want{
		log.Fatalf("Error ")
	}
}