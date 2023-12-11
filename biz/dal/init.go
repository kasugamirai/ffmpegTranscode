package dal

import (
	"freefrom.space/videoTransform/biz/dal/mysql"
	"freefrom.space/videoTransform/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
