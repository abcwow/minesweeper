package main

import (
	"github.com/go-vgo/robotgo"
)

func MoveAndClick(startX, startY int, grid []byte) {

	for i := 0; i < 16; i++ {
		for j := 0; j < 30; j++ {
			if grid[i*30+j] == 31 {

				ex := startX + j*16 + 8
				ey := startY + i*16 + 8
				robotgo.MoveMouseSmooth(ex, ey, 1.0, 3.0)
				robotgo.MouseClick("right", true)

			}
			if grid[i*30+j] == 30 {

				ex := startX + j*16 + 8
				ey := startY + i*16 + 8
				robotgo.MoveMouseSmooth(ex, ey, 1.0, 3.0)
				robotgo.MouseClick("left", true)
			}
		}
	}
}
