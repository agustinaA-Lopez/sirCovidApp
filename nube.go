package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type nube struct {
	X       [20]float64
	Y       [20]float64
	Alpha   [20]float64
	AlphaUp [20]bool
	img     *ebiten.Image
}

var (
	nube1         nube
	nubScale      = float64(.4)
	nubFrameHight = float64(300) * nubScale
	nubFrameWidth = float64(300) * nubScale
	numNube       = 2
)

func posicionNube() (float64, float64) {
	rand.Seed(time.Now().UnixNano())
	var x, y float64
	for i := 0; i < numNube; i++ {
		nube1.X[i], nube1.Y[i] = float64(rand.Intn(screenWidth/.4)), float64(rand.Intn(screenHeight/.4)) // enemigo.Y[i] //
		nube1.Alpha[i] = float64(rand.Intn(11)) * .1
		x, y = nube1.X[i], nube1.Y[i]
	}

	return x, y
}
func initNube() {
	switch {
	case banco || casita:
		numNube = 0
	default:
		numNube = Level * 2
		if numNube > 15 {
			numNube = 15
		}
	}

	posicionNube()

	nube1.img, _, err = ebitenutil.NewImageFromFile(`data/smoke.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

}

//// nubeCovid aumenta y disminuye transparencia de la nube (alpha)
func moverNube(n nube) nube {
	rand.Seed(time.Now().UnixNano())

	// creacion de nuevas nubes
	for i := 0; i < numNube; i++ {
		if n.Alpha[i] >= 0 {
			n.X[i]--
		}
		if n.Alpha[i] <= 0 {
			n.X[i], n.Y[i] = posicionNube()
			n.AlphaUp[i] = true
		} else if n.Alpha[i] > .9 {
			n.AlphaUp[i] = false
		}

		// movimiento nube
		if n.AlphaUp[i] {
			n.Alpha[i] += .009
		} else {
			n.Alpha[i] -= .003
		}
	}
	return n
}

func dibujarNube(n nube, screen *ebiten.Image) {
	for i := 0; i < numNube; i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(n.X[i], n.Y[i])
		op.ColorM.Scale(3, 2, 0, n.Alpha[i])
		op.GeoM.Scale(nubScale, nubScale)
		screen.DrawImage(n.img, op)
	}
}
