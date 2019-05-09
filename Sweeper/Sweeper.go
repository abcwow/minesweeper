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

func sweeperCalBesideUnit(dat []byte, a int, b int) bool {
	var update = false
	if dat[a] < dat[b] {
		return update
	}
	aAround := getAroundIndex(a)
	bAround := getAroundIndex(b)
	aStaticIndex, bStaticIndex, commIndex := getArea(aAround, bAround, a, b)
	aStaticState := getState(dat, aStaticIndex)
	bStaticState := getState(dat, bStaticIndex)
	commState := getState(dat, commIndex)
	aStaticBombNum := getBombNum(aStaticState)
	bStaticBombNum := getBombNum(bStaticState)
	commBombNum := getBombNum(commState)
	commBombNum = commBombNum

	aEmptyIndex := getEmptyIndex(dat, aStaticIndex)
	aEmptyNum := (byte)(len(aEmptyIndex))
	bEmptyIndex := getEmptyIndex(dat, bStaticIndex)
	bEmptyNum := (byte)(len(bEmptyIndex))

	if dat[a] == dat[b] {
		if aStaticBombNum < bStaticBombNum && (byte)(aEmptyNum) == bStaticBombNum-aStaticBombNum {
			for _, v := range aEmptyIndex {
				dat[v] = SWEEPUNIT
				update = true
			}
		}
		if aStaticBombNum > bStaticBombNum && (byte)(bEmptyNum) == aStaticBombNum-bStaticBombNum {
			for _, v := range bEmptyIndex {
				dat[v] = SWEEPUNIT
				update = true
			}
		}
		if aStaticBombNum == bStaticBombNum && len(aEmptyIndex) == 0 {
			for _, v := range bEmptyIndex {
				dat[v] = SAFEUNIT
				update = true
			}
		}
		if aStaticBombNum == bStaticBombNum && len(bEmptyIndex) == 0 {
			for _, v := range aEmptyIndex {
				dat[v] = SAFEUNIT
				update = true
			}
		}
	} else {
		if (byte)(aEmptyNum) <= dat[a]-dat[b]+bStaticBombNum-aStaticBombNum && bStaticBombNum > aStaticBombNum {
			for _, v := range aEmptyIndex {
				dat[v] = SWEEPUNIT
				update = true
			}
		}
		/*
			if (byte)(bEmptyNum) <= dat[a]-dat[b]+bStaticBomnNum-aStaticBombNum {
				for _, v := range bEmptyIndex {
					dat[v] = SAFEUNIT
					update = true
				}
			}
		*/
	}
	return update
}

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

func SweeperCreateMap() SweeperMap {
	rand.Seed(time.Now().Unix())
	sweeperMap := SweeperMap{}
	sweeperMap.valueRoot = make(map[int]int)
	sweeperMap.zeroRoot = make(map[int]int)
	for i := 0; i < 480; i++ {
		sweeperMap.unit[i].beside = getBesideIndex(i)
		sweeperMap.unit[i].corner = getCornerIndex(i)
		sweeperMap.unit[i].around = getAroundIndex(i)
	}
	return sweeperMap
}

func sweeperReset(sweeperMap *SweeperMap) {
	sweeperMap.valueRoot = make(map[int]int)
	sweeperMap.zeroRoot = make(map[int]int)
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

func sweeperCalMulUnit(dat []byte, n int, besideIndex []int) bool {
	var update bool = false
	for _, beside := range besideIndex {
		if datIsValue(dat[beside]) {
			up := sweeperCalBesideUnit(dat, n, beside)
			if up {
				update = up
			}
		}
	}
	return update
}

func sweeperSetDat(sw SweeperMap, dat []byte) int {
	var BombCnt int = 0
	for i := 0; i < len(dat); i++ {
		sw.unit[i].state = dat[i]
		if datIsValue(dat[i]) {
			sw.valueRoot[i] = i
		} else {
			if dat[i] == SWEEPUNIT || dat[i] == SWEEPDIDUNIT {
				BombCnt++
			} else if dat[i] == UNKNOWUNIT {
				sw.zeroRoot[i] = i
			}
		}
	}
	return BombCnt
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

		aroundIndex = sw.unit[v].around
		up := sweeperCalSingleUnit(dat, v, aroundIndex)
		if up {
			update = up
		}

		besideIndex := sw.unit[v].beside
		up = sweeperCalMulUnit(dat, v, besideIndex)
		if up {
			update = up
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

func calAroundProbability(dat []byte, aroundIndex []int) float32 {
	var pro float32 = 1
	for _, v := range aroundIndex {
		if datIsValue(dat[v]) == false {
			continue
		}
		around := getAroundIndex(v)
		aroundState := getState(dat, around)
		aroundBombNum := getBombNum(aroundState)
		empytNum := len(getEmptyIndex(dat, around))
		pro = pro * (1 - (float32)(dat[v]-aroundBombNum)/(float32)(empytNum))
	}
	return 1 - pro
}

func getBombProbability(dat []byte, n int, bombCnt int, zeroCnt int) float32 {
	if dat[n] != 0 {
		return 1
	}
	aroundIndex := getAroundIndex(n)
	empytNum := len(getEmptyIndex(dat, aroundIndex))
	aroundState := getState(dat, aroundIndex)
	aroundBombNum := getBombNum(aroundState)
	if empytNum == len(aroundIndex)-(int)(aroundBombNum) {
		return (float32)(bombCnt) / (float32)(zeroCnt)
	} else {
		return calAroundProbability(dat, aroundIndex)
	}
}

func getRandSweep() int {
	return rand.Intn(480)
}

func SweeperCal(sw SweeperMap, dat []byte) []byte {
	var bombCnt int = 0
	sw.step = SWEEPERSTEPINIT
	for {
		switch sw.step {
		case SWEEPERSTEPINIT:
			sweeperReset(&sw)
			bombCnt = sweeperSetDat(sw, dat)
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
			var pro float32
			var proMin float32 = 1
			var indexMin int = 0

			if len(sw.valueRoot) == 0 {
				index := getRandSweep()
				dat[index] = SAFEUNIT
			} else {
				for _, v := range sw.zeroRoot {
					pro = getBombProbability(dat, v, 99-bombCnt, len(sw.zeroRoot))
					if pro < proMin {
						proMin = pro
						indexMin = v
					}
				}
				fmt.Printf("b:%d, z:%d\n", 99-bombCnt, len(sw.zeroRoot))
				fmt.Printf("indexMin: %d, pro: %f\n", indexMin, proMin)
				dat[indexMin] = SAFEUNIT
			}
			return dat
		}
	}
}
