package main

import (
	"github.com/go-vgo/robotgo"
)

//GetGrid30X16Y 把16*30 的数组转换为坐标位置，并读取鼠标左右键点击选择；0xff左键 0xaa右键
func GetGrid30X16(grid *[30][16]byte) (int, int, bool) {
	flag := true
	for i := 0; i < 30; i++ {
		for j := 0; j < 16; j++ {
			if grid[i][j] == 0xff {
				flag = true
				return i, j, flag
			}
			if grid[i][j] == 0xaa {
				flag = false
				return i, j, flag
			}
		}
	}
	return -2, -2, false
}

//MouseMove 移动到输入的起始坐标位置 fx fy 终点x y坐标； 产生鼠标移动并点击当前的格子事件
func MoveAndClick(startX, startY int, grid *[30][16]byte) {

	y, x, flag := GetGrid30X16(grid)

	if x < 0 || y < 0 {
		return
	}

	ex := startX + x*16
	ey := startY + y*16

	robotgo.MoveMouseSmooth(ex, ey, 1.0, 3.0) //模拟人移鼠标轨迹
	if flag == true {
		robotgo.MouseClick("left", true) // true 鼠标双击
	} else if flag == false {
		robotgo.MouseClick("right", true)
	}
}
