package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// init carga los datos
func init() {

	//init ciudad img
	imgTiles, _, err = ebitenutil.NewImageFromFile(`data/ciudad Fere sat1(rezised).png`, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	//init intro
	intro1.initIntro(screenWidth, screenHeight)
	//inicia los textos
	initTextos()
	//inicializa a players
	initPlayer()
	//inicia sumarVidas
	initObjetos()
	//inicia nube
	initNube()
	//inicia enemigos
	initEnemigos()
	//iniciar otra variables
	iniciarVariables()
	//iniciar sonidos
	initSonido()

}

////////////////////////////
//////// Update ////////////
////////////////////////////

// Update se llama 60 veces por segundo
func (g *Game) Update(screen *ebiten.Image) error {
	// game counter
	comandos()
	g.count++
	count1++
	countVida++
	if count1 == 80 {
		count1 = 0
	}
	//func sonido
	sonido(player1)
	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		Commands = !Commands
		if Credits {
			Credits = !Credits
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		Credits = !Credits
		if Commands {
			Commands = !Commands
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		Commands = false
		Credits = false
	}
	switch {
	//mustra los creditos
	case Credits:
		//muestra ayuda
	case Commands:
	//pausar el juego

	case ModePause:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			ModePause = !ModePause
			// arreglo bug
		}
	case ModeWin:
	//toda la introduccion con eleccion de players, etc

	//introduccion al juego
	case ModeTitle:
		introduccion()
		//escribe en que nivel estas
	case ModeTitleLevel:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			ModeTitleLevel = false
			ModeGame = true
			casita = true
			interior()
		}
	//te muestra cual es tu mision

	case player1.Compras:
		if player1.Coins >= 2 {
			player1 = compar(player1)
		}
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			player1.Compras = false
		}
	case player2.Compras:
		if player2.Coins >= 2 {
			player2 = compar(player2)
		}
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			player2.Compras = false
		}
		//Game estas jugando
	case ModeGame:
		// casita = true
		//para mode pause y muestra los niveles
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			ModePause = !ModePause
		}
		// nube
		if !banco || !casita {
			nube1 = moverNube(nube1)
		}

		// palyer
		player1 = moverPlayer(player1)
		player2 = moverPlayer(player2)

		// vida
		player1, barbijo, plasma = vida(enemigo, player1, barbijo, plasma)
		player1, alchol, plasma = vida(enemigo, player1, alchol, plasma)

		if Game1.numPlayers == 2 {
			player2, barbijo, plasma = vida(enemigo, player2, barbijo, plasma)
			player2, alchol, plasma = vida(enemigo, player2, alchol, plasma)
		}
		//enemigos
		enemigo = moverHumanos(enemigo)

		//siguiente nivel
		Game1.siguienteNivel = siguienteNivel(player1)
		Game1.siguienteNivel = siguienteNivel(player2)

	case ModeGameOver:
		for i := 0; i < nivel; i++ {
			enemigo.FrameNum[i] = 1
			enemigo.FrameOX[i] = 0
		}
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			iniciarVariables()
			initPlayer()
			initNube()
			initObjetos()
			initEnemigos()
			banco = false
			ModeGameOver = false
		}
	}
	return nil

}

////////////////////////////
///////// Draw /////////////
////////////////////////////

// Draw dibuja la pantalla 60 veces por segundo
func (g *Game) Draw(screen *ebiten.Image) {
	// dibujar fondo
	op := &ebiten.DrawImageOptions{}

	switch {
	case Credits:
		dibujarObjetos(fondoNegroCommans, screen)
		dibujarTextos(screen)
	case Commands:
		dibujarObjetos(fondoNegroCommans, screen)
		dibujarTextos(screen)

	case banco:
		screen.DrawImage(imgBanco, op)
		dibujarObjetos(fondoNegroVidas1, screen)

		dibujarObjetos(monedas, screen)
		dibujarEnemigos(enemigo, screen)
		if player1.enBanco == true {
			dibujarPlayer(player1, screen)
		}
		if player2.enBanco == true {
			if Game1.numPlayers == 2 {
				dibujarPlayer(player2, screen)
				dibujarObjetos(fondoNegroVidas2, screen)

			}
		}
		screen.DrawImage(imgCintas, op)
		dibujarObjetos(alchol, screen)
		dibujarTextos(screen)

	case casita && !ModePause:
		screen.DrawImage(imgCasita, op)
		dibujarObjetos(fondoNegroVidas1, screen)

		//dibujar palyers
		if player1.enCasita == true {
			dibujarPlayer(player1, screen)
		}

		if player2.enCasita == true && Game1.numPlayers == 2 {
			dibujarObjetos(fondoNegroVidas2, screen)
			dibujarPlayer(player2, screen)
		}
		dibujarTextos(screen)

	default:
		screen.DrawImage(imgTiles, op)

		//dibuja al enemigo
		dibujarEnemigos(enemigo, screen)

		if !ModeTitleLevel && !ModePause {
			dibujarPlayer(player1, screen)
			if Game1.numPlayers == 2 {
				dibujarObjetos(fondoNegroVidas2, screen)
				dibujarPlayer(player2, screen)
			}
		}
		// dibujar nube
		dibujarNube(nube1, screen)

		dibujarObjetos(fondoNegroVidas1, screen)

		//dibujar objetos
		dibujarObjetos(barbijo, screen)
		dibujarObjetos(plasma, screen)

		if ModeTitle {
			dibujarObjetos(fondoNegro, screen)
			dibujarPlayer(player1, screen)
			if Game1.numPlayers == 2 {
				dibujarObjetos(fondoNegroVidas2, screen)
				dibujarPlayer(player2, screen)
			}
		}

		if ModePause || ModeTitleLevel {
			dibujarObjetos(fondoNegroPause, screen)

		}

		if player1.Compras || player2.Compras || ModeGameOver || ModeWin {
			dibujarObjetos(fondoNegroCompras, screen)
		}
		//dibujar textos compras
		dibujarTextoCompras(player1, screen)
		dibujarTextoCompras(player2, screen)
		// // dibujar nube
		// dibujarNube(nube1, screen)

		if ModeTitleLevel || ModePause {
			dibujarNiveles(screen)
			if Game1.numPlayers == 2 {
				dibujarObjetos(fondoNegroVidas2, screen)
			}
		}
		if ModeGame && count1 < 60 && !banco && !casita {
			dibujarNiveles(screen)
			if Game1.numPlayers == 2 {
				dibujarObjetos(fondoNegroVidas2, screen)
			}
		}
	}
	if player1.Inmune {
		dibujarObjetos(fondoNegroInmune1, screen)
	}
	if player2.Inmune {
		dibujarObjetos(fondoNegroInmune2, screen)
	}
	if player1.Fast {
		dibujarObjetos(fondoNegroFast1, screen)
	}
	if player2.Fast {
		dibujarObjetos(fondoNegroFast2, screen)
	}
	dibujarTextos(screen)
}

// Layout maneja las dimensiones de pantalla
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(screenWidth), int(screenHeight)
}

////////////////////////////
// Main
////////////////////////////

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Sir-covid")
	ebiten.SetWindowResizable(true)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
