package redisdb

import (
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
)

var log = logger.Get("default", true)

type RedisDb struct {
	redisPool *redis.Pool
}

type CommandEntry struct {
	Command string
	Args    []interface{}
}

func (this *RedisDb) Init(network, address, password string, db int) error {
	log.Info("Create redis connect pool: %v %v %v", network, address, password)
	this.redisPool = &redis.Pool{
		MaxIdle:     100,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			opts := []redis.DialOption{redis.DialDatabase(db)}
			if len(password) > 0 {
				opts = append(opts, redis.DialPassword(password))
			}
			var c redis.Conn
			var err error
			for i := 0; i < 3; i++ {
				c, err = redis.Dial(network, address, opts...)
				if err != nil {
					time.Sleep(10 * time.Millisecond)
				} else if c != nil {
					break
				}
			}
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	if _, err := this.Do("PING").String(); err != nil {
		return err
	}
	return nil
}

func (this *RedisDb) Get(key string) *RedisReply {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("GET", key))
}

func (this *RedisDb) Del(args ...interface{}) *RedisReply {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("DEL", args...))
}

func (this *RedisDb) Set(key string, value interface{}) *RedisReply {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("SET", key, value))
}

func (this *RedisDb) MSet(args ...interface{}) *RedisReply {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("MSET", args...))
}

//SetWithExpireSecond
//为了代码可读性好，请尽量使用 SetWithExpire
func (this *RedisDb) SetWithExpireSecond(key string, value interface{}, expire int) *RedisReply {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("SETEX", key, expire, value))
}

//SetWithExpire
//setex
func (this *RedisDb) SetWithExpire(key string, value interface{}, expire time.Duration) *RedisReply {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("SETEX", key, int(expire.Seconds()), value))
}

func (this *RedisDb) Smembers(key string) *RedisReply {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("SMEMBERS", key))
}

func (this *RedisDb) Srandmember(key string, count int) *RedisReply {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("SRANDMEMBER", key, count))
}

func (this *RedisDb) Do(commandName string, args ...interface{}) *RedisReply {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do(commandName, args...))
}

func (this *RedisDb) GetConn() redis.Conn {
	c := this.redisPool.Get()
	return c
}

func (this *RedisDb) FifoAdd(key, element interface{}, maxLength int) error {
	c := this.redisPool.Get()
	defer c.Close()

	newLength, err := redis.Int(c.Do("LPUSH", key, element))
	if err != nil {
		return err
	}
	if newLength > maxLength {
		_, err := c.Do("RPOP", key)
		if err != nil {
			fmt.Printf("FifoAdd err: = %+v\n", err)
		}
	}
	return nil

}

func (this *RedisDb) FifosAdd(key string, maxLength int, argss ...interface{}) error {
	c := this.redisPool.Get()
	defer c.Close()
	//	for _, v := range argss {
	//		fmt.Println("ss=", v)
	//	}
	newLength, err := redis.Int(c.Do("LPUSH", argss...))
	if err != nil {
		return err
	}
	if newLength > maxLength {
		_, err := c.Do("LTRIM", key, 0, maxLength-1)
		if err != nil {
			fmt.Printf("FifosAdd err: = %+v\n", err)
		}
	}
	return nil
}

func (this *RedisDb) HmsetStrStrMap(key string, m map[string]string) (string, error) {
	c := this.redisPool.Get()
	defer c.Close()
	args := make([]interface{}, 0, len(m)*2+1)
	args = append(args, key)
	for k, v := range m {
		args = append(args, k, v)
	}
	return redis.String(c.Do("HMSET", args...))
}

func (this *RedisDb) HmsetStringMap(key string, m map[int]string) (string, error) {
	c := this.redisPool.Get()
	defer c.Close()
	args := make([]interface{}, 0, len(m)*2+1)
	args = append(args, key)
	for k, v := range m {
		args = append(args, k, v)
	}
	return redis.String(c.Do("HMSET", args...))
}

func (this *RedisDb) HmsetIntMap(key string, m map[int]int) (string, error) {
	c := this.redisPool.Get()
	defer c.Close()
	args := make([]interface{}, 0, len(m)*2+1)
	args = append(args, key)
	for k, v := range m {
		args = append(args, k, v)
	}
	return redis.String(c.Do("HMSET", args...))
}

func (this *RedisDb) Hmset(kv ...interface{}) (string, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.String(c.Do("HMSET", kv...))
}
func (this *RedisDb) HgetInt(key string, subKey interface{}) (int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.Int(c.Do("HGET", key, subKey))
}

func (this *RedisDb) HSet(key string, field, value int) error {
	c := this.redisPool.Get()
	defer c.Close()
	_, err := c.Do("HSET", key, field, value)
	return err
}

