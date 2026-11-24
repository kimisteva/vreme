package vreme

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

func DodajPostajo(postaje map[string]Postaja, oznaka, ime, kraj string) {
	if _, obstaja := postaje[oznaka]; obstaja {
		fmt.Fprintf(os.Stdout, "Postaja %s že obstaja.\n", oznaka)
		return
	}

	postaje[oznaka] = Postaja{
		ime:     ime,
		kraj:    kraj,
		meritve: []float64{},
	}
}

func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura, minTemp, maxTemp float64) {
	postaja, obstaja := postaje[oznaka]

	if !obstaja {
		fmt.Fprintf(os.Stdout, "Postaja %s ne obstaja.\n", oznaka)
		return
	}

	if temperatura < minTemp || temperatura > maxTemp {
		fmt.Fprintf(os.Stdout, "Neveljavna temperatura %.1f °C (dovoljeno: %.1f do %.1f °C).\n",
			temperatura, minTemp, maxTemp)
		return
	}

	postaja.meritve = append(postaja.meritve, temperatura)
	postaje[oznaka] = postaja
}

func IzpisMeritev(postaje map[string]Postaja) {
	izpisMeritev(os.Stdout, postaje)
}

func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	izpisPorocila(os.Stdout, postaje, stMeritev)
}

func izpisMeritev(out io.Writer, postaje map[string]Postaja) {
	fmt.Fprintln(out, "MERITVE:")

	for _, oznaka := range oznake(postaje) {
		postaja := postaje[oznaka]
		fmt.Fprintf(out, "%s - %s (%s): %v\n",
			oznaka, postaja.ime, postaja.kraj, postaja.meritve)
	}
}

func izpisPorocila(out io.Writer, postaje map[string]Postaja, stMeritev int) {
	for _, oznaka := range oznake(postaje) {
		postaja := postaje[oznaka]

		min, max, povp, ok := statistika(postajaMap(postaje), oznaka, stMeritev)

		if !ok {
			fmt.Fprintf(out, "%s: premalo meritev\n", postaja.ime)
			continue
		}

		var opis string

		switch {
		case povp >= 25:
			opis = "Vroče"
		case povp >= 15:
			opis = "Toplo"
		case povp >= 5:
			opis = "Hladno"
		default:
			opis = "Mrzlo"
		}

		fmt.Fprintf(out, "%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> %s\n",
			postaja.ime, min, max, povp, opis)
	}
}

func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	postaja, obstaja := postaje[oznaka]
	stevMeritev := len(postaja.meritve)

	if stevMeritev < stMeritev || !obstaja || stMeritev < 1 {
		return 0, 0, 0, false
	}

	var avg float64
	najvisja = -50
	najnizja = 50

	for i := 0; i < len(postaja.meritve); i++ {
		avg += postaja.meritve[i]

		if postaja.meritve[i] < najnizja {
			najnizja = postaja.meritve[i]
		}

		if postaja.meritve[i] > najvisja {
			najvisja = postaja.meritve[i]
		}
	}

	return najnizja, najvisja, avg / float64(stevMeritev), true
}

func oznake(postaje map[string]Postaja) []string {
	rezultat := make([]string, 0, len(postaje))

	for oznaka := range postaje {
		rezultat = append(rezultat, oznaka)
	}

	sort.Strings(rezultat)
	return rezultat
}

func postajaMap(postaje map[string]Postaja) map[string]Postaja {
	return postaje
}
