package main

import (
	"bytes"
	"fmt"
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
		{
			text: `\begin{tabular}{lllrl}
  \hline
  Mode & Symm. &                     Desc. & Int.&   Freq. & \\
  \hline
  $\omega_{1}$ & $a_1$ &         symm. C-H stretch &    0&  3380.2 & \\
  $\omega_{2}$ & $t_2$ &    anti-symm. C-H stretch &   59&  3344.1 & \\
  $\omega_{3}$ & $a_1$ &                 breathing &    0&  1439.2 & \\
  $\omega_{4}$ & $t_2$ &    anti-symm. C-C stretch &   14&  1149.0 & \\
  $\omega_{5}$ & $t_1$ &                     H wag &    0&   888.7 & \\
  $\omega_{6}$ &   $e$ & C trapezoidal deformation &    0&   837.2 & \\
  $\omega_{7}$ & $t_2$ &                     H wag &  183&   774.9 & \\
  $\omega_{8}$ &   $e$ & H trapezoidal deformation &    0&   562.6 & \\
  \hline
  ZPT        &       &                          &      & 12836.7 &\\
  \hline
  $\nu_{1}$ & $a_1$ &         symm. C-H stretch &     &  3242.6 & \\
  $\nu_{2}$ & $t_2$ &    anti-symm. C-H stretch &     &  3210.6 & \\
  $\nu_{3}$ & $a_1$ &                 breathing &     &  1425.1 & \\
  $\nu_{4}$ & $t_2$ &    anti-symm. C-C stretch &     &  1095.6 & \\
  $\nu_{5}$ & $t_1$ &                     H wag &     &   849.7 & \\
  $\nu_{6}$ &   $e$ & C trapezoidal deformation &     &   808.1 & \\
  $\nu_{7}$ & $t_2$ &                     H wag &     &   752.5 & \\
  $\nu_{8}$ &   $e$ & H trapezoidal deformation &     &   489.8 & \\
  \hline
\end{tabular}
`,
			lines: [][]string{
				{`\begin{tabular}{lllrl}`},
				{`\hline`},
				{"Mode", "Symm.", "Desc.", "Int.", "Freq.", `\\`},
				{`\hline`},
				{`$\omega_{1}$`, "$a_1$", "symm. C-H stretch", "0", "3380.2", `\\`},
				{`$\omega_{2}$`, "$t_2$", "anti-symm. C-H stretch", "59", "3344.1", `\\`},
				{`$\omega_{3}$`, "$a_1$", "breathing", "0", "1439.2", `\\`},
				{`$\omega_{4}$`, "$t_2$", "anti-symm. C-C stretch", "14", "1149.0", `\\`},
				{`$\omega_{5}$`, "$t_1$", "H wag", "0", "888.7", `\\`},
				{`$\omega_{6}$`, "$e$", "C trapezoidal deformation", "0", "837.2", `\\`},
				{`$\omega_{7}$`, "$t_2$", "H wag", "183", "774.9", `\\`},
				{`$\omega_{8}$`, "$e$", "H trapezoidal deformation", "0", "562.6", `\\`},
				{`\hline`},
				{"ZPT", "", "", "", "12836.7", `\\`},
				{`\hline`},
				{`$\nu_{1}$`, "$a_1$", "symm. C-H stretch", "", "3242.6", `\\`},
				{`$\nu_{2}$`, "$t_2$", "anti-symm. C-H stretch", "", "3210.6", `\\`},
				{`$\nu_{3}$`, "$a_1$", "breathing", "", "1425.1", `\\`},
				{`$\nu_{4}$`, "$t_2$", "anti-symm. C-C stretch", "", "1095.6", `\\`},
				{`$\nu_{5}$`, "$t_1$", "H wag", "", "849.7", `\\`},
				{`$\nu_{6}$`, "$e$", "C trapezoidal deformation", "", "808.1", `\\`},
				{`$\nu_{7}$`, "$t_2$", "H wag", "", "752.5", `\\`},
				{`$\nu_{8}$`, "$e$", "H trapezoidal deformation", "", "489.8", `\\`},
				{`\hline`},
				{`\end{tabular}`},
				{""},
			},
			widths: []int{
				12, 5, 25, 4, 7, 2,
			},
		},
	}
	for _, test := range tests {
		lines, widths := ParseTab(test.text)
		if !reflect.DeepEqual(test.lines, lines) {
			for i, l := range test.lines {
				if !reflect.DeepEqual(l, lines[i]) {
					fmt.Printf("%q != %q\n\n", l, lines[i])
				}
			}
			t.Errorf("got\n%q,\nwanted\n%q\n", lines, test.lines)
		} else if !reflect.DeepEqual(test.widths, widths) {
			t.Errorf("got\n%v,\nwanted\n%v\n", widths, test.widths)
		}
	}
}

func TestWriteTab(t *testing.T) {
	var buf bytes.Buffer
	tests := []struct {
		res    string
		lines  [][]string
		widths []int
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
