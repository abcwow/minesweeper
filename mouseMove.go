package main

import (
	"github.com/go-vgo/robotgo"
)

func MoveAndClick(startX, startY int, grid *[30][16]byte) {

	for i := 0; i < 30; i++ {
		for j := 0; j < 16; j++ {
			if grid[i][j] == 31 {

				ex := startX + j*16
				ey := startY + i*16
				robotgo.MoveMouseSmooth(ex, ey, 1.0, 3.0)
				robotgo.MouseClick("right", true)

			}
			if grid[i][j] == 30 {

				ex := startX + j*16
				ey := startY + i*16
				robotgo.MoveMouseSmooth(ex, ey, 1.0, 3.0)
				robotgo.MouseClick("left", true)
			}
		}
	}
}
