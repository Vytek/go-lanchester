package main

import (
	"fmt"
	"go-lanchester/lanchester"
	"log"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Parametri iniziali
	R0 := 8000.0 // Numero di unità ROSSE
	B0 := 10000.0 // Numero di unità BLU
	rS := 0.05   // Efficacia al combattimento delle unità ROSSE
	bS := 0.04   // Efficacia al combattimento delle unità BLU
	T := 100     // Durata totale della simulazione
	dt := 1      // Intervallo di tempo

	// Simulazione utilizzando il modello quadratico
	R, B := lanchester.SquareLaw(R0, B0, rS, bS, T, dt)

	// Determinazione del vincitore
	var winner string
	if R[len(R)-1] > B[len(B)-1] {
		winner = "ROSSO"
	} else {
		winner = "BLU"
	}

	fmt.Printf("Risultato della battaglia:\n")
	fmt.Printf("Vincitore: %s\n", winner)
	fmt.Printf("Unità ROSSE rimanenti: %.2f\n", R[len(R)-1])
	fmt.Printf("Unità BLU rimanenti: %.2f\n", B[len(B)-1])

	// Creazione del grafico
	err := plotBattle(R, B, dt)
	if err != nil {
		log.Fatalf("Errore nella creazione del grafico: %v", err)
	}
}

func plotBattle(R, B []float64, dt int) error {
	p := plot.New()
	p.Title.Text = "Simulazione del Modello di Lanchester"
	p.X.Label.Text = "Tempo (round)"
	p.Y.Label.Text = "Numero di unità"

	// Creazione dei punti per le unità ROSSE
	redPoints := make(plotter.XYs, len(R))
	for i := range R {
		redPoints[i].X = float64(i * dt)
		redPoints[i].Y = R[i]
	}

	// Creazione dei punti per le unità BLU
	bluePoints := make(plotter.XYs, len(B))
	for i := range B {
		bluePoints[i].X = float64(i * dt)
		bluePoints[i].Y = B[i]
	}

	// Creazione delle linee per il grafico
	redLine, err := plotter.NewLine(redPoints)
	if err != nil {
		return err
	}
	redLine.LineStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Colore rosso

	blueLine, err := plotter.NewLine(bluePoints)
	if err != nil {
		return err
	}
	blueLine.LineStyle.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255} // Colore blu

	// Aggiunta delle linee al grafico
	p.Add(redLine, blueLine)
	p.Legend.Add("Unità ROSSE", redLine)
	p.Legend.Add("Unità BLU", blueLine)

	// Salvataggio del grafico come file PNG
	if err := p.Save(10*vg.Inch, 6*vg.Inch, "battle_simulation.png"); err != nil {
		return err
	}

	fmt.Println("Grafico salvato come 'battle_simulation.png'")
	return nil
}
