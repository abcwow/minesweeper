package mineUI

import (
	"github.com/go-vgo/robotgo"
)

type mineGrid struct {
	width     int
	heigh     int
	disetnece float64
	click     int //鼠标点击类型 0：不做操作；1：左键单击；2：右键单击;
}

const CLICK1 = "right"
const CLICK2 = "left"
const SAFENUM = 31
const MINENUM = 30
const DoubleClick = true
const SingleClick = false

func (p *mineGrid) getDistence(x, y int) {
	p.disetnece = float64(x*x + y*y) //只做距离判断，不计算实际距离
}

var nextWidth, nextHeigh int
var GameStart = 0

//输入起点坐标像素，w h；待点击的数组grid；生成对应坐标点的绝对位置地图
func AbsoluteMap(w, h int, grid []byte) []mineGrid {
	var path mineGrid
	mineMap := make([]mineGrid, 0, 480)

	for i := 0; i < 16; i++ {
		for j := 0; j < 30; j++ {
			if grid[i*30+j] == SAFENUM {
				path.width = w + j*16 + 8
				path.heigh = h + i*16 + 8
				//path.getDistence(path.width, path.heigh)
				path.click = 1
				mineMap = append(mineMap, path)

			}
			if grid[i*30+j] == MINENUM {
				path.width = w + j*16 + 8
				path.heigh = h + i*16 + 8
				//path.getDistence(path.width, path.heigh)
				path.click = 2
				mineMap = append(mineMap, path)
			}
		}
	}
	return mineMap
}

//搜索最小值--快速排序
//left 需要排序的数组起始位置， right 最右边的那个数的位置（下标），a 排序的数组
func quicksort(left int, right int, a []mineGrid) {
	if left >= right { //需要排序的起始位置大于或等于终止位置，就表明不再需要排序
		return
	}
	i := left       //左边的游标
	j := right      //右边的游标
	temp := a[left] // 基数
	for i != j {
		for a[j].disetnece >= temp.disetnece && i < j { //直到遇到了 <temp 的数就停下来
			j--
		}
		for a[i].disetnece <= temp.disetnece && i < j { //直到遇到了 >temp 的数就停下来
			i++
		}
		//交换这两个数
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}
	//将基数归位
	a[i], a[left] = temp, a[i]

	//递归处理基数左边未处理的
	quicksort(left, i-1, a)

	//递归处理基数右边未处理的
	if i != len(a)-1 {
		quicksort(i+1, right, a)
	}
}

//计算点对其他点的距离
func mapDistence(width, heigh int, a []mineGrid) {
	for i := 0; i < len(a); i++ {
		x := (a[i].width - width) * (a[i].width - width)
		y := (a[i].heigh - heigh) * (a[i].heigh - heigh)
		a[i].getDistence(x, y)
	}
}

//鼠标移动排序后距离
func mouseMove(a []mineGrid) {
	quicksort(0, len(a)-1, a)
	robotgo.MoveMouseSmooth(a[0].width, a[0].heigh, 1.0, 1.0)
}

func mouseClik(a mineGrid) (hasClickW, hasClickH int) {
	switch a.click {
	case 1:
		robotgo.MouseClick(CLICK1, SingleClick)
	case 2:
		robotgo.MouseClick(CLICK2, SingleClick)
	}
	return a.width, a.heigh
}

func MoveAndClick(w, h int, grid []byte) {
	mineMap := AbsoluteMap(w, h, grid)
	if GameStart == 1 {
		mapDistence(nextWidth, nextHeigh, mineMap)
	} else {
		mapDistence(w, h, mineMap)
	}
	GameStart = 1
	mouseMove(mineMap)
	for len(mineMap) > 1 {
		hasClickW, hasClickH := mouseClik(mineMap[0])
		mineMap = mineMap[1:]
		mapDistence(hasClickW, hasClickH, mineMap)
		mouseMove(mineMap)
	}
	nextWidth, nextHeigh = mouseClik(mineMap[0])
}
