package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Objetos struct {
	FrameOX     int
	FrameOY     int
	FrameNum    int
	FrameWidth  int
	FrameHeight int
	X           float64
	Y           float64
	img         *ebiten.Image
}

var (
	barbijo, alchol, plasma, fondoNegro, fondoNegroVidas1, fondoNegroVidas2, fondoNegroInmune1, fondoNegroInmune2, fondoNegroFast1, fondoNegroFast2, fondoNegroCommans, fondoNegroCompras, fondoNegroPause, home1, home, monedas, relato, papiro, ciudad, tpaper, vaccine, money, meds, mmeds, mhome, mhome1, bread, clothes Objetos
	objScale                                                                                                                                                                                                                                                                                                                 = .3
	barHScale                                                                                                                                                                                                                                                                                                                float64
	barWscale                                                                                                                                                                                                                                                                                                                float64
	coinHScale                                                                                                                                                                                                                                                                                                               float64
	coinWscale                                                                                                                                                                                                                                                                                                               float64
	plasmaHScale                                                                                                                                                                                                                                                                                                             float64
	plasmaWScale                                                                                                                                                                                                                                                                                                             float64
	alcholHScale                                                                                                                                                                                                                                                                                                             float64
	alcholWScale                                                                                                                                                                                                                                                                                                             float64
)

//cartFarmacy, cartSupermarket, cartStore, cartBanck

