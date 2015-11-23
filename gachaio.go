package gacha

type gachaIO interface {
	ReadUserGacha(uid int, gid int) (map[string]interface{}, error)
	WriteUserGacha(uid int, gid int, data map[string]interface{}) error
	DelUserGacha(uid int, gid int) error
}