func (this *RedisDb) HgetIntDef(key string, subKey interface{}, def int) (int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("HGET", key, subKey)).IntDef(def)
}

func (this *RedisDb) HgetStr(key string, subKey interface{}) (string, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.String(c.Do("HGET", key, subKey))
}

func (this *RedisDb) HgetByJson(key, subKey string) ([]byte, error) {
	c := this.redisPool.Get()
	defer c.Close()
	valueGet, _ := redis.Bytes(c.Do("HGET", key, subKey))
	return valueGet, nil
}

func (this *RedisDb) HmgetStr(key ...interface{}) ([]string, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.Strings(c.Do("HMGET", key...))
}

func (this *RedisDb) Hmget(key ...interface{}) (map[string]int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.IntMap(c.Do("HMGET", key...))
}

func (this *RedisDb) SetChivalryUser(key string, idStr interface{}) {
	c := this.redisPool.Get()
	defer c.Close()
	redis.String(c.Do("SET", key, idStr))
}

func (this *RedisDb) GetChivalryUser(key string) (string, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.String(c.Do("GET", key))
}

func (this *RedisDb) SetChivalry(key string, value interface{}) {
	c := this.redisPool.Get()
	defer c.Close()
	c.Do("SET", key, value)
}

func (this *RedisDb) GetChivalry(key string) (int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.Int(c.Do("GET", key))
}

func (this *RedisDb) DelChivalry(key interface{}) *RedisReply {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("DEL ", key))
}

func (this *RedisDb) Hgetall(key string) (map[string]string, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.StringMap(c.Do("HGETALL", key))
}

func (this *RedisDb) HgetallIntMap1(key string) (map[string]int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.IntMap(c.Do("HGETALL", key))
}

func (this *RedisDb) Hlen(key string) (int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.Int(c.Do("HLEN", key))
}

func (this *RedisDb) IncrBy(key string, delta int) (int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.Int(c.Do("incrby", key, delta))
}

func (this *RedisDb) Lrange(key string, minIndex, maxIndex int) ([]string, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.Strings(c.Do("LRANGE", key, minIndex, maxIndex))
}

func (this *RedisDb) LrangeInits(key string, minIndex, maxIndex int) ([]int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.Ints(c.Do("LRANGE", key, minIndex, maxIndex))
}

func (this *RedisDb) LpushAndLtrim(key string, msgStrs []string, minIndex, maxIndex int) (int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	args := make([]interface{}, 0, len(msgStrs))
	args = append(args, key)
	for _, v := range msgStrs {
		args = append(args, v)
	}
	newLength, err := redis.Int(c.Do("LPUSH", args...))
	if err != nil {
		return 0, err
	}
	if newLength > maxIndex+1 {
		_, err := c.Do("LTRIM", key, minIndex, maxIndex)
		if err != nil {
			fmt.Printf("LpushAndLtrim err: = %+v\n", err)
		}
	}
	return 0, nil
}

func (this *RedisDb) Hexists(key string, subKey interface{}) (int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.Int(c.Do("HEXISTS", key, subKey))
}

//Expire
//设置一个key过期．
func (this *RedisDb) Expire(key string, duration time.Duration) (int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.Int(c.Do("expire", key, int(duration.Seconds())))
}

func (this *RedisDb) HmgetIntMap(key string, subkey []int) (map[int]int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	args := make([]interface{}, 0, len(subkey)+1)
	args = append(args, key)
	for _, v := range subkey {
		args = append(args, v)
	}
	ints, err := redis.Ints(c.Do("HMGET", args...))
	if err != nil {
		return nil, err
	}
	m := make(map[int]int)
	for k, v := range subkey {
		m[v] = ints[k]
	}
	return m, nil
}

//ExpireSecond
//为了代码可读性好，请尽量使用 Expire
func (this *RedisDb) ExpireSecond(key string, second int) (int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.Int(c.Do("expire", key, second))
}

//注释掉了,因为tendis不把持 multi
//func (this *RedisDb) MultiExecOld(args []*CommandEntry) ([]interface{}, error) {
//c := this.redisPool.Get()
//defer c.Close()
//c.Send("MULTI")
//for _, entry := range args {
//c.Send(entry.Command, entry.Args...)
//}
//return redis.MultiBulk(c.Do("EXEC"))
//}

func (this *RedisDb) ZAdds(keyScoreSubKeys []interface{}) (int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.Int(c.Do("ZADD", keyScoreSubKeys...))
}

func (this *RedisDb) MultiExec(args []*CommandEntry) ([]interface{}, error) {
	c := this.redisPool.Get()
	defer c.Close()
	results := make([]interface{}, 0, len(args))
	for _, entry := range args {
		reply, errOne := c.Do(entry.Command, entry.Args...)
		results = append(results, reply)
		if errOne != nil {
			return results, errOne
		}
	}
	return results, nil
}

