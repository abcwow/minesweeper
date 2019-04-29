package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

const (
	GRID_H int = 16
	GRID_W int = 30
	CELL_N int = GRID_H * GRID_W

	CELL_W int = 16

	CHARACTER_N       int = 16
	CHARACTER_ELEMENT int = 3
)

const (
	CELL_BLANK_PUSH = iota
	CELL_1
	CELL_2
	CELL_3
	CELL_4
	CELL_5
	CELL_6
	CELL_7
	CELL_8
	CELL_QUESTION_PUSH
	CELL_BOMB_PUSH
	CELL_BOMB_CROSS
	CELL_BOMB_RED
	CELL_QUESTION
	CELL_FLAG
	CELL_BLANK
)

type Grid struct {
	ox int
	oy int

	resbmp robotgo.CBitmap

	screenbmp robotgo.CBitmap
}

func NewGrid() *Grid {

	var g Grid

	g.GetOrgPos()

	return &g

}

func (g *Grid) GetOrgPos() {

	bmp := robotgo.OpenBitmap("res.png")

	bits := robotgo.GetPortion(bmp, 0, 0, 16, 16)

	fx, fy := robotgo.FindBitmap(bits)
	g.ox = fx
	g.oy = fy

	fmt.Printf("ox = %d, oy = %d\n", g.ox, g.oy)

	g.resbmp = robotgo.CBitmap(bmp)

}

func (g *Grid) CheckMatchCharacter(row, col, index int) bool {

	fx := CELL_W * col
	fy := CELL_W * row

	x, y := fx+8, fy
	resx, resy := 8, index*CELL_W
	rescolor := robotgo.GetColor(robotgo.ToMMBitmapRef(g.resbmp), resx, resy)
	color := robotgo.GetColor(robotgo.ToMMBitmapRef(g.screenbmp), x, y)
	if row == 0 && col <= 1 {
		fmt.Printf("row 0 col %d color 1(scr %d, %d, bmp %d, %d): bmp %x scr %x\n", col, x, y, resx, resy, rescolor, color)
	}
	if rescolor != color {
		return false
	}

	x, y = fx+8, fy+7
	resx, resy = 8, index*CELL_W+7
	rescolor = robotgo.GetColor(robotgo.ToMMBitmapRef(g.resbmp), resx, resy)
	color = robotgo.GetColor(robotgo.ToMMBitmapRef(g.screenbmp), x, y)
	if row == 0 && col <= 1 {
		fmt.Printf("row 0 col %d color 2(scr %d, %d, bmp %d, %d): bmp %x scr %x\n", col, x, y, resx, resy, rescolor, color)
	}
	if rescolor != color {
		return false
	}

	x, y = fx+8, fy+15
	resx, resy = 8, index*CELL_W+15
	rescolor = robotgo.GetColor(robotgo.ToMMBitmapRef(g.resbmp), resx, resy)
	color = robotgo.GetColor(robotgo.ToMMBitmapRef(g.screenbmp), x, y)
	if row == 0 && col <= 1 {
		fmt.Printf("row 0 col %d color 3(scr %d, %d, bmp %d, %d): bmp %x scr %x\n", col, x, y, resx, resy, rescolor, color)
	}
	if rescolor != color {
		return false
	}

	return true
}

func (g *Grid) GetCellState(row, col int) byte {

	for i := 0; i < CHARACTER_N; i++ {
		if g.CheckMatchCharacter(row, col, i) {
			return byte(i)
		}
	}

	return 0

}

func (g *Grid) OrgPos() (ox, oy int) {
	return g.ox, g.oy
}

func (g *Grid) ForUseData(val byte) byte {
	if val >= 7 && val <= 14 {
		return 15 - val
	} else if val == 1 {
		return 0xff
	} else if val == 0 {
		return 0
	} else if val == 15 {
		return 0x0A
	}

	return 0
}

func (g *Grid) UpdateGridState() (info [CELL_N]byte) {

	fmt.Println("===================================================================")

	g.screenbmp = robotgo.CBitmap(robotgo.CaptureScreen(g.ox, g.oy, CELL_W*GRID_W, CELL_W*GRID_H))
	robotgo.SaveBitmap(robotgo.ToMMBitmapRef(g.screenbmp), "test.png", 1)
	for i := 0; i < GRID_H; i++ {
		szText := fmt.Sprintf("row%2d: ", i)
		for j := 0; j < GRID_W; j++ {
			info[i*GRID_W+j] = g.ForUseData(g.GetCellState(i, j))
			szText += fmt.Sprintf("%2d", info[i*GRID_W+j])
		}
		fmt.Println(szText)
	}

	return

}
