package crabgo

import (
	"log"
	"os"
)

var CrabLogger = log.New(os.Stdout, "", log.Ldate|log.Ltime)

func Dump(i interface{}) {
	CrabLogger.Printf("[D] %v \n", i)
}
