package main

import (
	"context"

	kvlib "github.com/hanzokv/go/v9"
	"github.com/hanzokv/lock/v4"
	"github.com/hanzokv/lock/v4/kv/kvgo"
	"github.com/stvp/tempredis"
)

func main() {
	server, err := tempredis.Start(tempredis.Config{})
	if err != nil {
		panic(err)
	}
	defer server.Term()

	client := kvlib.NewClient(&kvlib.Options{
		Network: "unix",
		Addr:    server.Socket(),
	})

	pool := kvgo.NewPool(client)

	rs := redsync.New(pool)

	mutex := rs.NewMutex("test-redsync")
	ctx := context.Background()

	if err := mutex.LockContext(ctx); err != nil {
		panic(err)
	}

	if _, err := mutex.UnlockContext(ctx); err != nil {
		panic(err)
	}
}
