package rueidis

import "github.com/hanzokv/lock/v4/redis"

var _ redis.Conn = (*conn)(nil)

var _ redis.Pool = (*pool)(nil)
