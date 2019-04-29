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

	"./Sweeper"
)

//"github.com/go-vgo/robotgo"

func aRobotgo() {

	g := NewGrid()

	for {
		//g.UpdateGridState()
		state, err := g.UpdateGridState()
		if err != nil {
			fmt.Println("----------------------------------------------------------------------")
			fmt.Printf("%s, press any key to start another game:\n", err.Error())
			var str string
			fmt.Scanln(&str)
			g.StartNewGame()
		}
		//
		control := sweeper.GetSweeper(state[:480])
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
		MoveAndClick(ox, oy, control)

		//time.Sleep(3e9)

	}

}

func main() {
	aRobotgo()
}
