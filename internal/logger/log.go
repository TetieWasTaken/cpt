package logger

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	"github.com/spf13/cobra"
)

var Verbose bool

func Vprintf(cmd *cobra.Command, format string, args ...any) error {
	if !Verbose {
		return nil
	}

	prefix := fmt.Sprintf("%s ", time.Now().Format("2006-01-02 15:04:05.000"))

	pc, file, line, ok := runtime.Caller(1)
	fn := runtime.FuncForPC(pc).Name()

	if ok {
		prefix = prefix + fmt.Sprintf("(%s:%d / %s) ", filepath.Base(file), line, fn)
	}

	_, err := fmt.Fprintf(cmd.ErrOrStderr(), prefix+format+"\n", args...)
	return err
}
