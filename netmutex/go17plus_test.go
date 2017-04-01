// +build go1.7

package netmutex

import (
	"testing"
	"time"
)

func BenchmarkMain(mb *testing.B) {

	nm, err := Open(10, time.Minute, addresses, nil)
	if err != nil {
		mb.Fatal(err)
	}
	defer nm.Close(1, time.Minute)

	retries := 10
	timeout := time.Minute
	ttl := time.Second
	key := "a"

	mb.Run("LockUnlock",
		func(b *testing.B) {

			lock := &Lock{}

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				nm.Lock(retries, timeout, lock, key, ttl)

				nm.Unlock(retries, timeout, lock)

			}
		})

	mb.Run("LockUpdateUnlock",
		func(b *testing.B) {
			lock := &Lock{}

			b.ResetTimer()

			for n := 0; n < b.N; n++ {
				nm.Lock(retries, timeout, lock, key, ttl)

				for i := 0; i < 8; i++ {
					err = nm.Update(retries, timeout, lock, ttl)
				}

				nm.Unlock(retries, timeout, lock)
			}
		})
}
