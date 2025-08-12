package logger

import (
	"fmt"
	"sync"

	loggerinterface "github.com/alexey/firstApp/pkg/logger/interface"
)

type RegisterLog struct{}

func NewRegisteredLog() *RegisterLog {
	return &RegisterLog{}
}

var (
	mutex         sync.RWMutex
	regLoggerData = make(map[string]loggerinterface.Logger)
)

func (r *RegisterLog) RegisterLogger(key string, log loggerinterface.Logger) error {
	mutex.Lock()
	defer mutex.Unlock()

	if _, exist := regLoggerData[key]; exist {
		return fmt.Errorf("Запись c ключом %s уже существует ", key)
	}

	regLoggerData[key] = log
	fmt.Printf("Map after adding %v c ключем %s", regLoggerData, key)
	fmt.Printf("размер мапы %d \n", len(regLoggerData))
	return nil
}

func (r *RegisterLog) GetRegLogger(key string) (loggerinterface.Logger, error) {
	mutex.Lock()
	defer mutex.Unlock()
	l, exist := regLoggerData[key]
	if !exist {
		return nil, fmt.Errorf("Ключа %s нет в базе ", key)
	}
	return l, nil
}

func (r *RegisterLog) GetAllRegLogger() map[string]loggerinterface.Logger {
	mutex.Lock()
	defer mutex.Unlock()

	copyOfRegLog := make(map[string]loggerinterface.Logger, len(regLoggerData))
	for k, v := range regLoggerData {
		copyOfRegLog[k] = v
	}

	return copyOfRegLog
}
