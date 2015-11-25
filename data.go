package gacha

// itemValueMap : 각 아이템의 가치의 정보를 가지는 맵. 키:아이템 id, 값:아이템의 가치를 보이는 정수혐
var itemValueMap = make(map[int]int)
// 각 가챠에 속하는 아이템 리스트를 가지는 맵. 키:가챠 id, 값:아이템 리스트
var gachaMap = make(map[int][]int)

func SetItemVal(itemId int, itemValue int) {
	itemValueMap[itemId] = itemValue
}

func GetItemVal(itemId int) int {
	iv, ok := itemValueMap[itemId]


}

func AddGacha(gid int, itemId int) {
	if itemList, ok := gachaMap[gid]; ok {
		itemList = append(itemList, itemId)
	} else {
		newItemList := []string{int}
		gachaMap[gid] = newItemList
	}
}

// relic id -> item list
func GetRelicList(rid int) []int {
	return gachaMap[rid]
}