func (this *RedisDb) HIncrBy(key string, field interface{}, delta int) (int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.Int(c.Do("hincrby", key, field, delta))
}

func (this *RedisDb) HgetallIntMap(key string) (map[int]int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	originMap, err := redis.IntMap(c.Do("HGETALL", key))
	if err != nil {
		return nil, err
	}
	mapRet := make(map[int]int, len(originMap))
	for k, v := range originMap {
		intK, err := strconv.Atoi(k)
		if err != nil {
			return nil, err
		}
		mapRet[intK] = v
	}
	return mapRet, nil
}

func (this *RedisDb) HGetAllIntAndStringMap(key string) (map[int]string, error) {
	c := this.redisPool.Get()
	defer c.Close()
	originMap, err := redis.StringMap(c.Do("HGETALL", key))
	if err != nil {
		return nil, err
	}
	mapRet := make(map[int]string, len(originMap))
	for k, v := range originMap {
		intK, err := strconv.Atoi(k)
		if err != nil {
			return nil, err
		}
		mapRet[intK] = v
	}
	return mapRet, nil
}

func (this *RedisDb) SmembersInt(key string) ([]int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.Ints(c.Do("SMEMBERS", key))
}

func (this *RedisDb) HDelIntList(key string, ids []int) (int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	args := make([]interface{}, 0, len(ids)+1)
	args = append(args, key)
	for _, v := range ids {
		args = append(args, v)
	}
	return redis.Int(c.Do("HDEL", args...))
}

func (this *RedisDb) HDel(args ...interface{}) (int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.Int(c.Do("HDEL", args...))
}
func (this *RedisDb) SRem(args ...interface{}) (int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.Int(c.Do("SREM", args...))
}
func (this *RedisDb) SAdd(args ...interface{}) (int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.Int(c.Do("SADD", args...))
}
func (this *RedisDb) SIsMember(key, member interface{}) (bool, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.Bool(c.Do("sismember", key, member))
}

func (this *RedisDb) Publish(channel, message interface{}) (bool, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.Bool(c.Do("publish", channel, message))
}

func (this *RedisDb) ZAdd(key, score, subKey interface{}) (int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.Int(c.Do("ZADD", key, score, subKey))
}

func (this *RedisDb) HsetRank(key, score, subKey interface{}) (int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.Int(c.Do("ZADD", key, score, subKey))
}

func (this *RedisDb) GetHsetRank(key interface{}) *RedisReply {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("ZREVRANGE", key, 0, 100, "WITHSCORES"))
}

func (this *RedisDb) DeleteHsetRank(key interface{}) {
	c := this.redisPool.Get()
	defer c.Close()
	c.Do("DEL", key)
}

func (this *RedisDb) ZIncrby(key, score, subKey interface{}) (int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.Int(c.Do("ZINCRBY", key, score, subKey))
}

func (this *RedisDb) ZCard(key string) int {
	c := this.redisPool.Get()
	defer c.Close()
	num, err := redis.Int(c.Do("ZCARD", key))
	if err != nil {
		return -1
	}
	return num
}

func (this *RedisDb) ZRevrank(key string, id interface{}) int {
	c := this.redisPool.Get()
	defer c.Close()
	num, err := redis.Int(c.Do("ZREVRANK", key, id))
	if err != nil {
		return -1
	}
	return num
}

func (this *RedisDb) ZScore(key, subKey interface{}) (int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("ZSCORE", key, subKey)).IntDef(0)
}

func (this *RedisDb) ZScoreByFloat(key, subKey interface{}) (float64, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("ZSCORE", key, subKey)).Float64Def(0)
}

func (this *RedisDb) ZRank(key, subKey interface{}) (int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("ZRANK", key, subKey)).IntDef(-1)
}

func (this *RedisDb) ZRange(key string, min, max int) *RedisReply {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("ZRANGE", key, min, max, "WITHSCORES"))
}

//[2000 40000]    >= 2000 [2000,"+inf" ]  <= 2000 [-1,2000]
func (this *RedisDb) ZRangeByScore(key string, min, max interface{}) *RedisReply {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("ZRANGEBYSCORE", key, min, max, "WITHSCORES"))
}

func (this *RedisDb) ZRem(key string, subKey interface{}) *RedisReply {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("ZREM", key, subKey))
}

func (this *RedisDb) ZRevrange(key string, min, max int) *RedisReply {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("ZREVRANGE", key, min, max, "WITHSCORES"))
}

func (this *RedisDb) ZRevrangeIntSlice(key string, min, max int) ([]int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	stringSlices, err := reply(c.Do("ZREVRANGE", key, min, max, "WITHSCORES")).ValuesStringSlice()
	if err != nil {
		return nil, err
	}
	l := len(stringSlices)
	intArr := make([]int, l, l)
	for i, v := range stringSlices {
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
		intArr[i] = int(f)
	}
	return intArr, nil
}

