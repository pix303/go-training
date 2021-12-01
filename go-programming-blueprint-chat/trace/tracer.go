package trace

import (
	"fmt"
	"io"
)

type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

func New(w io.Writer) Tracer {
	candidateTracer := tracer{out: w}
	return &candidateTracer
}

type nilTracer struct {
}

func (t *nilTracer) Trace(a ...interface{}) {}
func Off() Tracer {
	return &nilTracer{}
}
