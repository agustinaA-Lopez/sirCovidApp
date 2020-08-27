package main

import (
	"image"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func init() {

}

var (
	enemigo humanos
	count   int
	tmp     int
	obs     bool
	match   bool
	x       float64
	y       float64
	en      string
	i       int
	// eScaleW    float64
	// eScaleH    float64
	numEnemigo = 2
	randNum    int
)

func randXY() (x float64, y float64) {
	rand.Seed(time.Now().UnixNano())
	_, _, obs := obstaculos(x, y, x, y)
	for obs {
		x = float64(rand.Intn(screenWidth))
		y = float64(rand.Intn(screenHeight))
		_, _, obs = obstaculos(x, y, x, y)
	}
	return
}

func initEnemigos() {
	numEnemigo = Level * 2
	if numEnemigo > 20 {
		numEnemigo = 20
	}
	if numEnemigo < 20 {
		randNum = rand.Intn(20 - numEnemigo)
	} else {
		randNum = 0
	}
	if banco && numEnemigo > 3 {
		numEnemigo = 3
	}
	if casita {
		numEnemigo = 0
	}

	rand.Seed(time.Now().UnixNano())
	for i = randNum; i < numEnemigo+randNum; i++ {
		en = `sircovid\data\enemigo` + strconv.Itoa(i+1) + `.png`
		enemigo.FrameOX[i] = 48
		enemigo.FrameOY[i] = 72 * rand.Intn(4)
		enemigo.FrameNum[i] = 1
		enemigo.X[i], enemigo.Y[i] = randXY()
		enemigo.FrameWidth[i] = 48
		enemigo.FrameHeight[i] = 72
		enemigo.num[i] = rand.Intn(5)
		enemigo.cambio[i] = rand.Intn(50) + 100
		enemigo.img[i], _, err = ebitenutil.NewImageFromFile(en, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func moverHumanos(E humanos) humanos {
	// var i int
	rand.Seed(time.Now().UnixNano())
	if ModeGame {
		count++
	}
	for i = randNum; i < numEnemigo+randNum; i++ {

		E.FrameNum[i] = 3
		E.FrameOX[i] = 0

		if count >= E.cambio[i] {
			for tmp = E.num[i]; tmp == E.num[i]; tmp = rand.Intn(5) {
			}
			E.cambio[i] += rand.Intn(200) + 200
			E.num[i] = tmp
		}

		E.X1[i], E.Y1[i] = E.X[i], E.Y[i]

		switch E.num[i] {
		case 0:
			E.FrameNum[i] = 1
			E.FrameOX[i] = 48
		case 1:
			E.FrameOY[i] = 72
			E.X[i]--
		case 2:
			E.FrameOY[i] = 144
			E.X[i]++
		case 3:
			E.FrameOY[i] = 216
			E.Y[i]--
		case 4:
			E.FrameOY[i] = 0
			E.Y[i]++
		}

		E.X[i], E.Y[i], obs = obstaculos(E.X[i], E.Y[i], E.X1[i], E.Y1[i])
		for j := randNum; j < numEnemigo+randNum; j++ {
			if i != j && E.X[i]+(wth-7) > E.X[j]+5 && E.X[i]+5 < E.X[j]+(wth-7) && E.Y[i]+hgt > E.Y[j]+(hgt-10) && E.Y[i]+(hgt-10) < E.Y[j]+hgt {
				E.X[i] = E.X1[i]
				E.Y[i] = E.Y1[i]
				E.X[j] = E.X1[j]
				E.Y[j] = E.Y1[j]
				obs = true
				E.num[j] = 0
				E.cambio[j] = count + 10
			}
		}

		if E.X[i]+(wth-7) >= player1.X[0]+5 && E.X[i]+5 <= player1.X[0]+(wth-7) && E.Y[i]+hgt >= player1.Y[0]+(hgt-10) && E.Y[i]+(hgt-10) <= player1.Y[0]+hgt ||
			(E.X[i]+(wth-7) >= player2.X[0]+5 && E.X[i]+5 <= player2.X[0]+(wth-7) && E.Y[i]+hgt >= player2.Y[0]+(hgt-10) && E.Y[i]+(hgt-10) <= player2.Y[0]+hgt) {
			E.X[i], E.Y[i] = E.X1[i], E.Y1[i]
			obs = true
		}

		if obs {
			E.num[i] = 0
			E.cambio[i] = count + 10
		}
	}
	return E
}

func dibujarEnemigos(E humanos, screen *ebiten.Image) {
	for i = randNum; i < numEnemigo+randNum; i++ {
		if ModePause || ModeWin || ModeTitleLevel || ModeGameOver {
			E.FrameNum[i] = 1
			E.FrameOX[i] = 48
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(hScaleW, hScaleH)
		op.GeoM.Translate(E.X[i], E.Y[i])
		j := (count1 / 7) % E.FrameNum[i]
		hx, hy := E.FrameOX[i]+j*E.FrameWidth[i], E.FrameOY[i]
		screen.DrawImage(E.img[i].SubImage(image.Rect(hx, hy, hx+E.FrameWidth[i], hy+E.FrameHeight[i])).(*ebiten.Image), op)
	}
}
