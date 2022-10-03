package pool

import (
	"log"
	"testing"
	"time"
)

func TestPool_AddTask(t *testing.T) {
	pool := NewPool(2)
	for i := 0; i < 4; i++ {
		pool.AddTask(func() {
			id := i // actually id here is not correct!
			log.Printf("%d start working", id)
			time.Sleep(time.Second)
			log.Printf("%d end working", id)
		})
	}
}
