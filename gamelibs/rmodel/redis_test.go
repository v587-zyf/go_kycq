package rmodel

import (
	"fmt"
	"testing"
	"time"

	"cqserver/gamelibs/beans"
	"cqserver/golibs/redisdb"
)

var inited = false

var sampileKindSet = 0
var sampileKindGet = 1
var sampileKindAll = 2

func testInit() {
	if inited {
		return
	}
	//err := Init(&beans.RedisConfig{Network: "tcp", Address: "192.168.5.25:6379", Password: "hoodinn.HOODINN", DB: 0}, 1)
	err := Init(&beans.RedisConfig{Network: "tcp", Address: "127.0.0.1:6379", Password: "hoodinn.HOODINN", DB: 0}, 1)
	if err != nil {
		fmt.Printf("err  err:%v", err)
	}

	reply := redisDb.Set("test", "10")
	if reply.Err != nil {
		fmt.Printf("set  err:%v", err)
	}

	redisDb.SAdd("testset", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	inited = true
}

func BenchmarkRedisGet(b *testing.B) {
	testInit()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		redisDb.Get("test")
	}
}

func BenchmarkRedisRawGet(b *testing.B) {
	testInit()
	c := redisDb.GetConn()
	defer c.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Do("GET", "test")
	}
}

func BenchmarkRedisRawSet(b *testing.B) {
	testInit()
	c := redisDb.GetConn()
	defer c.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Do("SET", "test", "100")
	}
}

func BenchmarkRedisRawGetPipelining(b *testing.B) {
	testInit()
	c := redisDb.GetConn()
	defer c.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Send("GET", "test")
		c.Flush()
		c.Receive()
	}
}

func BenchmarkRedisGetInt(b *testing.B) {
	testInit()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		redisDb.Get("test").Int()
	}
}

func BenchmarkRedisIsMember(b *testing.B) {
	testInit()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		redisDb.SIsMember("testset", 10)
	}
}

func batchFor(count, kind int, b *testing.B) {
	testInit()
	samples := getSamples(count, kind)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		redisDb.MultiExec(samples)
	}
}
func batchMulti(count, kind int, b *testing.B) {
	testInit()
	samples := getSamples(100, sampileKindGet)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		redisDb.MultiExec(samples)
	}
}

func BenchmarkBatchForGet10(b *testing.B)   { batchFor(10, sampileKindGet, b) }
func BenchmarkBatchMulGet10(b *testing.B)   { batchMulti(10, sampileKindGet, b) }
func BenchmarkBatchForGet50(b *testing.B)   { batchFor(50, sampileKindGet, b) }
func BenchmarkBatchMulGet50(b *testing.B)   { batchMulti(50, sampileKindGet, b) }
func BenchmarkBatchForGet100(b *testing.B)  { batchFor(100, sampileKindGet, b) }
func BenchmarkBatchMulGet100(b *testing.B)  { batchMulti(100, sampileKindGet, b) }
func BenchmarkBatchForGet1000(b *testing.B) { batchFor(1000, sampileKindGet, b) }
func BenchmarkBatchMulGet1000(b *testing.B) { batchMulti(1000, sampileKindGet, b) }

func BenchmarkBatchForSet10(b *testing.B)   { batchFor(10, sampileKindSet, b) }
func BenchmarkBatchMulSet10(b *testing.B)   { batchMulti(10, sampileKindSet, b) }
func BenchmarkBatchForSet50(b *testing.B)   { batchFor(50, sampileKindSet, b) }
func BenchmarkBatchMulSet50(b *testing.B)   { batchMulti(50, sampileKindSet, b) }
func BenchmarkBatchForSet100(b *testing.B)  { batchFor(100, sampileKindSet, b) }
func BenchmarkBatchMulSet100(b *testing.B)  { batchMulti(100, sampileKindSet, b) }
func BenchmarkBatchForSet1000(b *testing.B) { batchFor(1000, sampileKindSet, b) }
func BenchmarkBatchMulSet1000(b *testing.B) { batchMulti(1000, sampileKindSet, b) }