func (this *RedisDb) ZRevrangeKv(key string, min, max int) ([]common.KV, error) {
	intArr, err := this.ZRevrangeIntSlice(key, min, max)
	if err != nil {
		return nil, err
	}
	l := len(intArr) / 2

	kvs := make([]common.KV, 0, l)
	for i := 0; i < l; i++ {
		kvs = append(kvs, common.KV{K: intArr[i*2], V: intArr[i*2+1]})
	}
	return kvs, nil
}

func (this *RedisDb) ZRevrangeStringMap(key string, min, max int) (map[string]string, error) {
	c := this.redisPool.Get()
	defer c.Close()
	stringMap, err := reply(c.Do("ZREVRANGE", key, min, max, "WITHSCORES")).HashValues()
	if err != nil {
		return nil, err
	}
	return stringMap, nil
}

func (this *RedisDb) ZRemRangeByRank(key string, min, max int) *RedisReply {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("ZREMRANGEBYRANK", key, min, max))
}

func (this *RedisDb) AddLock(key string, second int) {
	this.Set(key, 1)
	this.ExpireSecond(key, second)
}

func (this *RedisDb) GetLock(key string) bool {
	flag, err := this.Get(key).IntDef(-1)
	if err != nil {
		return false
	}
	if flag == -1 {
		return false
	}
	return true
}

func (this *RedisDb) UnLock(key string) {
	this.Del(key)
}

func (this *RedisDb) GetByJson(key string) ([]byte, error) {
	c := this.redisPool.Get()
	defer c.Close()
	valueGet, _ := redis.Bytes(c.Do("GET", key))
	return valueGet, nil
}

//Copy
//destinationTTL 是毫秒,如果没有过期时间,传0
func (this *RedisDb) Copy(sourceKey, destinationKey string, destinationTTL time.Duration) error {
	c := this.redisPool.Get()
	defer c.Close()
	result, err := c.Do("DUMP", sourceKey)
	if err != nil {
		return err
	}
	_, err = c.Do("RESTORE", destinationKey, int(destinationTTL.Seconds())*1000, result, "REPLACE")
	return err
}

func (this *RedisDb) Keys(pattern interface{}) ([]string, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("KEYS", pattern)).Strings()
}

func (this *RedisDb) FlushDb() (string, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("FLUSHDB")).String()
}

//获取reids池
func (this *RedisDb) GetRedisPool() *redis.Pool {
	return this.redisPool
}

//replay
func Reply(redisReply interface{}, err error) *RedisReply {
	return reply(redisReply, err)
}

func (this *RedisDb) LPushSlice(key string, data []int) *RedisReply {
	c := this.redisPool.Get()
	args := make([]interface{}, 0)
	args = append(args, key)
	for _, v := range data {
		args = append(args, v)
	}
	defer c.Close()
	return reply(c.Do("lpush", args...))
}

func (this *RedisDb) LPush(key string, data int) *RedisReply {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("lpush", key, data))
}

func (this *RedisDb) RPush(key string, data int) *RedisReply {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("rpush", key, data))
}

func (this *RedisDb) LRem(key string, data int) *RedisReply {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("lrem", key, 1, data))
}

func (this *RedisDb) LLen(key string) (int, error) {
	c := this.redisPool.Get()
	defer c.Close()
	return redis.Int(c.Do("llen", key))
}

func (this *RedisDb) LRange(key string) *RedisReply {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("lrange", key, 0, -1))
}

func (this *RedisDb) Sunionstore(targetKey string, sourceKeys ...string) *RedisReply {
	c := this.redisPool.Get()
	defer c.Close()
	args := make([]interface{}, 0)
	args = append(args, targetKey)
	for _, v := range sourceKeys {
		args = append(args, v)
	}
	return reply(c.Do("Sunionstore", args...))
}

func (this *RedisDb) ZunionstoreOne(targetKey string, sourceKeys string) error {
	c := this.redisPool.Get()
	defer c.Close()
	args := make([]interface{}, 0)
	args = append(args, targetKey, 1, sourceKeys)
	_, err := c.Do("Zunionstore", args...)
	if err != nil {
		logger.Error("Zunionstore targetKey:%v sourceKeys:%v,err:%v", targetKey, sourceKeys, err)
	}
	return err
}

func (this *RedisDb) Hvals(key string) *RedisReply {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("hvals", key))
}

func (this *RedisDb) Ttl(key string) *RedisReply {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("ttl", key))
}

func (this *RedisDb) HgetAll(key string) *RedisReply {
	c := this.redisPool.Get()
	defer c.Close()
	return reply(c.Do("hgetall", key))
}
