package main

type WaitGroup struct {
	count  int
	sem    chan struct{}
	waitCh chan struct{}
}

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{
		sem:    make(chan struct{}, 1),
		waitCh: make(chan struct{}),
	}
}

func (wg *WaitGroup) Add(delta int) {
	wg.sem <- struct{}{}
	defer func() { <-wg.sem }()

	wg.count += delta
	if wg.count < 0 {
		panic("negative counter")
	}

	if wg.count == 0 {
		close(wg.waitCh)
		wg.waitCh = make(chan struct{})
	}
}

func (wg *WaitGroup) Done() {
	wg.Add(-1)
}

func (wg *WaitGroup) Wait() {
	for {
		wg.sem <- struct{}{}
		if wg.count == 0 {
			<-wg.sem
			return
		}
		ch := wg.waitCh
		<-wg.sem

		<-ch
	}
}
