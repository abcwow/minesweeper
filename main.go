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

//"github.com/go-vgo/robotgo"

func aRobotgo() {

	g := NewGrid()

	for {
		g.UpdateGridState()
		//state := g.UpdateGridState()

		//

		//mouse
		//ox, oy := g.OrgPos()
		// grid := *[30][16]byte
		//MoveAndClick(ox, oy, grid)

		//time.Sleep(3e9)

	}

}

func main() {
	aRobotgo()
}
