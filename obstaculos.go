package main

func obstaculos(X float64, Y float64, X1 float64, Y1 float64) (float64, float64, bool) {
	var objetos [][]int
	var bul bool

	switch {
	case casita:
		objetos = make([][]int, 21)
		objetos[0] = []int{645, 20, 198, 30}
		objetos[1] = []int{634, 50, 170, 37}
		objetos[2] = []int{518, 78, 185, 25}
		objetos[3] = []int{389, 44, 179, 36}
		objetos[4] = []int{382, 81, 235, 31}
		objetos[5] = []int{548, 48, 244, 53}
		objetos[6] = []int{439, 54, 321, 60}
		objetos[7] = []int{415, 20, 330, 20}
		objetos[8] = []int{389, 17, 356, 41}
		objetos[9] = []int{666, 18, 354, 44}
		objetos[10] = []int{691, 368, 1, 529}
		objetos[11] = []int{576, 482, 410, 120}
		objetos[12] = []int{-1, 511, 411, 119}
		objetos[13] = []int{-2, 383, 0, 527}
		objetos[14] = []int{-2, 1061, -1, 184}
		objetos[15] = []int{491, 120, 447, 83}
		objetos[16] = []int{553, 50, 329, 22}
		//abajo para que no salga de la pantalla
		objetos[17] = []int{0, 0, 0, screenHeight}
		objetos[18] = []int{screenWidth, 0, 0, screenHeight}
		objetos[19] = []int{0, screenWidth, screenHeight, 32}
		objetos[20] = []int{0, screenWidth, 30, 0}

	case banco:
		objetos = make([][]int, 35)
		objetos[0] = []int{455, 128, 212, 55}
		objetos[1] = []int{615, 66, 262, 4}
		objetos[2] = []int{555, 94, 212, 22}
		objetos[3] = []int{678, 2, 173, 92}
		objetos[4] = []int{616, 63, 159, 26}
		objetos[5] = []int{454, 115, 160, 59}
		objetos[6] = []int{563, 56, 291, 28}
		objetos[7] = []int{423, 2, 265, 31}
		objetos[8] = []int{359, 65, 264, 3}
		objetos[9] = []int{359, 1, 183, 82}
		objetos[10] = []int{358, 49, 184, 1}
		objetos[11] = []int{406, 3, 158, 26}
		objetos[12] = []int{391, 65, 211, 25}
		objetos[13] = []int{420, 54, 293, 23}
		objetos[14] = []int{614, 2, 265, 30}
		objetos[15] = []int{542, 146, 157, 1}
		//limite
		objetos[16] = []int{0, 1057, -1, 159}
		objetos[17] = []int{704, 353, 157, 398}
		objetos[18] = []int{0, 350, 157, 394}
		objetos[19] = []int{349, 117, 336, 215}
		objetos[20] = []int{464, 182, 401, 150}
		objetos[21] = []int{560, 225, 351, 203}
		objetos[22] = []int{575, 143, 335, 18}
		objetos[23] = []int{463, 17, 352, 61}
		objetos[24] = []int{567, 8, 344, 8}
		objetos[25] = []int{463, 8, 344, 10}
		objetos[26] = []int{673, 30, 164, 10}
		objetos[27] = []int{355, 25, 322, 13}
		objetos[28] = []int{690, 10, 310, 26}
		objetos[29] = []int{673, 28, 324, 9}
		objetos[30] = []int{353, 14, 308, 18}

		objetos[31] = []int{0, 0, 0, screenHeight}
		objetos[32] = []int{screenWidth, 0, 0, screenHeight}
		objetos[33] = []int{0, screenWidth, screenHeight, 32}
		objetos[34] = []int{0, screenWidth, 30, 0}

	default:
		objetos = make([][]int, 76)
		objetos[0] = []int{232, 85, 123, 104}
		objetos[1] = []int{355, 247, 142, 76}
		objetos[2] = []int{297, 63, -1, 72}
		objetos[3] = []int{-1, 275, 0, 96}
		objetos[4] = []int{-3, 64, 0, 250}
		objetos[5] = []int{111, 15, 134, 114}
		objetos[6] = []int{111, 35, 235, 11}
		objetos[7] = []int{144, 17, 216, 40}
		objetos[8] = []int{134, 12, 141, 40}
		objetos[9] = []int{128, 25, 190, 38}
		objetos[10] = []int{62, 31, 128, 119}
		objetos[11] = []int{19, 38, 329, 83}
		objetos[12] = []int{25, 29, 448, 80}
		objetos[13] = []int{102, 73, 329, 48}
		objetos[14] = []int{121, 27, 385, 13}
		objetos[15] = []int{125, 22, 440, 12}
		objetos[16] = []int{153, 22, 444, 45}
		objetos[17] = []int{160, 26, 406, 30}
		objetos[18] = []int{100, 41, 460, 67}
		objetos[19] = []int{240, 12, 353, 64}
		objetos[20] = []int{254, 8, 391, 24}
		objetos[21] = []int{232, 8, 304, 27}
		objetos[22] = []int{345, 54, 278, 160}
		objetos[23] = []int{336, 68, 285, 64}
		objetos[24] = []int{364, 34, 448, 59}
		objetos[25] = []int{365, 39, 518, 10}
		objetos[26] = []int{443, 180, 315, 68}
		objetos[27] = []int{517, 105, 381, 147}
		objetos[28] = []int{452, 48, 439, 88}
		objetos[29] = []int{496, 22, 385, 143}
		objetos[30] = []int{523, 101, -2, 203}
		objetos[31] = []int{494, 10, 57, 11}
		objetos[32] = []int{504, 15, 77, 20}
		objetos[33] = []int{373, 16, 73, 25}
		objetos[34] = []int{521, 100, 292, 29}
		objetos[35] = []int{540, 24, 284, 16}
		objetos[36] = []int{592, 24, 284, 19}
		objetos[37] = []int{463, 40, 298, 27}
		objetos[38] = []int{453, 60, 303, 13}
		objetos[39] = []int{487, 39, 119, 26}
		objetos[40] = []int{372, 118, 132, 10}
		objetos[41] = []int{698, 17, 280, 108}
		objetos[42] = []int{688, 52, 291, 237}
		objetos[43] = []int{679, 12, 345, 45}
		objetos[44] = []int{749, 5, 505, 19}
		objetos[45] = []int{734, 11, 485, 44}
		objetos[46] = []int{812, 40, 292, 92}
		objetos[47] = []int{874, 108, 292, 235}
		objetos[48] = []int{809, 29, 435, 91}
		objetos[49] = []int{857, 17, 437, 92}
		objetos[50] = []int{837, 27, 460, 68}
		objetos[51] = []int{853, 21, 301, 101}
		objetos[52] = []int{862, 15, 400, 41}
		objetos[53] = []int{1010, 48, 331, 195}
		objetos[54] = []int{1022, 35, 311, 20}
		objetos[55] = []int{1030, 27, 304, 7}
		objetos[56] = []int{720, 338, -2, 213}
		objetos[57] = []int{703, 22, -1, 190}
		objetos[58] = []int{691, 38, -2, 137}
		objetos[59] = []int{-3, 646, -1, 29}
		objetos[60] = []int{642, 26, -1, 17}
		objetos[61] = []int{621, 34, -2, 26}
		objetos[62] = []int{659, 21, -2, 7}
		objetos[63] = []int{550, 44, 204, 22}
		objetos[64] = []int{315, 26, 52, 25}
		objetos[65] = []int{147, 34, 510, 16}
		objetos[66] = []int{311, 22, 442, 47}
		objetos[67] = []int{308, 24, 505, 21}
		objetos[68] = []int{52, 94, 80, 31}
		objetos[69] = []int{354, 169, 22, 33}
		objetos[70] = []int{723, 0, 60, 0}
		objetos[71] = []int{796, 74, 307, 76}
		//abajo para que no salga de la pantalla
		objetos[72] = []int{0, 0, 0, screenHeight}
		objetos[73] = []int{screenWidth, 0, 0, screenHeight}
		objetos[74] = []int{0, screenWidth, screenHeight, 32}
		objetos[75] = []int{0, screenWidth, 30, 0}

	}

	for i := 0; i < len(objetos); i++ {
		if int(X+(wth-7)) > objetos[i][0] && int(X+7) < objetos[i][0]+objetos[i][1] && int(Y+hgt) > objetos[i][2] && int(Y+(hgt-5)) < objetos[i][2]+objetos[i][3] {
			X = X1
			Y = Y1
			bul = true
		}
	}

	// for j := randNum; j < numEnemigo+randNum; j++ {
	// 	if i != j && X+(wth-7) > enemigo.X[j]+5 && X+5 < enemigo.X[j]+(wth-7) && Y+hgt > enemigo.Y[j]+(hgt-10) && Y+(hgt-10) <= enemigo.Y[j]+hgt {
	// 		//X+wth-7 > enemigo.X[j] && X+7 < enemigo.X[j]+wth-7 && Y+hgt > enemigo.Y[j]-15 && Y < enemigo.Y[j]+hgt {
	// 		// X = X1
	// 		// Y = Y1
	// 		bul = true
	// 	}

	// }

	return X, Y, bul
}
