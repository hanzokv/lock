package redsync

import (
	"context"
	"os"
	"strconv"
	"testing"

	kvlib "github.com/hanzokv/go/v9"
	"github.com/hanzokv/lock/v4/kv"
	"github.com/hanzokv/lock/v4/kv/kvgo"
	"github.com/stvp/tempredis"
)

var servers []*tempredis.Server

type testCase struct {
	poolCount int
	pools     []kv.Pool
}

func makeCases(poolCount int) map[string]*testCase {
	return map[string]*testCase{
		"kvgo": {
			poolCount,
			newMockPoolsKVGo(poolCount),
		},
	}
}

// Maintain separate blocks of servers for each type of driver
const (
	ServerPools    = 6
	ServerPoolSize = 8
	KVGoBlock = 0
)

func TestMain(m *testing.M) {
	os.Exit(testRunner(m))
}

func testRunner(m *testing.M) int {
	for i := range ServerPoolSize * ServerPools {
		server, err := tempredis.Start(tempredis.Config{
			"port": strconv.Itoa(51200 + i),
		})
		if err != nil {
			panic(err)
		}
		defer server.Term()

		servers = append(servers, server)
	}

	return m.Run()
}

func TestRedsync(t *testing.T) {
	ctx := context.Background()
	for k, v := range makeCases(8) {
		t.Run(k, func(t *testing.T) {
			rs := New(v.pools...)

			mutex := rs.NewMutex("test-redsync")
			_ = mutex.Lock()

			assertAcquired(ctx, t, v.pools, mutex)
		})
	}
}





func newMockPoolsKVGo(n int) []kv.Pool {
	pools := make([]kv.Pool, n)

	offset := KVGoBlock * ServerPoolSize

	for i := range n {
		client := kvlib.NewClient(&kvlib.Options{
			Network: "unix",
			Addr:    servers[i+offset].Socket(),
		})
		pools[i] = kvgo.NewPool(client)
	}
	return pools
}

