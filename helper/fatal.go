package helper

import (
	"fmt"
	"log"
)

func Panic(s string) {
	log.Println(s)
	TraceStack()
	panic("")
}

func PanicOnError(err error) {
	Panic(fmt.Sprintf("%v", err))
}
