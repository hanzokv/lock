package redsync

import (
	"context"
	"testing"

	kvlib "github.com/hanzokv/go/v9"
	"github.com/hanzokv/lock/v4/kv"
	"github.com/hanzokv/lock/v4/kv/kvgo"
	"github.com/stvp/tempredis"
)

func TestScripts_Hash(t *testing.T) {
	testScriptHash(t, "deleteScript", deleteScript)
	testScriptHash(t, "touchWithSetNXScript", touchWithSetNXScript)
	testScriptHash(t, "touchScript", touchScript)
}

func testScriptHash(t *testing.T, name string, script *kv.Script) {
	t.Run(name, func(t *testing.T) {
		server, err := tempredis.Start(tempredis.Config{})
		if err != nil {
			t.Fatal(err)
		}
		defer server.Term()

		client := kvlib.NewClient(&kvlib.Options{
			Network: "unix",
			Addr:    server.Socket(),
		})

		pool := kvgo.NewPool(client)
		conn, err := pool.Get(context.Background())
		if err != nil {
			t.Fatal(err)
		}

		dupscript := kv.NewScript(script.KeyCount, script.Src, "")
		err = conn.ScriptLoad(dupscript)
		if err != nil {
			t.Fatal(err)
		}

		if dupscript.Hash == "" {
			t.Fatal("script hash is empty")
		}
		if dupscript.Hash != script.Hash {
			t.Fatalf("script hash mismatch; want %s, got %s", script.Hash, dupscript.Hash)
		}
	})
}
