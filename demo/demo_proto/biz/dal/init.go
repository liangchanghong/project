package dal

import (
	"github.com/liangchanghong/project/gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/liangchanghong/project/gomall/demo/demo_proto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
