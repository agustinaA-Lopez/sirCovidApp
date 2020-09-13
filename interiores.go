package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

//bank ajuste interiores
func interior() {
	hScaleW = .65
	hScaleH = .55
	hgt = float64(player1.FrameHeight[0]) * hScaleH
	wth = float64(player1.FrameWidth[0]) * hScaleW
	switch {
	case casita:
		imgCasita, _, err = ebitenutil.NewImageFromFile(`data/casita.png`, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}

	case banco:
		imgBanco, _, err = ebitenutil.NewImageFromFile(`data/banco.png`, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
		imgCintas, _, err = ebitenutil.NewImageFromFile(`data/cintas.png`, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
	}
	initEnemigos()
	initNube()
}

func salida() {
	if banco || casita {
		banco = false
		casita = false
		hScaleW = .53
		hScaleH = .45
		hgt = float64(player1.FrameHeight[0]) * hScaleH
		wth = float64(player1.FrameWidth[0]) * hScaleW
		initNube()
		initEnemigos()
	}
}
