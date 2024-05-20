package calculator

import (
	"log"
	"sync"
)

type BasicEquation struct {
	x  int64
	y  int64
	op string
}

type Queue[T any] struct {
	queue []T
	Front func() T
	Push  func(T)
	Empty func() bool
}

type Task func() (any, error)

type Worker struct {
	readerChan <-chan Task
	resultChan chan<- any
	errChan    chan<- error
}

func (w *Worker) Run() {
	w.readerChan = make(<-chan Task, 1)
	w.resultChan = make(chan<- any, 1)
	w.errChan = make(chan<- error, 1)

	go func() {
		for task := range w.readerChan {
			// TODO recover gracefully
			// TODO maybe add quit signal to terminate the routine
			if result, err := task(); err != nil {
				w.errChan <- err
			} else {
				w.resultChan <- result
			}
		}
	}()
}

type WorkerPool struct {
	sync.Mutex
	numberOfWorkers int64
	pool            Queue[Worker]
	busy            Queue[Worker]
	tasks           Queue[Task]
	DoTask          Task
}

var workers = make([]func(read chan BasicEquation) chan int64, 0)
var readers = make([]chan BasicEquation, 0)
var writers = make([]chan int64, 0)

const (
	numberOfWorkers = 4
)

func init() {
	for i := range numberOfWorkers {
		log.Println(i)
		workers = append(workers, func(read chan BasicEquation) chan int64 {
			writeChan := make(chan int64, 1)
			go func() {
				for equ := range read {
					writeChan <- equ.x * equ.y
				}
			}()

			return writeChan
		})

		readers = append(readers, make(chan BasicEquation))
		writers = append(writers, workers[i](readers[i]))
	}
}

func Calc(equation string, numberOfWorkers int64) int64 {
	/*
		(a + b) * (c + d)
		a + b / x
	*/

	readers[0] <- BasicEquation{x: 2, y: 5}
	result := <-writers[0]

	log.Println("result", result)
	return result
}
