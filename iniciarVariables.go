package main

var nivel int

func iniciarVariables() {
	//Relato = true
	count = 0
	nivel = 1
	Level = 1
	//cosas de Game
	ModeTitle = true
	ElectNumPlayers = 0
	ElectPlayer = 0
	Game1.numPlayers = 1
	hScaleW = .53
	hScaleH = .45
	hgt = float64(player1.FrameHeight[0]) * hScaleH
	wth = float64(player1.FrameWidth[0]) * hScaleW
	player1.señaladorBool = false
	player2.señaladorBool = false
	// plyrScale = .65

}
func pasarNivel() {
	Level++
	// hack = false
	pasarNivelPlayer()
	ModeTitleLevel = true
	nivel = Level * 2
	if nivel > 20 {
		nivel = 20
	}
	mart1 = false
	bakery1 = false
	supermarket1 = false
	farmacia1 = false

	initObjetos()
	// initSonido()
	count = 0
	initEnemigos()
	//reinciar enemigos
	count1 = 0
	//nube
	initNube()

}
