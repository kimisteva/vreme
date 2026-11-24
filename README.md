# Vreme

Go modul za beleženje temperaturnih meritev vremenskih postaj. Modul vsebuje knjižnični paket `vreme` in demonstracijski program `cmd/vreme`.

## Zagon

```sh
go mod tidy
go run ./cmd/vreme
```

## Stikala

- `--minTemp`: najnižja dovoljena temperatura (privzeto `-50`).
- `--maxTemp`: najvišja dovoljena temperatura (privzeto `50`).
- `--stMeritev`: najmanjše število meritev za izračun statistike (privzeto `3`).

Primer:

```sh
go run ./cmd/vreme --minTemp=-30 --maxTemp=40 --stMeritev=2
```

## Uporaba paketa

```go
postaje := make(map[string]vreme.Postaja)
vreme.DodajPostajo(postaje, "LJ01", "Ljubljana Bežigrad", "Ljubljana")
vreme.DodajMeritev(postaje, "LJ01", 21.5, -50, 50)
vreme.IzpisMeritev(postaje)
vreme.IzpisPorocila(postaje, 1)
```

## Licenca

Projekt je objavljen pod licenco MIT. Glej datoteko [LICENSE](LICENSE).
