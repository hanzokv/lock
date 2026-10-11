module github.com/hanzokv/lock/v4

go 1.27.1

require (
	github.com/hanzokv/go/v9 v9.22.0
	github.com/stvp/tempredis v0.0.0-20181119212430-b82af8480203
	golang.org/x/sync v0.24.0
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/gomodule/redigo v1.9.3 // indirect
	github.com/stretchr/testify v1.11.1 // indirect
	go.uber.org/atomic v1.11.0 // indirect
)

replace github.com/stvp/tempredis => github.com/hjr265/tempredis v0.0.0-20231015061547-ad8aa5a343a2

retract v4.14.0
