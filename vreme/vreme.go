// Package vreme omogoča beleženje temperaturnih meritev vremenskih postaj
// ter izpis meritev in osnovnih statističnih poročil.
package vreme

import (
	"fmt"
	"io"
	"os"
	"sort"
)

// Postaja predstavlja vremensko postajo. Njena polja so zasebna, zato jih
// lahko uporabniki paketa spreminjajo samo prek funkcij tega paketa.
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// DodajPostajo doda v slovar postaj novo postajo z dano oznako, imenom in krajem.
// Če oznaka že obstaja, obstoječe postaje ne prepiše.
func DodajPostajo(postaje map[string]Postaja, oznaka, ime, kraj string) {
	if _, obstaja := postaje[oznaka]; obstaja {
		fmt.Fprintf(os.Stdout, "Postaja %s že obstaja.\n", oznaka)
		return
	}
	postaje[oznaka] = Postaja{ime: ime, kraj: kraj, meritve: []float64{}}
}

// DodajMeritev doda temperaturo postaji z dano oznako. Temperatura mora biti
// znotraj območja [minTemp, maxTemp]. Če postaja ne obstaja ali temperatura
// ni dovoljena, funkcija izpiše obvestilo in meritve ne doda.
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura, minTemp, maxTemp float64) {
	postaja, obstaja := postaje[oznaka]
	if !obstaja {
		fmt.Fprintf(os.Stdout, "Postaja %s ne obstaja.\n", oznaka)
		return
	}
	if temperatura < minTemp || temperatura > maxTemp {
		fmt.Fprintf(os.Stdout, "Neveljavna temperatura %.1f °C (dovoljeno: %.1f do %.1f °C).\n", temperatura, minTemp, maxTemp)
		return
	}
	postaja.meritve = append(postaja.meritve, temperatura)
	postaje[oznaka] = postaja
}

// IzpisMeritev izpiše oznake, imena, kraje in meritve vseh postaj.
func IzpisMeritev(postaje map[string]Postaja) {
	izpisMeritev(os.Stdout, postaje)
}

// IzpisPorocila izpiše statistično poročilo za vse postaje. Za izračun
// statistike mora imeti postaja vsaj stMeritev meritev.
func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	izpisPorocila(os.Stdout, postaje, stMeritev)
}

func izpisMeritev(out io.Writer, postaje map[string]Postaja) {
	fmt.Fprintln(out, "MERITVE:")
	for _, oznaka := range oznake(postaje) {
		p := postaje[oznaka]
		fmt.Fprintf(out, "%s - %s (%s): %v\n", oznaka, p.ime, p.kraj, p.meritve)
	}
}

func izpisPorocila(out io.Writer, postaje map[string]Postaja, stMeritev int) {
	for _, oznaka := range oznake(postaje) {
		p := postaje[oznaka]
		min, max, povp, ok := statistika(postaje, oznaka, stMeritev)
		if !ok {
			fmt.Fprintf(out, "%s: premalo meritev\n", p.ime)
			continue
		}
		fmt.Fprintf(out, "%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> %s\n",
			p.ime, min, max, povp, opisTemperature(povp))
	}
}

// statistika vrne najnižjo, najvišjo in povprečno temperaturo postaje.
// Če postaja ne obstaja ali nima dovolj meritev, vrne ok == false.
func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	p, obstaja := postaje[oznaka]
	if !obstaja || stMeritev < 1 || len(p.meritve) < stMeritev {
		return 0, 0, 0, false
	}

	najnizja, najvisja = p.meritve[0], p.meritve[0]
	var vsota float64
	for _, temperatura := range p.meritve {
		if temperatura < najnizja {
			najnizja = temperatura
		}
		if temperatura > najvisja {
			najvisja = temperatura
		}
		vsota += temperatura
	}
	return najnizja, najvisja, vsota / float64(len(p.meritve)), true
}

func opisTemperature(povprecje float64) string {
	switch {
	case povprecje >= 25:
		return "Vroče"
	case povprecje >= 15:
		return "Toplo"
	case povprecje >= 5:
		return "Hladno"
	default:
		return "Mrzlo"
	}
}

func oznake(postaje map[string]Postaja) []string {
	rezultat := make([]string, 0, len(postaje))
	for oznaka := range postaje {
		rezultat = append(rezultat, oznaka)
	}
	sort.Strings(rezultat)
	return rezultat
}
