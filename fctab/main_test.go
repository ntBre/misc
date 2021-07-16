package main

import (
	"reflect"
	"testing"
)

func TestRead15(t *testing.T) {
	tests := []struct {
		file string
		want []string
	}{
		{
			file: "testfiles/small.15",
			want: []string{
				"$F_{1,1}$ & 0.710424",
				"$F_{1,2}$ & -0.000000",
				"$F_{1,3}$ & -0.000000",
				"$F_{1,4}$ & -0.062542",
				"$F_{1,5}$ & -0.000000",
				"$F_{1,6}$ & 0.000000",
			},
		},
		{
			file: "testfiles/small5.15",
			want: []string{
				"$F_{1,1}$ & 0.710424",
				"$F_{1,2}$ & -0.000000",
				"$F_{1,3}$ & -0.000000",
				"$F_{1,4}$ & -0.062542",
				"$F_{1,5}$ & -0.000000",
			},
		},
	}
	for _, test := range tests {
		got := Read15(test.file)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got %v, wanted %v\n", got, test.want)
		}
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
	tests := []struct {
		file string
		want []string
	}{
		{
			file: "testfiles/small.30",
			want: []string{
				"$F_{1,1,1}$ & 0.0000",
				"$F_{1,1,2}$ & -1.3363",
				"$F_{1,2,2}$ & 0.0000",
				"$F_{2,2,2}$ & 1.3363",
				"$F_{1,1,3}$ & -0.8735",
				"$F_{1,2,3}$ & 0.0000",
			},
		},
		{
			file: "testfiles/small5.30",
			want: []string{
				"$F_{1,1,1}$ & 0.0000",
				"$F_{1,1,2}$ & -1.3363",
				"$F_{1,2,2}$ & 0.0000",
				"$F_{2,2,2}$ & 1.3363",
				"$F_{1,1,3}$ & -0.8735",
			},
		},
	}
	for _, test := range tests {
		got := Read30(test.file)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got %v, wanted %v\n", got, test.want)
		}
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
	tests := []struct {
		file string
		want []string
	}{
		{
			file: "testfiles/small.40",
			want: []string{
				"$F_{1,1,1,1}$ & 1.25",
				"$F_{1,1,1,2}$ & 0.00",
				"$F_{1,1,2,2}$ & 0.42",
				"$F_{1,2,2,2}$ & 0.00",
				"$F_{2,2,2,2}$ & 1.26",
				"$F_{1,1,1,3}$ & 0.00",
			},
		},
		{
			file: "testfiles/small5.40",
			want: []string{
				"$F_{1,1,1,1}$ & 1.25",
				"$F_{1,1,1,2}$ & 0.00",
				"$F_{1,1,2,2}$ & 0.42",
				"$F_{1,2,2,2}$ & 0.00",
				"$F_{2,2,2,2}$ & 1.26",
			},
		},
	}
	for _, test := range tests {
		got := Read40(test.file)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got %v, wanted %v\n", got, test.want)
		}
	}
}
