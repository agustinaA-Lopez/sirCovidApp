package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

//Game es la estructura del juego
type Game struct {
	count int
	nube
	numPlayers     int
	electPlayer    int
	siguienteNivel (player)
}

// Game1 es el juego
var Game1 Game
var intro1 intro

var (
	Commands                  bool
	Credits                   bool
	ModeGame                  bool
	ModeTitleLevel            bool
	Level                     = 1
	ModeWin                   bool
	ModeTitle                 bool
	ElectNumPlayers           int
	ElectPlayer               int
	ModeGameOver              bool
	count1                    int
	ModePause                 bool
	elecCompras               int
	farmacia, farmacia1       bool
	mart, mart1               bool
	bakery, bakery1           bool
	supermarket, supermarket1 bool
	vacunatorio               bool
	senia                     bool
	inconcebible              bool

	// imágenes
	imgTiles  *ebiten.Image
	imgBanco  *ebiten.Image
	imgCintas *ebiten.Image
	imgCasita *ebiten.Image

	//para start y game over
	arcadeFont      font.Face
	smallArcadeFont font.Face
	texts           = []string{}

	err error
)

const (
	// game
	screenWidth  = 66 * 16
	screenHeight = 33 * 16

	// tiles
	tileSize = 16
	tileXNum = 48

	//para start y game Over
	fontSize      = 32
	smallFontSize = fontSize / 2
)

//introduccion es la eleccion de los players, etc
func introduccion() {
	// intro update
	intro1.updateIntro(screenWidth, screenHeight)
	player1.X[0] = float64(390)
	player1.Y[0] = float64(350)
	player2.X[0] = float64(605)
	player2.Y[0] = float64(350)

	switch {
	case ElectNumPlayers == 0:
		if !hack && inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			Game1.numPlayers = 2
		}
		if !hack && inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			Game1.numPlayers = 1
		}
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) && (Game1.numPlayers == 1 || Game1.numPlayers == 2) {
			ElectNumPlayers = 1
		}
	case ElectPlayer == 0 && Game1.numPlayers == 1 || Game1.numPlayers == 2:
		if Game1.numPlayers == 1 || Game1.numPlayers == 2 {
			if !hack && inpututil.IsKeyJustPressed(ebiten.KeyUp) {
				player1.humanos.img = humano1.img
				player1.señaladorBool = false
			}
			if !hack && inpututil.IsKeyJustPressed(ebiten.KeyDown) {
				player1.humanos.img = humano2.img
				player1.señaladorBool = true
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyW) {
				player2.humanos.img = humano1.img
				player2.señaladorBool = false

			}
			if inpututil.IsKeyJustPressed(ebiten.KeyS) {
				player2.humanos.img = humano2.img
				player2.señaladorBool = true
			}
		}
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			ElectPlayer = 1
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && ElectPlayer == 1 {
		ModeTitle = false
		player1.X[0] = 396
		player1.Y[0] = 195
		player2.X[0] = 430
		player2.Y[0] = 195
		ModeTitleLevel = true
	}

}

