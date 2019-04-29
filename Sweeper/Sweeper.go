package sweeper

import "fmt"

type unit struct {
	state   uint8
	besides []int
}

func getBesideIndex(n int) []int {
	var dat [8]int

	if n == 0 {
		dat[0] = n + 1
		dat[1] = n + 30
		dat[2] = n + 31
		return dat[:3]
	} else if n == 29 {
		dat[0] = n - 1
		dat[1] = n + 29
		dat[2] = n + 30
		return dat[:3]
	} else if n == 450 {
		dat[0] = n + 1
		dat[1] = n - 30
		dat[2] = n - 29
		return dat[:3]
	} else if n == 479 {
		dat[0] = n - 1
		dat[1] = n - 30
		dat[2] = n - 31
		return dat[:3]
	} else if n < 30 {
		dat[0] = n - 1
		dat[1] = n + 1
		dat[2] = n + 29
		dat[3] = n + 30
		dat[4] = n + 31
		return dat[:5]
	} else if n > 450 {
		dat[0] = n - 1
		dat[1] = n + 1
		dat[2] = n - 29
		dat[3] = n - 30
		dat[4] = n - 31
		return dat[:5]
	} else if n%30 == 0 {
		dat[0] = n - 30
		dat[1] = n - 29
		dat[2] = n + 1
		dat[3] = n + 30
		dat[4] = n + 31
		return dat[:5]
	} else if n%30 == 29 {
		dat[0] = n - 31
		dat[1] = n - 30
		dat[2] = n - 1
		dat[3] = n + 29
		dat[4] = n + 30
		return dat[:5]
	}
	dat[0] = n - 31
	dat[1] = n - 30
	dat[2] = n - 29
	dat[3] = n - 1
	dat[4] = n + 1
	dat[5] = n + 29
	dat[6] = n + 30
	dat[7] = n + 31
	return dat[:8]
}

func getBesideState(dat []byte, index []int) []byte {
	var state []byte

	for i := 0; i < len(index); i++ {
		if dat[index[i]] != 10 {
			state = append(state, (byte)(dat[index[i]]))
		}

	}
	return state
}

func getEmptyIndex(dat []byte, index []int) []int {
	var state []int

	for i := 0; i < len(index); i++ {
		if dat[index[i]] == 0 {
			state = append(state, index[i])
		}
	}
	return state
}

func getSweeperNum(dat []byte) byte {
	var n byte = 0
	for i := 0; i < len(dat); i++ {
		if dat[i] == 0xff {
			n++
		}
	}
	return n
}

func GetSweeper(dat []byte) []byte {

	var update bool = true

	for update == true {
		update = false
		for i := 0; i < 480; i++ {
			if dat[i] == 10 {
				continue
			}
			besides := getBesideIndex(i)
			emptySt := getEmptyIndex(dat, besides)

			st := getBesideState(dat, besides)
			n := getSweeperNum(st)

			if dat[i] == (byte)(len(emptySt))+n {
				for j := 0; j < len(emptySt); j++ {
					dat[emptySt[j]] = 0xff
					update = true
				}
			}
			if dat[i] == n && dat[i] != 0 && dat[i] != 0xff {
				for j := 0; j < len(emptySt); j++ {
					dat[emptySt[j]] = 0xaa
					update = true
				}
			}
		}
	}

	return dat
}

type SweeperMap struct {
	unit [480]unit
}

func SweeperCreateMap() SweeperMap {
	sm := SweeperMap{}
	for i := 0; i < 480; i++ {
		sm.unit[i].besides = getBesideIndex(i)
	}
	fmt.Println(sm)
	return sm
}
