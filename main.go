// Copyright 2016 The go-vgo Project Developers. See the COPYRIGHT
// file at the top-level directory of this distribution and at
// https://github.com/go-vgo/robotgo/blob/master/LICENSE
//
// Licensed under the Apache License, Version 2.0 <LICENSE-APACHE or
// http://www.apache.org/licenses/LICENSE-2.0> or the MIT license
// <LICENSE-MIT or http://opensource.org/licenses/MIT>, at your
// option. This file may not be copied, modified, or distributed
// except according to those terms.

package main

import (
	"fmt"
	"time"

	"./Sweeper"
	"./mineUI"
)

//"github.com/go-vgo/robotgo"
func Succeed(dat []byte) bool {

	fmt.Println("!!!!!!check end!!!!!!: ", dat)
	n := len(dat)

	for i := 0; i < n; i++ {
		if dat[i] == 0 {
			return false
		}
	}

	return true
}

func aRobotgo() {

	g := NewGrid()

	sw := sweeper.SweeperCreateMap()

	g.StartNewGame()

	for {

		//g.UpdateGridState()
		fmt.Println("begin calc")
		state, err := g.UpdateGridState()
		if err != nil {
			fmt.Println("----------------------------------------------------------------------")
			fmt.Printf("%s, press any key to start another game:\n", err.Error())
			var str string
			fmt.Scanln(&str)
			g.StartNewGame()
			time.Sleep(1e9)
			continue
		}
		//

		control := sweeper.SweeperCal(sw, state[:480])
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		fmt.Println("len of control is: ", len(control))
		for i := 0; i < 16; i++ {
			szText := fmt.Sprintf("row%2d: ", i)
			for j := 0; j < 30; j++ {
				szText += fmt.Sprintf("%2d", control[i*30+j])
			}
			fmt.Println(szText)
		}

		//mouse
		ox, oy := g.OrgPos()
		// grid := *[30][16]byte
		mineUI.MoveAndClick(ox, oy, control)

		//
		if Succeed(control) {

			fmt.Println("----------------------------------------------------------------------")
			fmt.Printf("bingo, press any key to start another game:\n")
			var str string
			fmt.Scanln(&str)
			g.StartNewGame()
			time.Sleep(1e9)
			continue
		}

	}

}

func main() {
	aRobotgo()
}
