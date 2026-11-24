// Program vreme je primer uporabe paketa vreme za beleženje temperatur.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/kimisteva/vreme/vreme"
	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "vreme",
		Usage: "beleženje meritev vremenskih postaj in izpis statistike",
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:  "minTemp",
				Value: -50,
				Usage: "najnižja dovoljena temperatura meritve",
			},
			&cli.Float64Flag{
				Name:  "maxTemp",
				Value: 50,
				Usage: "najvišja dovoljena temperatura meritve",
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Value: 3,
				Usage: "najmanjše število meritev za izračun statistike",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			minTemp := cmd.Float64("minTemp")
			maxTemp := cmd.Float64("maxTemp")
			stMeritev := cmd.Int("stMeritev")
			if minTemp > maxTemp {
				return cli.Exit("minTemp ne sme biti večji od maxTemp", 2)
			}
			if stMeritev < 1 {
				return cli.Exit("stMeritev mora biti vsaj 1", 2)
			}

			postaje := make(map[string]vreme.Postaja)
			vreme.DodajPostajo(postaje, "KP01", "Koper Luka", "Koper")
			vreme.DodajPostajo(postaje, "KR01", "Kredarica", "Triglav")
			vreme.DodajPostajo(postaje, "LJ01", "Ljubljana Bežigrad", "Ljubljana")
			vreme.DodajPostajo(postaje, "MB01", "Maribor Tabor", "Maribor")

			// Začetne meritve iz demonstracijskega primera domače naloge 1.
			for _, m := range []struct {
				oznaka string
				temp   float64
			}{
				{"KP01", 24.5}, {"KP01", 27}, {"KP01", 26.1}, {"KP01", 28.4},
				{"KR01", 2}, {"KR01", -1.5}, {"KR01", 3.5},
				{"LJ01", 18.2}, {"LJ01", 21.5}, {"LJ01", 19.8}, {"LJ01", 16.5},
				{"MB01", 12.4},
			} {
				vreme.DodajMeritev(postaje, m.oznaka, m.temp, minTemp, maxTemp)
			}

			vreme.IzpisMeritev(postaje)
			fmt.Println()
			vreme.IzpisPorocila(postaje, stMeritev)
			return nil
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
