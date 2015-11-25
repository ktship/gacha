package gacha
import (
	"strconv"
	"runtime"
	"fmt"
)

/*
가챠(유물) 시스템 테스트

기본 테스트
	1. 쿠키런의 유물 보상 시스템.
	   - 여러가지 보상중에 하나씩 주는 시스템. 리셋 가능. 기본 보상 셋팅 가능.
 */

type itemStatus int
type relicId	int
type itemId		int

type client struct {
	uid				int
	relics			map[relicId]map[itemId]itemStatus
}

type clientTask struct {
	curNum			int
	remainedTime	int
}













func init() {
	fmt.Printf("Running On %s, %s, %s, %d-bit \n", runtime.Compiler, runtime.GOARCH, runtime.GOOS, strconv.IntSize)

	// item 가치 테이블
	// 0:1000, 1:900, 2:800, 3:500, 4:300, 5:200, 6:100, 7:50, 8:20, 9:10
	SetItemVal(0, 1000)
	SetItemVal(1, 900)
	SetItemVal(2, 800)
	SetItemVal(3, 500)
	SetItemVal(4, 300)
	SetItemVal(5, 200)
	SetItemVal(6, 100)
	SetItemVal(7, 50)
	SetItemVal(8, 20)
	SetItemVal(9, 10)

	// 유물 구성
	// 유물 id 0: 0,3,6,7,9, id 1: 7,8,9 id 2: 0,3,6
	AddGacha(0, 0)
	AddGacha(0, 3)
	AddGacha(0, 6)
	AddGacha(0, 7)
	AddGacha(0, 9)

	AddGacha(1, 7)
	AddGacha(1, 8)
	AddGacha(1, 9)

	AddGacha(2, 0)
	AddGacha(2, 3)
	AddGacha(2, 6)
}
