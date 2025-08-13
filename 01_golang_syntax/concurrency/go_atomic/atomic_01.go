package main

import "sync/atomic"

var n1 int64

func addSyncByAtomic(delta int64) int64 {
	return atomic.AddInt64(&n1, delta)
}
func readSyncByAtomic() int64 {
	return atomic.LoadInt64(&n1)
}