func initObjetos() {
	//objetos
	barbijo.FrameOX = 0
	barbijo.FrameOY = 62
	barbijo.FrameNum = 1
	barbijo.FrameWidth = 38
	barbijo.FrameHeight = 17
	barbijo.X = float64(650)
	barbijo.Y = float64(150)
	barbijo.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\objetos.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	plasma.FrameOX = 0
	plasma.FrameOY = 79
	plasma.FrameNum = 1
	plasma.FrameWidth = 30
	plasma.FrameHeight = 48
	plasma.X = float64(489)
	plasma.Y = float64(73)
	plasma.img = barbijo.img

	alchol.FrameOX = 0
	alchol.FrameOY = 124
	alchol.FrameNum = 1
	alchol.FrameWidth = 30
	alchol.FrameHeight = 39
	alchol.X = float64(353)
	alchol.Y = float64(260)
	alchol.img = barbijo.img

	monedas.FrameOX = 0
	monedas.FrameOY = 170
	monedas.FrameNum = 1
	monedas.FrameWidth = 24
	monedas.FrameHeight = 30
	monedas.X = float64(430)
	monedas.Y = float64(125)
	monedas.img = barbijo.img

	//carteles
	fondoNegro.FrameOX = 344
	fondoNegro.FrameOY = 876
	fondoNegro.FrameNum = 1
	fondoNegro.FrameWidth = 905
	fondoNegro.FrameHeight = 460
	fondoNegro.X = float64(305)
	fondoNegro.Y = float64(230)
	fondoNegro.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\mission-fondo.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	fondoNegroPause.FrameOX = 344
	fondoNegroPause.FrameOY = 145
	fondoNegroPause.FrameNum = 1
	fondoNegroPause.FrameWidth = 435
	fondoNegroPause.FrameHeight = 155
	fondoNegroPause.X = float64(300)
	fondoNegroPause.Y = float64(150)
	fondoNegroPause.img = fondoNegro.img

	fondoNegroVidas1.FrameOX = 0
	fondoNegroVidas1.FrameOY = 485
	fondoNegroVidas1.FrameNum = 1
	fondoNegroVidas1.FrameWidth = 435
	fondoNegroVidas1.FrameHeight = 155
	fondoNegroVidas1.X = float64(0)
	fondoNegroVidas1.Y = float64(470)
	fondoNegroVidas1.img = fondoNegro.img

	fondoNegroVidas2.FrameOX = 0
	fondoNegroVidas2.FrameOY = 485
	fondoNegroVidas2.FrameNum = 1
	fondoNegroVidas2.FrameWidth = 435
	fondoNegroVidas2.FrameHeight = 155
	fondoNegroVidas2.X = float64(750)
	fondoNegroVidas2.Y = float64(470)
	fondoNegroVidas2.img = fondoNegro.img

	fondoNegroInmune1.FrameOX = 0
	fondoNegroInmune1.FrameOY = 500
	fondoNegroInmune1.FrameNum = 1
	fondoNegroInmune1.FrameWidth = 435
	fondoNegroInmune1.FrameHeight = 50
	fondoNegroInmune1.X = float64(0)
	fondoNegroInmune1.Y = float64(-10)
	fondoNegroInmune1.img = fondoNegro.img

	fondoNegroInmune2.FrameOX = 0
	fondoNegroInmune2.FrameOY = 500
	fondoNegroInmune2.FrameNum = 1
	fondoNegroInmune2.FrameWidth = 435
	fondoNegroInmune2.FrameHeight = 50
	fondoNegroInmune2.X = float64(745)
	fondoNegroInmune2.Y = float64(-10)
	fondoNegroInmune2.img = fondoNegro.img

	fondoNegroFast1.FrameOX = 0
	fondoNegroFast1.FrameOY = 500
	fondoNegroFast1.FrameNum = 1
	fondoNegroFast1.FrameWidth = 435
	fondoNegroFast1.FrameHeight = 50
	fondoNegroFast1.X = float64(200)
	fondoNegroFast1.Y = float64(-10)
	fondoNegroFast1.img = fondoNegro.img

	fondoNegroFast2.FrameOX = 0
	fondoNegroFast2.FrameOY = 500
	fondoNegroFast2.FrameNum = 1
	fondoNegroFast2.FrameWidth = 435
	fondoNegroFast2.FrameHeight = 50
	fondoNegroFast2.X = float64(510)
	fondoNegroFast2.Y = float64(-10)
	fondoNegroFast2.img = fondoNegro.img

	fondoNegroCommans.FrameOX = 0
	fondoNegroCommans.FrameOY = 0
	fondoNegroCommans.FrameNum = 1
	fondoNegroCommans.FrameWidth = 600
	fondoNegroCommans.FrameHeight = 500
	fondoNegroCommans.X = float64(180)
	fondoNegroCommans.Y = float64(0)
	fondoNegroCommans.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\fondoNegro.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	fondoNegroCompras.FrameOX = 0
	fondoNegroCompras.FrameOX = 0
	fondoNegroCompras.FrameOY = 0
	fondoNegroCompras.FrameNum = 1
	fondoNegroCompras.FrameWidth = 1100
	fondoNegroCompras.FrameHeight = 300
	fondoNegroCompras.X = float64(30)
	fondoNegroCompras.Y = float64(100)
	fondoNegroCompras.img, _, err = ebitenutil.NewImageFromFile(`sircovid\data\fondoNegro1.png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	mhome.FrameOX = 0
	mhome.FrameOY = 202
	mhome.FrameNum = 1
	mhome.FrameWidth = 144
	mhome.FrameHeight = 287
	mhome.X = float64(0)
	mhome.Y = float64(200)
	mhome.img = fondoNegro.img

	money.FrameOX = 0
	money.FrameOY = 600
	money.FrameNum = 1
	money.FrameWidth = 290
	money.FrameHeight = 232
	money.X = float64(50)
	money.Y = float64(50)
	money.img = fondoNegro.img

	meds.FrameOX = 297
	meds.FrameOY = 310
	meds.FrameNum = 1
	meds.FrameWidth = 300
	meds.FrameHeight = 180
	meds.X = float64(315)
	meds.Y = float64(310)
	meds.img = fondoNegro.img

	bread.FrameOX = 975
	bread.FrameOY = 140
	bread.FrameNum = 1
	bread.FrameWidth = 140
	bread.FrameHeight = 257
	bread.X = float64(950)
	bread.Y = float64(145)
	bread.img = fondoNegro.img

	clothes.FrameOX = 278
	clothes.FrameOY = 0
	clothes.FrameNum = 1
	clothes.FrameWidth = 387
	clothes.FrameHeight = 142
	clothes.X = float64(270)
	clothes.Y = float64(0)
	clothes.img = fondoNegro.img

	tpaper.FrameOX = 794
	tpaper.FrameOY = 610
	tpaper.FrameNum = 1
	tpaper.FrameWidth = 210
	tpaper.FrameHeight = 259
	tpaper.X = float64(750)
	tpaper.Y = float64(35)
	tpaper.img = fondoNegro.img

	vaccine.FrameOX = 590
	vaccine.FrameOY = 312
	vaccine.FrameNum = 1
	vaccine.FrameWidth = 360
	vaccine.FrameHeight = 182
	vaccine.X = float64(600)
	vaccine.Y = float64(335)
	vaccine.img = fondoNegro.img
}

func dibujarNiveles(screen *ebiten.Image) {
	if ModeTitleLevel || (ModeGame && monedas.X != 1500) {
		dibujarObjetos(money, screen)
	}
	if ModeTitleLevel || (ModeGame && (player1.CompleteLevel || player2.CompleteLevel)) {
		dibujarObjetos(mhome, screen)
	}

	if (ModeTitleLevel || ((player1.Coins > 2 || player2.Coins > 2) && !farmacia1)) && (Level == 1 || Level == 5 || Level == 6 || Level == 8) {
		dibujarObjetos(meds, screen)

	}
	if (ModeTitleLevel || ((player1.Coins > 2 || player2.Coins > 2) && !bakery1)) && (Level == 2 || Level == 5 || Level == 7 || Level == 10) {
		dibujarObjetos(bread, screen)

	}
	if (ModeTitleLevel || ((player1.Coins > 2 || player2.Coins > 2) && !mart1)) && (Level == 3 || Level == 7 || Level == 8 || Level == 9) {
		dibujarObjetos(clothes, screen)

	}
	if (ModeTitleLevel || ((player1.Coins > 2 || player2.Coins > 2) && !supermarket1)) && (Level == 4 || Level == 6 || Level == 9 || Level == 10) {
		dibujarObjetos(tpaper, screen)

	}
	if !ModeWin && (ModeTitleLevel || (player1.Coins > 2 || player2.Coins > 2)) && Level > 10 {
		dibujarObjetos(vaccine, screen)

	}
}

func dibujarObjetos(B Objetos, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	if B == mhome || B == money || B == meds || B == bread || B == clothes || B == tpaper || B == vaccine || B == monedas {
		op.GeoM.Scale(.8, .8)
	}
	op.GeoM.Translate(B.X, B.Y)
	bx, by := B.FrameOX, B.FrameOY
	screen.DrawImage(B.img.SubImage(image.Rect(bx, by, bx+B.FrameWidth, by+B.FrameHeight)).(*ebiten.Image), op)
}