func siguienteNivel(p player) player {
	if p.CompleteLevel && (p.X[0] >= home.X-40 && p.X[0] <= home.X+40 && p.Y[0] == -40) {
		pasarNivel()
		player1.CompleteLevel = false
		player2.CompleteLevel = false
		farmacia, mart, supermarket, bakery, vacunatorio = false, false, false, false, false
	}
	return p
}
func compar(p player) player {
	//compras

	if (inpututil.IsKeyJustPressed(ebiten.KeyDown) && p.señalador == 0) || (inpututil.IsKeyJustPressed(ebiten.KeyS) && p.señalador == 1) {
		elecCompras++
	}
	if (inpututil.IsKeyJustPressed(ebiten.KeyUp) && p.señalador == 0) || (inpututil.IsKeyJustPressed(ebiten.KeyW) && p.señalador == 1) {
		elecCompras--
	}
	if elecCompras > 1 && (bakery || vacunatorio) {
		elecCompras = 1
	}

	if elecCompras > 2 {
		elecCompras = 2
	}
	if elecCompras == 0 && vacunatorio && p.Coins < 10 {
		elecCompras = 1
	}
	if elecCompras < 0 {
		elecCompras = 0
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if (farmacia || supermarket) && elecCompras == 0 {
			p.Coins = p.Coins - 3
			sBarbijo.Play()
			sBarbijo.Rewind()
			p.vidas++
		}
		if mart && elecCompras == 1 {
			p.Coins = p.Coins - 5
			p.X[0] = 396
			p.Y[0] = 195
			inconcebible = true
			pasarNivel()
		}
		if mart && elecCompras == 0 {
			p.Coins = p.Coins - 2
			p.Inmune = true
			p.CountPoder = 600
		}
		if (elecCompras == 1 && (farmacia || supermarket)) || (elecCompras == 0 && bakery) {
			p.Coins = p.Coins - 2
			p.Fast = true
			p.CountPoder = 600
			sFast.Play()
			sFast.Rewind()
		}
		if elecCompras == 0 && vacunatorio {
			ModeWin = true
		}

		switch {
		case elecCompras == 1 && bakery:
			p.Coins = p.Coins - 2
			bakery1 = true
		case elecCompras == 2 && farmacia:
			p.Coins = p.Coins - 2
			farmacia1 = true
		case elecCompras == 2 && mart:
			p.Coins = p.Coins - 2
			mart1 = true
		case elecCompras == 2 && supermarket:
			p.Coins = p.Coins - 2
			supermarket1 = true
		}
		switch {
		case farmacia1 && Level == 1:
			p.CompleteLevel = true
		case bakery1 && Level == 2:
			p.CompleteLevel = true
		case mart1 && Level == 3:
			p.CompleteLevel = true
		case supermarket1 && Level == 4:
			p.CompleteLevel = true
		case farmacia1 && bakery1 && Level == 5:
			p.CompleteLevel = true

		case farmacia1 && supermarket1 && Level == 6:
			p.CompleteLevel = true
		case mart1 && bakery1 && Level == 7:
			p.CompleteLevel = true
		case farmacia1 && mart1 && Level == 8:
			p.CompleteLevel = true
		case supermarket1 && mart1 && Level == 9:
			p.CompleteLevel = true
		case supermarket1 && bakery1 && Level == 10:
			p.CompleteLevel = true
		case vacunatorio && elecCompras == 1 && Level > 10:
			p.CompleteLevel = true

		}

		farmacia, mart, supermarket, bakery = false, false, false, false
		elecCompras = 0
		p.Compras = false

	}
	return p
}
func comandos() {
	// coder shortcuts
	if inpututil.IsKeyJustPressed(ebiten.KeyControl) {
		hack = true
	} else if inpututil.IsKeyJustReleased(ebiten.KeyControl) {
		hack = false
	}
	switch {
	case hack && inpututil.IsKeyJustPressed(ebiten.KeyF):
		player1.Fast = !player1.Fast

	case hack && inpututil.IsKeyJustPressed(ebiten.KeyN):
		banco = false
		casita = false
		pasarNivel()

	case hack && inpututil.IsKeyJustPressed(ebiten.KeyI):
		player1.Inmune = !player1.Inmune

	case hack && inpututil.IsKeyJustPressed(ebiten.KeyM):
		player1.Coins += 10
	}
}
func dibujarTextoCompras(p player, screen *ebiten.Image) {

	if p.Compras {
		if p.Coins < 2 {
			jugadores := fmt.Sprintf("YOU HAVE no MONEY\n  COME BACK SOON")
			text.Draw(screen, jugadores, arcadeFont, 235, 250, color.White)
		}

		//EN FARMACIA
		switch {
		case farmacia && elecCompras == 0 && p.Coins >= 2:
			jugadores := fmt.Sprintf(">$3-PLASMA -GET LIFE-\n $2-ASPIRIN -GO FAST-\n $2-MEDICINE")
			text.Draw(screen, jugadores, arcadeFont, 220, 250, color.White)
		case farmacia && elecCompras == 1 && p.Coins >= 2:
			jugadores := fmt.Sprintf(" $3-PLASMA -GET LIFE-\n>$2-ASPIRIN -GO FAST-\n $2-MEDICINE")
			text.Draw(screen, jugadores, arcadeFont, 220, 250, color.White)
		case farmacia && elecCompras == 2 && p.Coins >= 2:
			jugadores := fmt.Sprintf(" $3-PLASMA -GET LIFE-\n $2-ASPIRIN -GO FAST-\n>$2-MEDICINE")
			text.Draw(screen, jugadores, arcadeFont, 220, 250, color.White)

			//EN Comida China
		case bakery && elecCompras == 0 && p.Coins >= 2:
			jugadores := fmt.Sprintf(">$2-CAFE -GO FAST-\n $2-FOOD")
			text.Draw(screen, jugadores, arcadeFont, 250, 250, color.White)
		case bakery && elecCompras == 1 && p.Coins >= 2:
			jugadores := fmt.Sprintf(" $2-CAFE -GO FAST-\n>$2-FOOD")
			text.Draw(screen, jugadores, arcadeFont, 250, 250, color.White)

			//EN MART
		case mart && elecCompras == 0 && p.Coins >= 2:
			jugadores := fmt.Sprintf(">$2-face mask-GET INMUNE\n $5-HAT\n $2-CLOTHES")
			text.Draw(screen, jugadores, arcadeFont, 130, 250, color.White)
		case mart && elecCompras == 1 && p.Coins >= 2:
			jugadores := fmt.Sprintf(" $2-face mask-GET INMUNE\n>$5-HAT\n $2-CLOTHES")
			text.Draw(screen, jugadores, arcadeFont, 130, 250, color.White)
		case mart && elecCompras == 2 && p.Coins >= 2:
			jugadores := fmt.Sprintf(" $2-face mask-GET INMUNE\n $5-HAT\n>$2-CLOTHES")
			text.Draw(screen, jugadores, arcadeFont, 130, 250, color.White)

			//en SUPERMARKET
		case supermarket && elecCompras == 0 && p.Coins >= 2:
			jugadores := fmt.Sprintf(">$3-PASTA-GET LIFE-\n $2-ENERgy drink -GO FAST\n $2-TOILET PAPER")
			text.Draw(screen, jugadores, arcadeFont, 150, 250, color.White)
		case supermarket && elecCompras == 1 && p.Coins >= 2:
			jugadores := fmt.Sprintf(" $3-PASTA-GET LIFE-\n>$2-ENERgy drink -GO FAST\n $2-TOILET PAPER")
			text.Draw(screen, jugadores, arcadeFont, 150, 250, color.White)
		case supermarket && elecCompras == 2 && p.Coins >= 2:
			jugadores := fmt.Sprintf(" $3-PASTA-GET LIFE-\n $2-ENERgy drink -GO FAST\n>$2-TOILET PAPER")
			text.Draw(screen, jugadores, arcadeFont, 150, 250, color.White)
			//en SUPERMARKET
		case vacunatorio && elecCompras == 0 && p.Coins >= 2:
			jugadores := fmt.Sprintf(">$10-VACCINE-\n    -LEAVE-")
			text.Draw(screen, jugadores, arcadeFont, 200, 250, color.White)
		case vacunatorio && elecCompras == 1 && p.Coins >= 2:
			jugadores := fmt.Sprintf(" $10-VACCINE-\n>   -LEAVE-")
			text.Draw(screen, jugadores, arcadeFont, 200, 250, color.White)
		}
	}
}
