package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestParseTab(t *testing.T) {
	tests := []struct {
		text   string
		lines  [][]string
		widths []int
	}{
		{
			text: `HCN                & $B$ &   44580.6 & 44386.4  & 44316 & 70.4 & 70.4  & 0.0 \\
HCO$^+$            & $B$ &   44851.4 & 44611.1  & 44677.3 & -66.2 & 66.2  & 0.0 \\
HNC                & $B$ &   45571.5 & 45405.3  & 45332 & 73.3 & 73.3  & 0.0 \\
C$_2$H             & $B$ &   44187.6 & 43702.2  & 43674.5 & 27.7 & 27.7  & 0.0 \\`,
			lines: [][]string{
				{"HCN", "$B$", "44580.6", "44386.4", "44316", "70.4", "70.4", "0.0 \\\\"},
				{"HCO$^+$", "$B$", "44851.4", "44611.1", "44677.3", "-66.2", "66.2", "0.0 \\\\"},
				{"HNC", "$B$", "45571.5", "45405.3", "45332", "73.3", "73.3", "0.0 \\\\"},
				{"C$_2$H", "$B$", "44187.6", "43702.2", "43674.5", "27.7", "27.7", "0.0 \\\\"},
			},
			widths: []int{7, 3, 7, 7, 7, 5, 4, 6},
		},

		{
			text: `A$_0$, B$_0$, \& C$_0$ & MHz & 536.5 \\
B$_0$ \& C$_0$ & MHz & 93.2 \\`,
			lines: [][]string{
				{"A$_0$, B$_0$, \\& C$_0$", "MHz", "536.5 \\\\"},
				{"B$_0$ \\& C$_0$", "MHz", "93.2 \\\\"},
			},
			widths: []int{22, 3, 8},
		},
		{
			text: `&  & \multicolumn{2}{c}{CcCR} & \\
Molecule &  & Equil. & Vib.~Avg. & Experiment & Difference & $\vert$Diff.$\vert$ & \% Error$^a$\\`,
			lines: [][]string{
				{"", "", "\\multicolumn{2}{c}{CcCR}", "\\\\"},
				{"Molecule", "", "Equil.", "Vib.~Avg.", "Experiment", "Difference", "$\\vert$Diff.$\\vert$", "\\% Error$^a$\\\\"},
			},
			widths: []int{8, 0, 6, 9, 10, 10, 19, 14},
		},
	}
	for _, test := range tests {
		lines, widths := ParseTab(test.text)
		if !reflect.DeepEqual(test.lines, lines) {
			t.Errorf("got\n%q,\nwanted\n%q\n", lines, test.lines)
		} else if !reflect.DeepEqual(test.widths, widths) {
			t.Errorf("got\n%v,\nwanted\n%v\n", widths, test.widths)
		}
	}
}

func TestWriteTab(t *testing.T) {
	var buf bytes.Buffer
	tests := []struct {
		lines  [][]string
		widths []int
		res    string
	}{
		{
			lines: [][]string{
				{"     5", "   3", " 1 \\\\"},
				{"   3", "    4", "  2 \\\\"},
				{"     5", "0", "   3 \\\\"},
			},
			widths: []int{6, 5, 7},
			res: "     5 &     3 &    1 \\\\\n" +
				"   3   &     4 &    2 \\\\\n" +
				"     5 &     0 &    3 \\\\\n",
		},
	}
	for _, test := range tests {
		buf.Reset()
		WriteTab(&buf, test.lines, test.widths)
		if test.res != buf.String() {
			t.Errorf("got\n%q,\nwanted\n%q\n", buf.String(), test.res)
		}
	}
}
