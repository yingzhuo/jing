/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *      _ _
 *     | (_)_ __   __ _
 *  _  | | | '_ \ / _` |
 * | |_| | | | | | (_| |
 *  \___/|_|_| |_|\__, |                        https://github.com/yingzhuo/jing
 *                |___/
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package goroutine

import "sync"

type Semaphore struct {
	threads chan int
	wg      *sync.WaitGroup
}

func NewSemaphore(n int) *Semaphore {
	inst := new(Semaphore)
	inst.threads = make(chan int, n)
	inst.wg = &sync.WaitGroup{}
	return inst
}

func (s *Semaphore) Add() {
	s.threads <- 1
	s.wg.Add(1)
}

func (s *Semaphore) Done() {
	s.wg.Done()
	<-s.threads
}

func (s *Semaphore) Wait() {
	s.wg.Wait()
}

func (s *Semaphore) Close() {
	close(s.threads)
}
