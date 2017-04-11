package log

import (
	"fmt"
)

type Log struct {
}

func New() *Log {
	return &Log{}
}

func (l *Log) Log(line interface{}) {
	fmt.Println(line)
}
