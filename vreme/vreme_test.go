package vreme

import "testing"

func TestDodajPostajoInMeritev(t *testing.T) {
	postaje := make(map[string]Postaja)
	DodajPostajo(postaje, "LJ01", "Ljubljana", "Ljubljana")
	DodajMeritev(postaje, "LJ01", 10, -50, 50)
	if got := len(postaje["LJ01"].meritve); got != 1 {
		t.Fatalf("pričakovana 1 meritev, dobil %d", got)
	}
}

func TestStatistika(t *testing.T) {
	postaje := map[string]Postaja{"LJ01": {ime: "Ljubljana", kraj: "Ljubljana", meritve: []float64{10, 20, 30}}}
	min, max, povp, ok := statistika(postaje, "LJ01", 3)
	if !ok || min != 10 || max != 30 || povp != 20 {
		t.Fatalf("napačna statistika: %v, %v, %v, %v", min, max, povp, ok)
	}
	if _, _, _, ok := statistika(postaje, "LJ01", 4); ok {
		t.Fatal("statistika bi morala biti nedostopna pri premalo meritvah")
	}
}
