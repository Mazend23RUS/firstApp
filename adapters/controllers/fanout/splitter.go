package fanout

import (
	"context"
	"fmt"
	"sync"

	"github.com/alexey/firstApp/pkg/logger"
)

// Fan-Out реализация

type SplitterInterface[T any] interface {
	Split(input <-chan T) []<-chan T // Fan out Логика
	CloseOutputChanels()
	RegisteredOutChanel() <-chan T
	// WorkingPool(pool WorkingPoolInterface[T]) SplitterInterface[T]
}

type Splitter[T any] struct {
	input         <-chan T
	outputchenel  []chan T
	buffersize    int
	mt            sync.RWMutex
	logRegistered *logger.RegisterLog
}

func NewSplitter[T any](bufsize int, logRegistered *logger.RegisterLog) *Splitter[T] {
	return &Splitter[T]{
		outputchenel:  make([]chan T, 0),
		buffersize:    bufsize,
		logRegistered: logRegistered,
	}
}

func (s *Splitter[T]) RegisteredOutChanel() chan T {
	s.mt.Lock()
	defer s.mt.Unlock()

	ch := make(chan T, s.buffersize)
	s.outputchenel = append(s.outputchenel, ch)

	return ch
}

func (s *Splitter[T]) Start(input <-chan T) {
	go s.processMessages(input)
}

// отдельная горутина для обработки входящих сообщений
func (s *Splitter[T]) processMessages(input <-chan T) {
	defer s.CloseOutputChanels()

	for inchan := range input {

		s.mt.RLock()
		outputs := s.outputchenel
		regLoggers := s.logRegistered.GetAllRegLogger()

		s.mt.RUnlock()

		for _, out := range outputs {
			select {
			case out <- inchan:
			default:
				for _, log := range regLoggers {
					log.PrintError(context.Background(), "Splitter: Канал переполнен, сообщение потеряно", nil)
				}
			}
		}

		for name, log := range regLoggers {
			log.PrintInfo(context.Background(), fmt.Sprintf("Splitter: Отправка сообщения зареги-нному логгеру:  "+
				name+" Само сообщение:  %v", inchan))
		}
	}
}

func (s *Splitter[T]) CloseOutputChanels() {
	s.mt.Lock()
	defer s.mt.Unlock()

	for _, ch := range s.outputchenel {
		close(ch)
	}
}
