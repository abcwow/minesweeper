package sweeper

import "fmt"
import "time"
import "math/rand"

const UNKNOWUNIT byte = 0
const EMPYTUNIT byte = 10
const SAFEUNIT byte = 30  //30
const SWEEPUNIT byte = 31 //31
const SWEEPDIDUNIT byte = 32

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

func getAroundIndex(n int) []int {
	var around []int
	beside := getBesideIndex(n)
	corner := getCornerIndex(n)

	around = beside
	for _, v := range corner {
		around = append(around, v)
	}
	return around
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

func getBombNum(dat []byte) byte {
	var n byte = 0
	for i := 0; i < len(dat); i++ {
		if dat[i] == SWEEPUNIT || dat[i] == SWEEPDIDUNIT {
			n++
		}
	}
	return n
}

func getStaticArea(a []int, comm []int, n int) []int {
	var aStatic []int
	NumExist := make(map[int]int)
	for _, v := range comm {
		NumExist[v] = v
	}
	for _, v := range a {
		_, exist := NumExist[v]
		if !exist && v != n {
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

func getArea(a []int, b []int, aIndex int, bIndex int) (aStatic []int, bStatic []int, comm []int) {
	comm = getCommArea(a, b)
	aStatic = getStaticArea(a, comm, bIndex)
	bStatic = getStaticArea(b, comm, aIndex)
	return aStatic, bStatic, comm
}

func getAbs(n int) int {
	if n > 0 {
		return n
	} else {
		return 0 - n
	}
}

func sweeperCalMulUnit(dat []byte, a int, b int) {
	if dat[a] < dat[b] {
		return
	}
	aAround := getAroundIndex(a)
	bAround := getAroundIndex(b)
	aStaticIndex, bStaticIndex, commIndex := getArea(aAround, bAround, a, b)
	aStaticState := getState(dat, aStaticIndex)
	bStaticState := getState(dat, bStaticIndex)
	commState := getState(dat, commIndex)
	aStaticBombNum := getBombNum(aStaticState)
	bStaticBomnNum := getBombNum(bStaticState)
	commBombNum := getBombNum(commState)
	commBombNum = commBombNum

	aEmptyIndex := getEmptyIndex(dat, aStaticIndex)
	aEmptyNum := (byte)(len(aEmptyIndex))
	bEmptyIndex := getEmptyIndex(dat, bStaticIndex)
	bEmptyNum := (byte)(len(bEmptyIndex))

	if dat[a] == dat[b] {
		if aStaticBombNum != bStaticBomnNum {
			if aStaticBombNum < bStaticBomnNum && (byte)(aEmptyNum) == bStaticBomnNum-aStaticBombNum {
				for _, v := range aEmptyIndex {
					dat[v] = SWEEPUNIT
				}
			}
			if aStaticBombNum > bStaticBomnNum && (byte)(bEmptyNum) == aStaticBombNum-bStaticBomnNum {
				for _, v := range bEmptyIndex {
					dat[v] = SWEEPUNIT
				}
			}
		}
	} else {

	}
}

func getRandSweep() int {
	//fix
	return rand.Intn(480)
}

func GetSweeper(dat []byte) []byte {
	var update bool = true
	var cnt int = 0

	for update == true {
		update = false
		for i := 0; i < 480; i++ {
			if dat[i] == EMPYTUNIT || dat[i] == SWEEPDIDUNIT {
				continue
			}
			beside := getBesideIndex(i)
			around := getAroundIndex(i)

			emptySt := getEmptyIndex(dat, around)
			aroundState := getState(dat, around)
			n := getBombNum(aroundState)

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
				v = v
				if dat[v] != UNKNOWUNIT && dat[v] != 0x10 && dat[v] != SWEEPUNIT {
					sweeperCalMulUnit(dat, i, v)
				}

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

//V2

type unit struct {
	state byte

	beside []int
	corner []int
	around []int
}
type SweeperMap struct {
	unit [480]unit

	step      int
	valueRoot map[int]int
	zeroRoot  map[int]int
}

var sw SweeperMap

func SweeperInit() {
	sw = SweeperCreateMap()
}

func SweeperCreateMap() SweeperMap {
	rand.Seed(time.Now().Unix())
	sm := SweeperMap{}
	sm.valueRoot = make(map[int]int)
	sm.zeroRoot = make(map[int]int)
	for i := 0; i < 480; i++ {
		sm.unit[i].beside = getBesideIndex(i)
		sm.unit[i].corner = getCornerIndex(i)
		sm.unit[i].around = getAroundIndex(i)
	}
	return sm
}

func datIsValue(dat byte) bool {
	if dat != UNKNOWUNIT && dat != EMPYTUNIT && dat != SAFEUNIT && dat != SWEEPUNIT && dat != SWEEPDIDUNIT {
		return true
	} else {
		return false
	}
}

func sweeperCalSingleUnit(dat []byte, n int, around []int) bool {
	var update bool = false
	emptyIndex := getEmptyIndex(dat, around)
	aroundState := getState(dat, around)
	aroundBombNum := getBombNum(aroundState)

	if dat[n] == (byte)(len(emptyIndex))+aroundBombNum {
		for i := 0; i < len(emptyIndex); i++ {
			dat[emptyIndex[i]] = SWEEPUNIT
			update = true
		}
	}
	if dat[n] == aroundBombNum && dat[n] != 0 && dat[n] != SWEEPUNIT {
		for i := 0; i < len(emptyIndex); i++ {
			dat[emptyIndex[i]] = SAFEUNIT
			update = true
		}
	}
	return update
}

/*
func SweeperCal(dat []byte) []byte {
	var update bool = true
	var cnt int = 0

	for update == true {
		update = false
		for i := 0; i < 480; i++ {
			if dat[i] == EMPYTUNIT || dat[i] == SWEEPDIDUNIT {
				continue
			}
			around := getAroundIndex(i)
			if (byte)(len(getEmptyIndex(dat, around))) == 0 {
				continue
			}

			beside := sw.unit[i].beside
			around = sw.unit[i].around

			up := sweeperCalSingleUnit(dat, i, around)
			if up {
				update = up
			}

			for _, v := range beside {
				if datIsValue(dat[v]) {
					sweeperCalMulUnit(dat, i, v)
				}
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
*/
func sweeperSetDat(sw SweeperMap, dat []byte) int {
	var notZeroCnt int = 0
	for i := 0; i < len(dat); i++ {
		sw.unit[i].state = dat[i]
		if datIsValue(dat[i]) {
			sw.valueRoot[i] = i
			notZeroCnt++
		} else {
			sw.zeroRoot[i] = i
		}
	}
	return notZeroCnt
}

const SWEEPERSTEPINIT int = 0
const SWEEPERSTEPCAL int = 10
const SWEEPERSTEPRAMDON int = 20

func sweeperCalOnce(sw SweeperMap, dat []byte) bool {
	var update = false

	for _, v := range sw.valueRoot {
		aroundIndex := getAroundIndex(v)
		if (byte)(len(getEmptyIndex(dat, aroundIndex))) == 0 {
			delete(sw.valueRoot, v)
			continue
		}
		if datIsValue(dat[v]) == false {
			delete(sw.valueRoot, v)
			continue
		}
		besideIndex := sw.unit[v].beside
		aroundIndex = sw.unit[v].around

		up := sweeperCalSingleUnit(dat, v, aroundIndex)
		if up {
			update = up
		}

		for _, beside := range besideIndex {
			if datIsValue(dat[beside]) {
				sweeperCalMulUnit(dat, v, beside)
			}
		}
	}
	return update
}

func sweeperCal(sw SweeperMap, dat []byte) bool {
	var updateOnce bool = false
	for {
		update := sweeperCalOnce(sw, dat)
		if update {
			updateOnce = true
		} else {
			if updateOnce {
				return true
			} else {
				return false
			}
		}
	}
}

func getBombProbability(dat []byte, n int) int {
	var pro int = 0
	if dat[n] != 0 {
		return 0
	}
	aroundIndex := getAroundIndex(n)
	for _, v := range aroundIndex {
		around := getAroundIndex(v)
		aroundState := getState(dat, around)
		aroundBombNum := getBombNum(aroundState)
		empytNum := len(getEmptyIndex(dat, around))
		pro = pro + (int)(empytNum)/(int)(dat[v]-aroundBombNum)
	}
	return pro
}

func SweeperCal(sw SweeperMap, dat []byte) []byte {
	var notZeroCnt int = 0
	sw.step = SWEEPERSTEPINIT
	for {
		switch sw.step {
		case SWEEPERSTEPINIT:
			notZeroCnt = sweeperSetDat(sw, dat)
			notZeroCnt = notZeroCnt
			sw.step = SWEEPERSTEPCAL
			break

		case SWEEPERSTEPCAL:
			if sweeperCal(sw, dat) == true {
				return dat
			} else {
				sw.step = SWEEPERSTEPRAMDON
			}
			break

		case SWEEPERSTEPRAMDON:
			/*
				for _, v := range sw.zeroRoot {
					pro := getBombProbability(dat, v)
					pro = pro
				}
			*/
			index := getRandSweep()
			if dat[index] != SWEEPUNIT && dat[index] == 0 {
				dat[index] = SAFEUNIT
			}
			return dat
		}
	}
}

func debug() {
	var a string
	fmt.Scanf("%s", &a)
}
