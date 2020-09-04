package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type player struct {
	humanos
	vidas         int
	señalador     int
	señaladorBool bool
	a, b, c, d    int
	MovX          int
	MovY          int
	Inmune        bool
	CountPoder    int
	Coins         int
	Fast          bool
	Compras       bool
	CompleteLevel bool
	enBanco       bool
	enCasita      bool
}

var (
	player1, player2 player
	humano1, humano2 humanos
	// nivel            = int(1)
	// plyrScale float64
	banco     bool
	casita    bool
	countVida int
)

func initPlayer() {

	//////////////   Imangen VIEJO  //////////////////////////////
	humano1.img[0], _, err = ebitenutil.NewImageFromFile(`data/player1 (2).png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	//imagen chic
	humano2.img[0], _, err = ebitenutil.NewImageFromFile(`data/player2 (2).png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	// player1

	player1.FrameOX[0] = 48
	player1.FrameOY[0] = 0
	player1.FrameNum[0] = 1
	player1.X[0] = 396
	player1.Y[0] = 195
	player1.FrameWidth[0] = 48
	player1.FrameHeight[0] = 72
	player1.MovX = 0
	player1.MovY = 0
	//player
	// player1.humanos = player1
	player1.vidas = 3
	player1.señalador = 0
	player1.Coins = 0
	player1.enCasita = true
	player1.a, player1.b, player1.c, player1.d = 0, 0, 0, 0

	player1.humanos.img = humano1.img

	//player2
	player2.FrameOX[0] = 48
	player2.FrameOY[0] = 72
	player2.FrameNum[0] = 1
	player2.X[0] = 430
	player2.Y[0] = 195
	player2.FrameWidth[0] = 48
	player2.FrameHeight[0] = 72
	player1.MovX = 0
	player1.MovY = 0
	player2.enCasita = true

	player2.vidas = 3
	player2.señalador = 1
	player2.Coins = 0
	player2.a, player2.b, player2.c, player2.d = 0, 0, 0, 0
	player2.humanos.img = humano2.img

}
func pasarNivelPlayer() {
	//player1

	player1.X[0] = 396
	player1.Y[0] = 195
	player1.FrameOX[0] = 48
	player1.FrameOY[0] = 0
	player1.FrameNum[0] = 1

	player1.MovX = 0
	player1.MovY = 0
	player1.enCasita = true
	player1.a, player1.b, player1.c, player1.d = 0, 0, 0, 0

	//player2
	player2.FrameNum[0] = 1
	player2.X[0] = 430
	player2.Y[0] = 195
	player2.MovX = 0
	player2.MovY = 0
	player2.enCasita = true

	player2.a, player2.b, player2.c, player2.d = 0, 0, 0, 0

}
func moverPlayer(p player) player {
	// leer tecla
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) && p.señalador == 0 || inpututil.IsKeyJustPressed(ebiten.KeyD) && p.señalador == 1 {
		p.a = 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) && p.señalador == 0 || inpututil.IsKeyJustPressed(ebiten.KeyA) && p.señalador == 1 {
		p.b = 1
	}
	if (!hack && inpututil.IsKeyJustPressed(ebiten.KeyUp)) && p.señalador == 0 || inpututil.IsKeyJustPressed(ebiten.KeyW) && p.señalador == 1 {
		p.c = 1
	}
	if (!hack && inpututil.IsKeyJustPressed(ebiten.KeyDown)) && p.señalador == 0 || inpututil.IsKeyJustPressed(ebiten.KeyS) && p.señalador == 1 {
		p.d = 1
	}
	if p.a == 1 && p.MovY != 1 && p.MovY != 2 {
		p.FrameOX[0] = 0
		p.FrameOY[0] = 144
		p.FrameNum[0] = 3
		p.MovX = 1
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyRight) && p.señalador == 0 || inpututil.IsKeyJustReleased(ebiten.KeyD) && p.señalador == 1 {
		p.FrameOX[0] = 48
		p.FrameNum[0] = 1
		p.MovX = 0
		p.a = 0
	}
	if p.b == 1 && p.MovY != 1 && p.MovY != 2 {
		p.FrameOX[0] = 0
		p.FrameOY[0] = 72
		p.FrameNum[0] = 3
		p.MovX = 2
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyLeft) && p.señalador == 0 || inpututil.IsKeyJustReleased(ebiten.KeyA) && p.señalador == 1 {
		p.FrameOX[0] = 48
		p.FrameNum[0] = 1
		p.MovX = 0
		p.b = 0
	}
	if p.c == 1 && p.MovX != 1 && p.MovX != 2 {
		p.FrameOX[0] = 0
		p.FrameOY[0] = 216
		p.FrameNum[0] = 3
		p.MovY = 1
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyUp) && p.señalador == 0 || inpututil.IsKeyJustReleased(ebiten.KeyW) && p.señalador == 1 {
		p.FrameOX[0] = 48
		p.FrameNum[0] = 1
		p.MovY = 0
		p.c = 0
	}
	if p.d == 1 && p.MovX != 1 && p.MovX != 2 {
		p.FrameOX[0] = 0
		p.FrameOY[0] = 0
		p.FrameNum[0] = 3
		p.MovY = 2
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyDown) && p.señalador == 0 || inpututil.IsKeyJustReleased(ebiten.KeyS) && p.señalador == 1 {
		p.FrameOX[0] = 48
		p.FrameNum[0] = 1
		p.MovY = 0
		p.d = 0
	}

	// trasladar player1
	if Level == 2 && p.posicionInicial[0] != 1 {
		p.posicionInicial[0] = 1
	}
	var X1 = p.X[0]
	var Y1 = p.Y[0]
	switch {
	case p.MovX == 1 && p.Fast:
		p.X[0] += 2.3
	case p.MovX == 2 && p.Fast:
		p.X[0] -= 2.3
	case p.MovY == 1 && p.Fast:
		p.Y[0] -= 2.3
	case p.MovY == 2 && p.Fast:
		p.Y[0] += 2.3

	case p.MovX == 1:
		p.X[0] += 1.5
	case p.MovX == 2:
		p.X[0] -= 1.5
	case p.MovY == 1:
		p.Y[0] -= 1.5
	case p.MovY == 2:
		p.Y[0] += 1.5
	}
	p.X[0], p.Y[0], obs = obstaculos(p.X[0], p.Y[0], X1, Y1)
	if obs {
		p.FrameOX[0] = 48
		p.FrameNum[0] = 1
	}
	///ENTRADAS a puertas///
	switch {
	//banco
	case p.c == 1 && (p.X[0] > 90 && p.X[0] < 103 && p.Y[0] < 89):
		p.Y[0] = 360
		p.X[0] = 507
		sonidoPuerta()
		banco = true
		p.enBanco = true
		interior()

		//vacunatorio
	case Level >= 10 && p.c == 1 && p.X[0] > 814 && p.X[0] < 834 && p.Y[0] < 363 && p.Y[0] > 359:
		p.Y[0] = -40
		p.a, p.b, p.c, p.d = 0, 0, 0, 0
		p.MovX = 0
		p.Compras = true
		vacunatorio = true
		sonidoPuerta()

		//home1
	case p.c == 1 && p.X[0] < 19 && p.Y[0] < 225:

		switch {
		case !p.CompleteLevel:
			p.Y[0] = 379
			p.X[0] = 507
			casita = true
			p.enCasita = true
			interior()
		case p.CompleteLevel:
			p.Y[0] = -40
		}
		sonidoPuerta()
		p.a, p.b, p.c, p.d = 0, 0, 0, 0
		p.MovX = 0

	//pharmacy
	case !banco && !casita && p.c == 1 && p.X[0] > 455 && p.X[0] < 485 && p.Y[0] < 366 && p.Y[0] > 350:
		p.Y[0] = -40
		p.a, p.b, p.c, p.d = 0, 0, 0, 0
		p.MovX = 0
		p.Compras = true
		farmacia = true
		sonidoPuerta()
		//comida china
	case p.c == 1 && p.X[0] > 990 && p.X[0] < 1020 && p.Y[0] < 186 && p.Y[0] > 170:
		p.Y[0] = -40
		p.a, p.b, p.c, p.d = 0, 0, 0, 0
		p.MovX = 0
		p.Compras = true
		bakery = true
		sonidoPuerta()
		//PARA MART
	case p.c == 1 && p.X[0] > 310 && p.X[0] < 330 && p.Y[0] < 52 && p.Y[0] > 46:
		p.Y[0] = -40
		p.a, p.b, p.c, p.d = 0, 0, 0, 0
		p.MovX = 0
		sonidoPuerta()
		p.Compras = true
		mart = true
		//SUPERMAKET
	case p.c == 1 && (p.X[0] > 795 && p.X[0] < 855 && p.Y[0] < 186 && p.Y[0] > 171):
		p.Y[0] = -80
		p.a, p.b, p.c, p.d = 0, 0, 0, 0
		p.MovX = 0
		p.Compras = true
		supermarket = true
		sonidoPuerta()

		///SALIDAS///
		// banco
	case p.Y[0] > 361 && p.X[0] > 480 && p.X[0] < 542 && banco:
		p.Y[0] = 90
		p.X[0] = 95
		sonidoPuerta()
		salida()
		banco = false
		p.enBanco = false
		//home1
	case p.Y[0] > 380 && p.X[0] > 480 && p.X[0] < 542 && casita:
		p.Y[0] = 226
		p.X[0] = 10
		p.enCasita = false
		if Game1.numPlayers == 2 && player1.enCasita == false || player2.enCasita == false {
			sonidoPuerta()
			salida()
			casita = false

		}
		if Game1.numPlayers == 1 {
			sonidoPuerta()
			salida()
			casita = false
		}

	//vacunatorio
	case p.Y[0] < -36 && p.Y[0] > -39 && p.X[0] > 814 && p.X[0] < 834:
		p.Y[0] = 364
		sonidoPuerta()
		//mart
	case p.Y[0] < -36 && p.Y[0] > -39 && p.X[0] > 310 && p.X[0] < 330:
		p.Y[0] = 53
		mart = false
		sonidoPuerta()

		//Pharmacy
	case p.Y[0] < -36 && p.Y[0] > -39 && p.X[0] > 454 && p.X[0] < 486:
		p.Y[0] = 367
		farmacia = false
		sonidoPuerta()
		//salida de Comida China
	case p.Y[0] < -36 && p.Y[0] > -39 && p.X[0] > 990 && p.X[0] < 1020:
		p.Y[0] = 187
		bakery = false
		sonidoPuerta()
		//supermarket
	case p.Y[0] < -76 && p.Y[0] > -79 && (p.X[0] > 795 && p.X[0] < 856):
		p.Y[0] = 187
		supermarket = false
		sonidoPuerta()

	}

	//Vuelta a la realidad
	if p.Y[0] < 0 {
		p.X[0] = X1
	}
	if p.Y[0] < -40 && p.Y[0] > -60 {
		p.Y[0] = -40
	}
	//porque coincide la salida del vacunatorio con la del super esto es necesario
	if p.Y[0] < -80 {
		p.Y[0] = -80
	}

	if inconcebible {
		p.X[0] = 396
		p.Y[0] = 195
		p.FrameOX[0] = 48
		p.FrameOY[0] = 0
		p.FrameNum[0] = 1
		p.MovX = 0
		p.MovY = 0
		p.enCasita = true
		inconcebible = false
	}
	return p
}

