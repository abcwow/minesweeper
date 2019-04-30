package sweeper

import "fmt"
import "time"
import "math/rand"

type unit struct {
	state uint8

	besides []int
	corner  []int
}

const EMPYTUNIT byte = 10
const SAFEUNIT byte = 0xaa  //30
const SWEEPUNIT byte = 0xff //31
const SWEEPDIDUNIT byte = 0xff

func getCornerIndex(n int) []int {
	var dat [8]int

	if n == 0 {
		dat[0] = n + 31
		return dat[:1]
	} else if n == 29 {
		dat[0] = n + 29
		return dat[:1]
	} else if n == 450 {
		dat[0] = n - 29
		return dat[:1]
	} else if n == 479 {
		dat[0] = n - 31
		return dat[:1]
	} else if n < 30 {
		dat[0] = n + 29
		dat[1] = n + 31
		return dat[:2]
	} else if n > 450 {
		dat[0] = n - 29
		dat[1] = n - 31
		return dat[:2]
	} else if n%30 == 0 {
		dat[0] = n - 29
		dat[1] = n + 31
		return dat[:2]
	} else if n%30 == 29 {
		dat[0] = n - 31
		dat[1] = n + 29
		return dat[:2]
	} else {
		dat[0] = n - 31
		dat[1] = n - 29
		dat[2] = n + 29
		dat[3] = n + 31
		return dat[:4]
	}
}

func getBesideIndex(n int) []int {
	var dat [8]int

	if n == 0 {
		dat[0] = n + 1
		dat[1] = n + 30
		return dat[:2]
	} else if n == 29 {
		dat[0] = n - 1
		dat[1] = n + 30
		return dat[:2]
	} else if n == 450 {
		dat[0] = n + 1
		dat[1] = n - 30
		return dat[:2]
	} else if n == 479 {
		dat[0] = n - 1
		dat[1] = n - 30
		return dat[:2]
	} else if n < 30 {
		dat[0] = n - 1
		dat[1] = n + 1
		dat[2] = n + 30
		return dat[:3]
	} else if n > 450 {
		dat[0] = n - 1
		dat[1] = n + 1
		dat[2] = n - 30
		return dat[:3]
	} else if n%30 == 0 {
		dat[0] = n - 30
		dat[1] = n + 1
		dat[2] = n + 30
		return dat[:3]
	} else if n%30 == 29 {
		dat[0] = n - 30
		dat[1] = n - 1
		dat[2] = n + 30
		return dat[:3]
	} else {
		dat[0] = n - 30
		dat[1] = n - 1
		dat[2] = n + 1
		dat[3] = n + 30
		return dat[:4]
	}
}

func getBesideIndexAll(n int) []int {
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
	} else {
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
}

func getState(dat []byte, index []int) []byte {
	var state []byte

	for i := 0; i < len(index); i++ {
		if dat[index[i]] != EMPYTUNIT {
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
		if dat[i] == SWEEPUNIT || dat[i] == SWEEPDIDUNIT {
			n++
		}
	}
	return n
}

func getStaticArea(a []int, comm []int) []int {
	var aStatic []int
	NumExist := make(map[int]int)
	for _, v := range comm {
		NumExist[v] = v
	}
	for _, v := range a {
		_, exist := NumExist[v]
		if !exist {
			aStatic = append(aStatic, v)
		}
	}
	return aStatic
}

func getCommArea(a []int, b []int) []int {
	var comm []int
	NumExist := make(map[int]int)
	for _, v := range a {
		NumExist[v] = v
	}
	for _, v := range b {
		_, exist := NumExist[v]
		if exist {
			comm = append(comm, v)
		}
	}
	return comm
}

func getArea(a []int, b []int) (aStatic []int, bStatic []int, comm []int) {
	comm = getCommArea(a, b)
	aStatic = getStaticArea(a, comm)
	bStatic = getStaticArea(b, comm)
	return aStatic, bStatic, comm
}

func CalBorderState(dat []byte, a int, b int) {
	var tmp byte
	n1 := dat[a]
	n2 := dat[b]
	if (n2 == 0) || (n2 > 9) {
		return
	}
	if n1 < n2 {
		return
	}
	aAround := getBesideIndexAll(a)
	bAround := getBesideIndexAll(b)
	aStaticIndex, bStaticIndex, commIndex := getArea(aAround, bAround)
	commIndex = commIndex
	aStaticState := getState(dat, aStaticIndex)
	aBombNum := getSweeperNum(aStaticState)
	bStaticState := getState(dat, bStaticIndex)
	bBombNum := getSweeperNum(bStaticState)
	//commState := getState(dat, commIndex)
	//commBombNum := getSweeperNum(commState)

	aEmptyIndex := getEmptyIndex(dat, aStaticIndex)
	aEmptyNum := (byte)(len(aEmptyIndex))
	bEmptyIndex := getEmptyIndex(dat, bStaticIndex)
	bEmptyNum := (byte)(len(bEmptyIndex))

	if aBombNum > bBombNum {
		tmp = aBombNum - bBombNum
	} else {
		tmp = bBombNum - aBombNum
	}

	if n1-n2 != tmp {
		if (aBombNum > bBombNum) && bEmptyNum == tmp {
			//fmt.Println(bStaticIndex)
			for _, v := range bEmptyIndex {
				dat[v] = SWEEPUNIT
			}
		}
		if (aBombNum < bBombNum) && aEmptyNum == tmp {
			//fmt.Println(aStaticIndex)
			for _, v := range aEmptyIndex {
				dat[v] = SWEEPUNIT

			}
		}
		if aBombNum == bBombNum {

		}
	}
}

func getRandSweep() int {
	return rand.Intn(480)
}
func stateAllNotZero(dat []byte) bool {
	for i := 0; i < len(dat); i++ {
		if dat[i] == 0 {
			return false
		}
	}
	return true
}

func GetSweeper(dat []byte) []byte {
	var update bool = true
	var cnt int = 0

	for update == true {
		update = false
		for i := 0; i < 480; i++ {
			var around []int
			if dat[i] == EMPYTUNIT || dat[i] == SWEEPDIDUNIT {
				continue
			}
			//besides := getBesideIndex(i)

			beside := getBesideIndex(i)
			corner := getCornerIndex(i)

			around = beside
			for _, v := range corner {
				around = append(around, v)
			}

			emptySt := getEmptyIndex(dat, around)

			aroundState := getState(dat, around)
			n := getSweeperNum(aroundState)

			if dat[i] == (byte)(len(emptySt))+n {
				for j := 0; j < len(emptySt); j++ {
					dat[emptySt[j]] = SWEEPUNIT
					update = true
					cnt++
				}
			}
			if dat[i] == n && dat[i] != 0 && dat[i] != SWEEPUNIT {
				for j := 0; j < len(emptySt); j++ {
					dat[emptySt[j]] = SAFEUNIT
					update = true
					cnt++
				}
			}

			for _, v := range beside {
				CalBorderState(dat, i, v)
			}
		}
	}
	if cnt == 0 {
		index := getRandSweep()
		if dat[index] != SWEEPUNIT && dat[index] == 0 {
			dat[index] = SAFEUNIT
		}
	}
	return dat
}

type SweeperMap struct {
	unit [480]unit
}

func SweeperCreateMap() SweeperMap {
	rand.Seed(time.Now().Unix())
	sm := SweeperMap{}
	for i := 0; i < 480; i++ {
		sm.unit[i].besides = getBesideIndex(i)
	}
	fmt.Println(sm)
	return sm
}
