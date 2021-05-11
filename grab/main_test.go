package main

import (
	"path/filepath"
	"testing"
)

func TestProcess(t *testing.T) {
	res, _ := Process(filepath.Join(base, "job.0000007241.out"))
	pat := `H 0.0000000000 0.7574590974 0.5217905143
O 0.0000000000 0.0000000000 -0.0657441568
H 0.0000000000 -0.7574590974 0.5017905143
`
	if res.Geom != pat {
		t.Errorf("got %v, wanted %v\n", res.Geom, pat)
	}
	want := -76.369682932866
	if res.Val != want {
		t.Errorf("got %v, wanted %v\n", res.Val, want)
	}
}
