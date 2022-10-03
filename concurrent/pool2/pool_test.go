package pool2

import (
	"log"
	"sync"
	"testing"
	"time"
)

func TestPool_AddTask(t *testing.T) {
	pool := NewPool(4)
	doneNum := 0
	lock := sync.Mutex{}
	taskNum := 10
	for i := 0; i < taskNum; i++ {
		pool.AddTask(func() {
			id := i // actually id here is not correct!
			log.Printf("%d start working", id)
			time.Sleep(time.Second)
			log.Printf("%d end working", id)
			lock.Lock()
			doneNum++
			lock.Unlock()
		})
	}
	pool.WaitAllDone()
	if doneNum != taskNum {
		t.Errorf("doneNum %d != taskNum %d", doneNum, taskNum)
	}
}
