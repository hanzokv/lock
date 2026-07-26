package kvgo

import redsynckv "github.com/hanzokv/lock/v4/kv"

var _ redsynckv.Conn = (*conn)(nil)

var _ redsynckv.Pool = (*pool)(nil)