func BenchmarkBatchForAll10(b *testing.B)   { batchFor(10, sampileKindAll, b) }
func BenchmarkBatchMulAll10(b *testing.B)   { batchMulti(10, sampileKindAll, b) }
func BenchmarkBatchForAll50(b *testing.B)   { batchFor(50, sampileKindAll, b) }
func BenchmarkBatchMulAll50(b *testing.B)   { batchMulti(50, sampileKindAll, b) }
func BenchmarkBatchForAll100(b *testing.B)  { batchFor(100, sampileKindAll, b) }
func BenchmarkBatchMulAll100(b *testing.B)  { batchMulti(100, sampileKindAll, b) }
func BenchmarkBatchForAll1000(b *testing.B) { batchFor(1000, sampileKindAll, b) }
func BenchmarkBatchMulAll1000(b *testing.B) { batchMulti(1000, sampileKindAll, b) }

func getSamples(count, kind int) []*redisdb.CommandEntry {
	samples := make([]*redisdb.CommandEntry, 0, count)
	for i := 0; i < count; i++ {
		if kind&sampileKindSet == sampileKindSet {
			samples = append(samples, &redisdb.CommandEntry{
				Command: "set",
				Args:    []interface{}{fmt.Sprintf("test_%d", i), fmt.Sprintf("tttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt%d", i)},
			})
		}
		if kind&sampileKindGet == sampileKindGet {
			samples = append(samples, &redisdb.CommandEntry{
				Command: "get",
				Args:    []interface{}{fmt.Sprintf("test_%d", i)},
			})
		}
	}
	return samples
}

func TestRedisGet(t *testing.T) {
	testInit()
	Temple.SetDefender(1, 1, 555, 1000)
	defenderId, endTs, err := Temple.GetOneBuildingDefender(1, 1)
	if err != nil {
		t.Fatal(err)
	}
	if defenderId != 555 || endTs != 1000 {
		t.Fatal(fmt.Sprintf("TestRedisGet err:%d,%d", defenderId, endTs))
	}
}

func TestRedis(t *testing.T) {
	testInit()
	redisDb.Hmset("hk111", "w", 1)
	redisDb.Hmset("hk111", "q", 2)
	redisDb.Hmset("hk111", "wq", 3)

	value, err := redisDb.Hgetall("hk111")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("---", value)

	keys := make([]interface{}, 0)
	keys = append(keys, "hk111")
	keys = append(keys, "w")
	keys = append(keys, "fff")
	keys = append(keys, "wq")
	v, err := redisDb.HmgetStr(keys...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("---", v)
	delkeys := make([]interface{}, 0)
	delkeys = append(delkeys, "hk111")
	delkeys = append(delkeys, "w")
	delkeys = append(delkeys, "fff")
	_, err = redisDb.HDel(delkeys...)
	if err != nil {
		t.Fatal(err)
	}
	vv, err := redisDb.Hgetall("hk111")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("---", vv)
	redisDb.SAdd("saddwq", "ww1", "ww2", "ww2")
	aa, _ := redisDb.Smembers("saddwq22").ValuesStringSlice()
	fmt.Println("---", aa)

	stemkeys := make([]interface{}, 0)
	stemkeys = append(stemkeys, "saddwq")
	stemkeys = append(stemkeys, "ww2")
	redisDb.SRem(stemkeys...)
	bb, _ := redisDb.Smembers("saddwq").ValuesStringSlice()
	fmt.Println("---", bb)

}
func TestEnterAt(t *testing.T) {
	testInit()

	now, err := time.Parse("2006-01-02 15:04:05", "2018-06-24 12:00:00")
	err = WhiteList.SetAllowEnterAt("")
	if err != nil {
		t.Fatal(err)
	}
	enterAt, err := WhiteList.GetAllowEnterAt()
	if err != nil {
		t.Fatal(err)
	}
	if !enterAt.IsZero() {
		t.Fatal("should'nt happend", enterAt)
	}
	if now.Before(enterAt) {
		t.Fatal("should'nt happend:", enterAt)
	}

	err = WhiteList.SetAllowEnterAt("2018-06-30 12:00:00")
	if err != nil {
		t.Fatal(err)
	}
	enterAt, err = WhiteList.GetAllowEnterAt()
	if err != nil {
		t.Fatal(err)
	}
	if now.After(enterAt) {
		t.Fatal("should'nt happend,", enterAt)
	}

	err = WhiteList.SetAllowEnterAt("2018-06-10 12:00:00")
	if err != nil {
		t.Fatal(err)
	}
	enterAt, err = WhiteList.GetAllowEnterAt()
	if err != nil {
		t.Fatal(err)
	}
	if now.Before(enterAt) {
		t.Fatal("should'nt happend,", enterAt)
	}
}
