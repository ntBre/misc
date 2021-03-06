# misc
Miscellaneous command line tools written in Go

## orgtab
orgtab converts Emacs org-mode tables to LaTeX tables

## pdfinfgo
pdfinfgo converts the output of `pdfinfo` to a format useable by BibTex. 
I use it along with `entr` to automatically generate citations for PDF
files I add to my references directory

## tabfmt
tabfmt formats columns LaTeX tables based on the width of their widest cells

## fctab
fctab formats force constants from Spectro input files into LaTeX
tables for reporting in SI sections of papers

## regexpl
regexpl is a simple command line regex tester for Go

## tmpcln
tmpcln is a utility for deleting tmp files on the nodes on Maple

## matrview
matrview is a tool for viewing what parts of matrices meet
conditions. I used it for deriving expressions for the number of
points a Cartesian pbqff calculation would entail, so it's pretty
cluttered from that.  The actual matrix visualization part may be more
broadly applicable though.

## stripcite
stripcite is a very basic tool for grabbing the \cite{} commands out
of a LaTeX document