func vida(h humanos, p player, b Objetos, pl Objetos) (player, Objetos, Objetos) {
	barHScale = float64(barbijo.FrameHeight) * objScale
	barWscale = float64(barbijo.FrameWidth) * objScale
	coinHScale = float64(monedas.FrameHeight) * objScale
	coinWscale = float64(monedas.FrameWidth) * objScale
	alcholHScale = float64(alchol.FrameHeight) * objScale
	alcholWScale = float64(alchol.FrameWidth) * objScale
	plasmaHScale = float64(plasma.FrameHeight) * objScale
	plasmaWScale = float64(plasma.FrameWidth) * objScale

	if !p.Inmune && countVida > 60 {

		//pierde vidas con la nube
		for i := 0; i < numNube; i++ {
			nubX := nube1.X[i] * nubScale
			nubY := nube1.Y[i] * nubScale
			nubFrameWidth := nubFrameWidth * nubScale
			nubFrameHight := nubFrameHight * nubScale

			if countVida > 90 && p.X[0]+wth > nubX+35 && p.X[0] < nubX+nubFrameWidth && p.Y[0]+hgt > nubY+20 && p.Y[0] < nubY+nubFrameHight && nube1.Alpha[i] > .4 {
				p.vidas--
				countVida = 0
				sonidoVidas(p)
			}
		}

		//pierde vidas con humanos
		for i := randNum; i < numEnemigo+randNum; i++ {
			if p.X[0]+(wth-5) > h.X[i]+5 && p.X[0]+5 < h.X[i]+(wth-5) && p.Y[0]+hgt > h.Y[i]+(hgt-8) && p.Y[0]+(hgt-8) < h.Y[i]+hgt {
				p.vidas--
				countVida = 0
				sonidoVidas(p)
			}
		}

	}

	//inmune con barbijo o alchol en gel
	if !casita && !banco && p.X[0]+wth > b.X && p.X[0] < b.X+barWscale+20 && p.Y[0]+hgt-45 > b.Y && p.Y[0]+hgt-45 < b.Y+barHScale {
		sBarbijo.Play()
		sBarbijo.Rewind()
		b.X = 1500
		p.Inmune = true
		p.CountPoder = 1200
	}

	if b == alchol && p.X[0]+wth > b.X && p.X[0] < b.X+barWscale+20 && p.Y[0]+hgt-45 > b.Y && p.Y[0]+hgt-45 < b.Y+barHScale {
		sBarbijo.Play()
		sBarbijo.Rewind()
		b.X = 1500
		p.Inmune = true
		p.CountPoder = 1200
	}

	if p.Inmune || p.Fast {

		p.CountPoder--
	}
	if p.CountPoder == 0 {
		p.Inmune = false
		p.Fast = false
	}

	//gana/pierde vida
	if !casita && !banco && p.X[0]+wth > pl.X && p.X[0] < pl.X+plasmaWScale+20 && p.Y[0]+hgt > pl.Y && p.Y[0]+hgt < pl.Y+plasmaHScale+20 {
		sBarbijo.Play()
		sBarbijo.Rewind()
		p.vidas++
		pl.X = 1500
	}

	if p.vidas == 0 {
		ModeGameOver = true
		ModeGame = false
		banco = false
	}
	//PIERDE POR falta de plata
	if player1.Coins < 2 && player2.Coins < 2 && !p.CompleteLevel && monedas.X == 1500 {
		ModeGame = false
		ModeGameOver = true
	}
	//gana monedas
	if p.X[0]+wth > monedas.X && p.X[0] < monedas.X+coinWscale && p.Y[0]+hgt > monedas.Y+40 && p.Y[0]+hgt < monedas.Y+40+coinHScale {
		monedas.X = 1500
		p.Coins += 5
		sonidomonedas()
	}
	return p, b, pl
}

func dibujarPlayer(P player, screen *ebiten.Image) {
	if ModePause {
		P.FrameNum[0] = 1
		P.FrameOX[0] = 48
	}
	op := &ebiten.DrawImageOptions{}
	if !ModeTitle {
		op.GeoM.Scale(hScaleW, hScaleH)
	}
	op.GeoM.Translate(P.X[0], P.Y[0])
	i := (count1 / 7) % P.FrameNum[0]
	sx, sy := P.FrameOX[0]+i*P.FrameWidth[0], P.FrameOY[0]
	screen.DrawImage(P.img[0].SubImage(image.Rect(sx, sy, sx+P.FrameWidth[0], sy+P.FrameHeight[0])).(*ebiten.Image), op)
}
