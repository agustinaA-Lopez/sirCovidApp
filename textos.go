package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

func initTextos() {

	tt, err := truetype.Parse(fonts.ArcadeN_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	arcadeFont = truetype.NewFace(tt, &truetype.Options{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	smallArcadeFont = truetype.NewFace(tt, &truetype.Options{
		Size:    smallFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func dibujarTextos(screen *ebiten.Image) {
	// dibujar vidas y monedas
	lifesP1 := fmt.Sprintf("Life:%02d  Coins:%d", player1.vidas, player1.Coins)
	text.Draw(screen, lifesP1, smallArcadeFont, 20, 515, color.Black)
	if Game1.numPlayers == 2 {
		lifesP2 := fmt.Sprintf("Life:%02d  Coins:%d", player2.vidas, player2.Coins)
		text.Draw(screen, lifesP2, smallArcadeFont, 785, 515, color.Black)
	}

	//dibujar inmunidad y fast
	if player1.Inmune {
		Inm := fmt.Sprintf("Inmune for:%02d", player1.CountPoder/60)
		text.Draw(screen, Inm, smallArcadeFont, fontSize, 25, color.Black)
	}
	if player2.Inmune {
		Inm := fmt.Sprintf("Inmune for:%02d", player2.CountPoder/60)
		text.Draw(screen, Inm, smallArcadeFont, 800, 25, color.Black)
	}
	//dibujar Fast
	if player1.Fast {
		Inm := fmt.Sprintf("Fast for:%02d", player1.CountPoder/60)
		text.Draw(screen, Inm, smallArcadeFont, 260, 25, color.Black)
	}
	if player2.Fast {
		Inm := fmt.Sprintf("Fast for:%02d", player2.CountPoder/60)
		text.Draw(screen, Inm, smallArcadeFont, 600, 25, color.Black)
	}

	come := "& come back safe"
	switch {
	case Commands:
		// ModeGameOver = false
		plasmaVida := fmt.Sprintf("Player 1 movements:\n\n Up: Arrow Up\n Dow: Arrow Down\n Rigth: Arrow Rigth\n Left: Arrow Left")
		text.Draw(screen, plasmaVida, smallArcadeFont, 325, 150, color.White)
		alcholInmune := fmt.Sprintf("\nPlayer 2 movements:\n\n Up: W\n Dow: S\n Rigth: D\n Left: A")
		text.Draw(screen, alcholInmune, smallArcadeFont, 325, 250, color.White)
		barbijoInmune := fmt.Sprintf("\nPause: Space Bar\n\nMute: X\n\nVolume Up: Ctrl+Arrow Up\nVolume Down: Ctrl+Arrow Down\n\nCredits: C")
		text.Draw(screen, barbijoInmune, smallArcadeFont, 325, 360, color.White)

		//dibujar creditos
	case Credits:
		// ModeGameOver = false
		credits := fmt.Sprintf("Credits:\n\nMusic and Sound: Kim Kaos\n\nGraphics:City by Fere Duelli\n\nHome: tileset by Kymotonian & Speedialga\n\nBank: titelest by Princess Phoenix\n\nSoftware:L.A.T. Software Factory\n\nCaracters graphics: Matthias1")
		text.Draw(screen, credits, smallArcadeFont, 275, 150, color.White)

	//switch {
	case ModeTitle:
		// intro draw
		intro1.drawIntro(screen, screenWidth, screenHeight)

	case ModeTitleLevel:
		nivel := fmt.Sprintf("    LEVEL %d", Level)
		text.Draw(screen, nivel, arcadeFont, 290, 200, color.Black)
		nivel1 := fmt.Sprintf("PRESS SPACE KEY")
		text.Draw(screen, nivel1, smallArcadeFont, 400, 290, color.Black)

		if Level == 1 {
			nivel := fmt.Sprintf("MISSION: GET MEDS\n" + come)
			text.Draw(screen, nivel, smallArcadeFont, 380, 230, color.Black)
		}
		if Level == 2 {
			nivel := fmt.Sprintf("MISSION: GET FOOD\n" + come)
			text.Draw(screen, nivel, smallArcadeFont, 380, 230, color.Black)
		}
		if Level == 3 {
			nivel := fmt.Sprintf("MISSION GET CLOTHES\n" + come)
			text.Draw(screen, nivel, smallArcadeFont, 370, 230, color.Black)
		}
		if Level == 4 {
			nivel := fmt.Sprintf("MISSION: GET TOLIET PAPER\n   " + come)
			text.Draw(screen, nivel, smallArcadeFont, 314, 230, color.Black)
		}
		if Level == 5 {
			nivel := fmt.Sprintf("MISSION: GET MEDS,\nFOOD\n" + come)
			text.Draw(screen, nivel, smallArcadeFont, 370, 230, color.Black)
		}
		if Level == 6 {
			nivel := fmt.Sprintf("MISSION: GET TOLET PAPER,\nMEDS " + come)
			text.Draw(screen, nivel, smallArcadeFont, 320, 230, color.Black)
		}
		if Level == 7 {
			nivel := fmt.Sprintf("MISSION: GET GLOTHES,\nFOOD\n" + come)
			text.Draw(screen, nivel, smallArcadeFont, 350, 230, color.Black)
		}
		if Level == 8 {
			nivel := fmt.Sprintf("MISSION: GET MEDS,\nCLOTHES " + come)
			text.Draw(screen, nivel, smallArcadeFont, 320, 225, color.Black)
		}
		if Level == 9 {
			nivel := fmt.Sprintf("MISSION: Toilet PAPER,\nCLOTHES " + come)
			text.Draw(screen, nivel, smallArcadeFont, 320, 230, color.Black)
		}
		if Level == 10 {
			nivel := fmt.Sprintf("MISSION: GET food,\nTOILET PAPER\n" + come)
			text.Draw(screen, nivel, smallArcadeFont, 320, 225, color.Black)
		}
		if Level > 10 {
			nivel := fmt.Sprintf("MISSION: GET VACCINE\n&come back free")
			text.Draw(screen, nivel, smallArcadeFont, 350, 230, color.Black)
		}

	case ModeGameOver:
		lost := fmt.Sprintf("  GAME OVER!\n\n  TRAY AGAIN?\n\nPRESS SPACE KEY\n\n c for credits\n\n ")
		noMoney := fmt.Sprintf("Mission failed!\n you ran out of money\nTRAY AGAIN?\n\nPRESS SPACE KEY")

		if player1.Coins < 2 && player2.Coins < 2 && monedas.X == 1500 {
			text.Draw(screen, noMoney, arcadeFont, 220, 200, color.White)
		} else {
			text.Draw(screen, lost, arcadeFont, 310, 200, color.White)
		}
	case ModeWin:
		win := fmt.Sprintf("you succeded surviving\nthe pandemic!\n\n       YOU WIN\n\n  press c for creidts")
		text.Draw(screen, win, arcadeFont, 200, 200, color.White)
	}
	switch {
	case Commands:
		// ModeGameOver = false
		plasmaVida := fmt.Sprintf("Player 1 movements:\n\n Up: Arrow Up\n Dow: Arrow Down\n Rigth: Arrow Rigth\n Left: Arrow Left")
		text.Draw(screen, plasmaVida, smallArcadeFont, 325, 150, color.White)
		alcholInmune := fmt.Sprintf("\nPlayer 2 movements:\n\n Up: W\n Dow: S\n Rigth: D\n Left: A")
		text.Draw(screen, alcholInmune, smallArcadeFont, 325, 250, color.White)
		barbijoInmune := fmt.Sprintf("\nPause: Space Bar\n\nMute: X\n\nVolume Up: Ctrl+Arrow Up\nVolume Down: Ctrl+Arrow Down\n\nCredits: C")
		text.Draw(screen, barbijoInmune, smallArcadeFont, 325, 360, color.White)

		//dibujar creditos
	case Credits:
		// ModeGameOver = false
		credits := fmt.Sprintf("Credits:\n\nMusic and Sound: Kim Kaos\n\nGraphics:City by Fere Duelli\n\nHome: tileset by Kymotonian & Speedialga\n\nBank: titelest by Princess Phoenix\n\nSoftware:L.A.T. Software Factory\n\nCaracters graphics: Matthias1n\n\nmade whith ebiten library")
		text.Draw(screen, credits, smallArcadeFont, 275, 150, color.White)
	case ModePause && count1 < 40:
		jugadores := fmt.Sprintf("PAUSE")
		text.Draw(screen, jugadores, arcadeFont, 440, 220, color.Black)

		//elegir numero de jugadores
	case ElectNumPlayers == 0 && Game1.numPlayers == 1:
		jugadores := fmt.Sprintf(">1 PLAYER\n 2 PLAYERS")
		text.Draw(screen, jugadores, arcadeFont, 350, 300, color.Black)
	case ElectNumPlayers == 0 && Game1.numPlayers == 2:
		jugadores := fmt.Sprintf(" 1 PLAYER\n>2 PLAYERS")
		text.Draw(screen, jugadores, arcadeFont, 350, 300, color.Black)

		//elegir jugador
	case ElectPlayer == 0:
		jugadores := fmt.Sprintf("CHOOSE PLAYER")
		text.Draw(screen, jugadores, arcadeFont, 320, 300, color.Black)

	case ElectNumPlayers == 1:
		for i, l := range texts {
			x := (screenWidth - len(l)*fontSize) / 2
			text.Draw(screen, l, arcadeFont, x, (i+5)*fontSize, color.White)
		}
	}
	//dibujar Comandos

}
