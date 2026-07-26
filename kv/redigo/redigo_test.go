package redigo

import "github.com/hanzokv/lock/v4/kv"

var _ kv.Conn = (*conn)(nil)

var _ kv.Pool = (*pool)(nil)
