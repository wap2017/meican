package config

type Config struct {
	DislikeDishWordList       []string // "辣"
	DislikeRestaurantWordList []string // "丽"
	PreferDishWordList        []string // "猪排|汤粉|牛扒|鸡扒"
	PreferRestaurantWordList  []string // "便当"
	//已废弃
	//Floor    uint32 // 17|18|19
	FloorInd uint32 //楼层地址的下标
	Location string
}
