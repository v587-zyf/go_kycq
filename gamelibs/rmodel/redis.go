package rmodel

import (
	"cqserver/gamelibs/beans"
	"cqserver/golibs/redisdb"
)

var (
	redisDb      = &redisdb.RedisDb{}
)

func Init(rc *beans.RedisConfig, sId int) error {
	return redisDb.Init(rc.Network, rc.Address, rc.Password, rc.DB)
}

func RedisDb() *redisdb.RedisDb {
	return redisDb
}