package coroutine

import (
	"log"
	"reflect"
	"runtime"
	"runtime/debug"
	"sync"
)

// Pool is a struct that manages a collection of workers, each with their own goroutine.
type Pool struct {
	job func(...interface{})
	ch  chan struct{}
	wg  sync.WaitGroup
	log *log.Logger
}

// NewPool creates a new Pool of workers that starts with n workers.
func NewPool(n int, logger *log.Logger, job func(...interface{})) *Pool {
	if n <= 0 {
		return nil
	}
	return &Pool{job: job, ch: make(chan struct{}, n), log: logger}
}

// Process will use the Pool to process job asynchronously
func (p *Pool) Process(param ...interface{}) {
	p.wg.Add(1)
	go func() {
		p.ch <- struct{}{}
		defer func() {
			p.wg.Done()
			<-p.ch
			if err := recover(); err != nil && p.log != nil {
				p.log.Printf("Panic: Func Name:%s, pool err recovered:%v, stack trace:%s",
					runtime.FuncForPC(reflect.ValueOf(p.job).Pointer()).Name(), err, string(debug.Stack()))
			}
		}()
		p.job(param...)
	}()
}

// Wait will wait until all workers finish their job.
func (p *Pool) Wait() {
	p.wg.Wait()
}
