package cache

import (
	"IMChat/core/config"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/redis/go-redis/v9"
)

var testCache *redis.Client

func TestMain(m *testing.M) {

	conf, err := config.LoadConfig("../..")
	if err != nil {
		log.Fatal("load config err: ", err)
	}

	testCache = InitRedis(&conf.Redis)

	os.Exit(m.Run())
}

type User struct {
	Id   uint32
	Age  uint32
	Name string
}

func (u *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func TestHSet(t *testing.T) {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("user:%d", i)
			err := testCache.HMSet(context.Background(), key, User{Id: uint32(i), Age: uint32(i) * 2, Name: key}).Err()
			if err != nil {
				t.Logf("HSET ERROR: %v\n", err)
			}

		}(i)
	}

	wg.Wait()
}

func TestHGet(t *testing.T) {
	var user User
	err := testCache.HGetAll(context.Background(), "user:1").Scan(&user)
	if err != nil {
		t.Logf("HGETALL ERROR: %v", err)
		return
	}

	t.Logf("user: %+v\n", user)
}

func TestHScan(t *testing.T) {
	var user User
	keys, cursor, err := testCache.HScan(context.Background(), "user:1", uint64(1), "", 0).Result()
	if err != nil {
		t.Logf("HGETALL ERROR: %v", err)
		return
	}

	t.Logf("keys: %v\n", keys)
	t.Logf("cursor: %d\n", cursor)
	t.Logf("user: %+v\n", user)
}
