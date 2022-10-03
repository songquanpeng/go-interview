package pool

// Reference: https://juejin.cn/post/7028780454257885214

type Pool struct {
	work  chan func()
	count chan struct{} // remain goroutines that can be created
}

func NewPool(size int) *Pool {
	return &Pool{
		work:  make(chan func()), // no buffer
		count: make(chan struct{}, size),
	}
}

func (p *Pool) worker(task func()) {
	defer func() {
		<-p.count
	}()
	for {
		task()
		task = <-p.work
	}
}

func (p *Pool) AddTask(task func()) bool {
	select {
	case p.work <- task:
	case p.count <- struct{}{}:
		go p.worker(task)
	}
	return true
}
