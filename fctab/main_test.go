package main

import (
	"reflect"
	"testing"
)

func TestRead15(t *testing.T) {
	got := Read15("testfiles/small.15")
	want := []string{
		"$F_{1,1}$ & 0.7104237500",
		"$F_{1,2}$ & -0.0000001512",
		"$F_{1,3}$ & -0.0000000252",
		"$F_{1,4}$ & -0.0625424581",
		"$F_{1,5}$ & -0.0000000980",
		"$F_{1,6}$ & 0.0000001456",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v\n", got, want)
	}
}

func TestDimMap3(t *testing.T) {
	got := len(DimMap3(9))
	want := 165
	if got != want {
		t.Errorf("got %v, wanted %v\n", got, want)
	}
}

func TestRead30(t *testing.T) {
	got := Read30("testfiles/small.30")
	want := []string{
		"$F_{1,1,1}$ & 0.0000109630",
		"$F_{1,1,2}$ & -1.3363443890",
		"$F_{1,2,2}$ & 0.0000008844",
		"$F_{2,2,2}$ & 1.3362504436",
		"$F_{1,1,3}$ & -0.8735378010",
		"$F_{1,2,3}$ & 0.0000028134",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v\n", got, want)
	}
}

func TestDimMap4(t *testing.T) {
	got := len(DimMap4(9))
	want := 495
	if got != want {
		t.Errorf("got %v, wanted %v\n", got, want)
	}
}

func TestRead40(t *testing.T) {
	got := Read40("testfiles/small.40")
	want := []string{
		"$F_{1,1,1,1}$ & 1.2548200557",
		"$F_{1,1,1,2}$ & 0.0013250835",
		"$F_{1,1,2,2}$ & 0.4185298029",
		"$F_{1,2,2,2}$ & 0.0018503920",
		"$F_{2,2,2,2}$ & 1.2569845853",
		"$F_{1,1,1,3}$ & 0.0003606063",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v\n", got, want)
	}
}
