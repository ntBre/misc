package main

import (
	"testing"
)

func TestPDFinfo(t *testing.T) {
	tests := []struct {
		pdf  string
		want ref
	}{
		{
			pdf: "testfiles/Westbrook20.pdf",
			want: ref{
				name:    "Westbrook20",
				author:  "Brent R. Westbrook and Ryan C. Fortenberry",
				title:   "jp0c01609 1..14",
				journal: "J. Phys. Chem. A 2020.124:3191-3204",
			},
		},
	}
	for _, test := range tests {
		got := pdfinfo(test.pdf)
		want := test.want.String()
		if got != want {
			t.Errorf("got\n%q, wanted\n%q\n", got, want)
		}
	}
}
