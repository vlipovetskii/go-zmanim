package helper

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

func TraceStack() {
	for i := 1; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			panic(fmt.Sprintf("runtime.Caller(%d) failed", i))
		}
		if strings.Contains(file, "go/src/runtime") {
			break
		}

		if s := strings.Split(file, "/src/"); len(s) == 2 {
			log.Printf("%d : %s\n", line, s[1])
		} else if s := strings.Split(file, "/go/"); len(s) == 2 {
			log.Printf("%d : %s\n", line, s[1])
		} else {
			log.Printf("%d : %s\n", line, file)
		}
	}
}

/*
CurrentFuncName
How to get the current function name https://stackoverflow.com/questions/25927660/how-to-get-the-current-function-name/46289376#46289376
*/
func CurrentFuncName() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	// fmt.Printf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
	/*
		PRB: frame.Function returns github.com/username/.../src/package.FuncName
		WO: strip path and return package.FuncName only
	*/
	s := strings.Split(frame.Function, "/src/")
	return s[len(s)-1]
}
