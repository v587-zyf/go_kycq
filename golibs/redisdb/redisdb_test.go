package redisdb

import (
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	redisDb := &RedisDb{}
	err := redisDb.Init("tcp", "127.0.0.1:6379", "hoodinn.HOODINN", 0)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCopy(t *testing.T) {
	redisDb := &RedisDb{}
	err := redisDb.Init("tcp", "127.0.0.1:6379", "hoodinn.HOODINN", 0)
	if err != nil {
		t.Fatal(err)
	}

	err = redisDb.Copy("sourceA", "destinationB", time.Hour)
	if err != nil {
		t.Fatal(err)
	}

}
