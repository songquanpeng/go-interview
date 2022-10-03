package pool2

import "sync"

// Reference: https://juejin.cn/post/7028780454257885214

type Pool struct {
	work              chan func()
	remain            chan struct{} // remain goroutines that can be created
	closing           bool
	closingLock       sync.RWMutex
	close             chan struct{}
	activeCounter     uint
	activeCounterLock sync.Mutex
}

func NewPool(size uint) *Pool {
	return &Pool{
		work:          make(chan func()), // no buffer
		remain:        make(chan struct{}, size),
		closing:       false,
		close:         make(chan struct{}),
		activeCounter: 0,
	}
}

func (p *Pool) worker(task func()) {
	defer func() { // this is useless
		<-p.remain
	}()
	for {
		task()
		p.closingLock.RLock()
		if p.closing {
			p.closingLock.RUnlock()
			break
		}
		p.closingLock.RUnlock()
		task = <-p.work
	}
	p.activeCounterLock.Lock()
	p.activeCounter -= 1
	if p.activeCounter == 0 {
		p.close <- struct{}{}
	}
	p.activeCounterLock.Unlock()
}

func (p *Pool) AddTask(task func()) bool {
	p.closingLock.RLock()
	if p.closing {
		p.closingLock.RUnlock()
		return false
	}
	p.closingLock.RUnlock()

	select {
	case p.work <- task:
	case p.remain <- struct{}{}:
		p.activeCounterLock.Lock()
		p.activeCounter += 1
		p.activeCounterLock.Unlock()
		go p.worker(task)
	}
	return true
}

func (p *Pool) WaitAllDone() {
	p.closingLock.Lock()
	p.closing = true
	p.closingLock.Unlock()
	<-p.close
}
