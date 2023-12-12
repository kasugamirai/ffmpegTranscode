package dal

import "freefrom.space/videoTransform/biz/dal/sqlite"

func Init() {
	//redis.Init()
	//mysql.Init()
	sqlite.Init()
}
