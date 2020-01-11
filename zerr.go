package zerr

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// Wrap returns a wrapped error with a stack trace at the point Wrap is called, and the supplied message.
// If err is nil, Wrap returns nil.
func Wrap(err error, msg ...string) error {
	if err == nil {
		return nil
	}

	format := strings.Builder{}

	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()

	format.WriteString(file[strings.LastIndex(file, "/")+1:] + "/")
	format.WriteString(name[strings.LastIndex(name, "/")+1:] + "():")
	format.WriteString(strconv.Itoa(line) + ": ")

	for _, t := range msg {
		format.WriteString(t + ": ")
	}
	format.WriteString("%w")

	return fmt.Errorf(format.String(), err)
}
