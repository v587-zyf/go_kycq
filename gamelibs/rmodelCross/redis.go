package rmodelCross

import (
	"cqserver/golibs/redisdb"
	"fmt"
	"github.com/pkg/errors"
)

const (
	CROSS_USER_REDIS = 1
)
var (
	redisMap = make(map[int]*redisdb.RedisDb)
)

func InitCrossMap(id int, network, address, password string, db int) error {

	if redisMap == nil {
		redisMap = make(map[int]*redisdb.RedisDb)
	}
	if redisMap[id] != nil {
		fmt.Println(fmt.Sprintf("init Cross map id:%v,address:%v has init", id, address))
		return errors.New("redis has init")
	}
	r := &redisdb.RedisDb{}
	err := r.Init(network, address, password, db)
	if err != nil {
		return err
	}
	redisMap[id] = r

	return nil
}

func GetRedisCrossDb(id int) *redisdb.RedisDb {
	return redisMap[id]
}
