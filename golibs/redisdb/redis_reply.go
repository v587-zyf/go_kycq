package redisdb

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type RedisReply struct {
	Reply interface{}
	Err   error
}

type HashValues map[string]string

func reply(reply interface{}, err error) *RedisReply {
	return &RedisReply{reply, err}
}

func (this *RedisReply) Raw() (interface{}, error) {
	return this.Reply, this.Err
}

func (this *RedisReply) String() (string, error) {
	return this.StringDef("")
}

func (this *RedisReply) StringDef(def string) (string, error) {
	if this.Err == redis.ErrNil {
		return def, nil
	}
	return redis.String(this.Reply, this.Err)
}

func (this *RedisReply) MustString() (string, error) {
	return redis.String(this.Reply, this.Err)
}

func (this *RedisReply) Int() (int, error) {
	return this.IntDef(0)
}

func (this *RedisReply) IntDef(def int) (int, error) {
	if this.Err == redis.ErrNil {
		return def, nil
	}
	if this.Err == nil && this.Reply == nil {
		return def, nil
	}
	return redis.Int(this.Reply, this.Err)
}

func (this *RedisReply) MustInt() (int, error) {
	return redis.Int(this.Reply, this.Err)
}

func (this *RedisReply) Int64() (int64, error) {
	return this.Int64Def(0)
}

func (this *RedisReply) Int64Def(def int64) (int64, error) {
	if this.Err == redis.ErrNil {
		return def, nil
	}
	if this.Err == nil && this.Reply == nil {
		return def, nil
	}
	return redis.Int64(this.Reply, this.Err)
}

func (this *RedisReply) MustInt64() (int64, error) {
	return redis.Int64(this.Reply, this.Err)
}

func (this *RedisReply) Float64() (float64, error) {
	return this.Float64Def(0)
}

func (this *RedisReply) Float64Def(def float64) (float64, error) {
	if this.Err == redis.ErrNil {
		return def, nil
	}
	return redis.Float64(this.Reply, this.Err)
}

func (this *RedisReply) MustFloat64() (float64, error) {
	return redis.Float64(this.Reply, this.Err)
}

func (this *RedisReply) Bool() (bool, error) {
	return this.BoolDef(false)
}

func (this *RedisReply) BoolDef(def bool) (bool, error) {
	if this.Err == redis.ErrNil {
		return def, nil
	}
	return redis.Bool(this.Reply, this.Err)
}

func (this *RedisReply) MustBool() (bool, error) {
	return redis.Bool(this.Reply, this.Err)
}

func (this *RedisReply) Bytes() ([]byte, error) {
	if this.Err == redis.ErrNil {
		return nil, nil
	}
	return redis.Bytes(this.Reply, this.Err)
}

func (this *RedisReply) MustBytes() ([]byte, error) {
	return redis.Bytes(this.Reply, this.Err)
}

func (this *RedisReply) Strings() ([]string, error) {
	if this.Err == redis.ErrNil {
		return nil, nil
	}
	return redis.Strings(this.Reply, this.Err)
}

func (this *RedisReply) MustStrings() ([]string, error) {
	return redis.Strings(this.Reply, this.Err)
}

func (this *RedisReply) Values() ([]interface{}, error) {
	if this.Err == redis.ErrNil {
		return nil, nil
	}
	return redis.Values(this.Reply, this.Err)
}

func (this *RedisReply) ValuesIntSlice() ([]int, error) {
	if this.Err == redis.ErrNil {
		return nil, nil
	}
	v, err := redis.Values(this.Reply, this.Err)
	if err != nil {
		return nil, err
	}
	ids := make([]int, 0)
	if err = redis.ScanSlice(v, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

func (this *RedisReply) ValuesFloat64Slice() ([]float64, error) {
	if this.Err == redis.ErrNil {
		return nil, nil
	}
	v, err := redis.Values(this.Reply, this.Err)
	if err != nil {
		return nil, err
	}
	ids := make([]float64, 0)
	if err = redis.ScanSlice(v, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

func (this *RedisReply) ValuesInt64Slice() ([]int64, error) {
	if this.Err == redis.ErrNil {
		return nil, nil
	}
	v, err := redis.Values(this.Reply, this.Err)
	if err != nil {
		return nil, err
	}
	ids := make([]int64, 0)
	if err = redis.ScanSlice(v, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

func (this *RedisReply) ValuesInt64SliceForE10() ([]int64, error) {
	if this.Err == redis.ErrNil {
		return nil, nil
	}
	v, err := redis.Values(this.Reply, this.Err)
	if err != nil {
		return nil, err
	}
	ids := make([]string, 0)
	if err = redis.ScanSlice(v, &ids); err != nil {
		return nil, err
	}
	nums := make([]int64,len(ids))
	for k,v := range ids {
		var (
			new float64
		)
		n, err := fmt.Sscanf(v, "%e", &new)
		if err != nil {
			fmt.Println(err.Error())
		} else if 1 != n {
			fmt.Println("n is not one")
		}
		nums[k] = int64(new)
	}
	return nums, nil
}

func (this *RedisReply) ValuesStringSlice() ([]string, error) {
	if this.Err == redis.ErrNil {
		return nil, nil
	}
	v, err := redis.Values(this.Reply, this.Err)
	if err != nil {
		return nil, err
	}
	ids := make([]string, 0)
	if err = redis.ScanSlice(v, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

func (this *RedisReply) MustValues() ([]interface{}, error) {
	return redis.Values(this.Reply, this.Err)
}

func (this *RedisReply) HashValues() (HashValues, error) {
	reply, err := redis.Strings(this.Reply, this.Err)
	if err != nil {
		return nil, err
	}
	values := make(HashValues)
	for i := 0; i < len(reply); i = i + 2 {
		values[reply[i]] = reply[i+1]
	}
	return values, nil
}

func (this HashValues) Get(key string) string {
	v, ok := this[key]
	if ok {
		return v
	}
	return ""
}

func (this HashValues) Empty() bool {
	return len(this) == 0
}
