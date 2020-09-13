package main

import (
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Intro logo y adornos
type intro struct {

	// logo
	logoX        int
	logoWidth    int
	logoY        float64
	logoHeight   int
	logoScale    float64
	logoScaleDir string
	logoAlpha    float64
	logoAlphaDir string
	imgLogo      *ebiten.Image

	// covids
	imgCovid       *ebiten.Image
	imgCovidWidth  int
	imgCovidHeight int

	covidX          [20]float64
	covidY          [20]float64
	covidSpeed      [20]float64
	covidAngle      [20]float64
	covidAngleSpeed [20]float64
	covidScale      [20]float64
	covidAlpha      [20]float64
}

func (i *intro) initIntro(screenWidth int, screenHeight int) {
	// sIntro.Play()

	// logo
	i.logoX = screenWidth / 2
	i.logoY = float64(0)
	i.logoScale = float64(1)
	i.logoScaleDir = "up"
	i.logoAlpha = float64(0)
	i.logoAlphaDir = "up"

	i.imgLogo, _, err = ebitenutil.NewImageFromFile(`data/sircovid.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	i.logoWidth, i.logoHeight = i.imgLogo.Size()

	// covids
	i.imgCovid, _, err = ebitenutil.NewImageFromFile(`data/covid.png`, ebiten.FilterDefault)
	i.imgCovidWidth, i.imgCovidHeight = i.imgCovid.Size()
	if err != nil {
		log.Fatal(err)
	}

	for j := 0; j < 20; j++ {
		i.covidY[j] = float64(rand.Intn(screenHeight))
		i.covidSpeed[j] = (rand.Float64() - 0.5)
		i.covidAngle[j] = float64(rand.Intn(360))
		i.covidAngleSpeed[j] = (rand.Float64() - 0.5) / 10 // 0.01
		i.covidScale[j] = 0.1 + rand.Float64()*1.5
		i.covidAlpha[j] = rand.Float64()
		if i.covidSpeed[j] > 0 {
			i.covidX[j] = float64(0 - i.imgCovidWidth)
		} else {
			i.covidX[j] = float64(screenWidth + i.imgCovidWidth)
		}
	}
}

func (i *intro) updateIntro(screenWidth int, screenHeight int) {

	// logo moves down
	if i.logoY < float64(screenHeight/5) {
		i.logoY += 0.3
	}

	// logo scale
	if i.logoScaleDir == "up" {
		i.logoScale += 0.001
	}
	if i.logoScaleDir == "down" {
		i.logoScale -= 0.001
	}
	if i.logoScale < 1.4 {
		i.logoScaleDir = "up"
	}
	if i.logoScale > 1.6 {
		i.logoScaleDir = "down"
	}

	// logo alpha
	if i.logoAlphaDir == "up" {
		i.logoAlpha += 0.01
	}
	if i.logoAlphaDir == "down" {
		i.logoAlpha -= 0.01
	}
	if i.logoAlpha < 0.9 {
		i.logoAlphaDir = "up"
	}
	if i.logoAlpha > 1.1 {
		i.logoAlphaDir = "down"
	}

	// covids
	for j := 0; j < 20; j++ {
		i.covidX[j] += i.covidSpeed[j]
		i.covidAngle[j] += i.covidAngleSpeed[j]

		if i.covidX[j] < float64(-i.imgCovidWidth) || i.covidX[j] > float64(screenWidth+i.imgCovidWidth) {
			i.covidY[j] = float64(rand.Intn(screenHeight))
			i.covidSpeed[j] = (rand.Float64() - 0.5)
			i.covidAngle[j] = float64(rand.Intn(360))
			i.covidAngleSpeed[j] = (rand.Float64() - 0.5) / 10 // 0.01
			i.covidScale[j] = 0.1 + rand.Float64()*1.5
			i.covidAlpha[j] = rand.Float64()
			if i.covidSpeed[j] > 0 {
				i.covidX[j] = float64(0 - i.imgCovidWidth)
			} else {
				i.covidX[j] = float64(screenWidth + i.imgCovidWidth)
			}
		}
	}
}

func (i *intro) drawIntro(screen *ebiten.Image, screenWidth int, screenHeight int) {

	op := &ebiten.DrawImageOptions{}

	// covids
	for j := 0; j < 20; j++ {
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(-i.imgCovidWidth/2), float64(-i.imgCovidHeight/2))
		op.GeoM.Scale(i.covidScale[j], i.covidScale[j])
		op.GeoM.Rotate(i.covidAngle[j])
		op.GeoM.Translate(float64(i.covidX[j]), float64(i.covidY[j]))
		op.ColorM.Scale(1, 1, 1, i.covidAlpha[j])
		screen.DrawImage(i.imgCovid, op)
	}

	// logo
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(i.logoX-i.logoWidth/2), i.logoY-float64(i.logoHeight/2))
	op.GeoM.Scale(1, i.logoScale)
	op.ColorM.Scale(1, 1, 1, i.logoAlpha)
	screen.DrawImage(i.imgLogo, op)
}
