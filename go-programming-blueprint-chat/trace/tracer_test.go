package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {

	var buffer bytes.Buffer
	tracer := New(&buffer)

	if tracer == nil {
		t.Error("Return from New should be not null")
	} else {
		tracer.Trace("Hello trace package.")
		if buffer.String() != "Hello trace package.\n" {
			t.Errorf("Trace should not write '%s'", buffer.String())
		}
	}
}

func testOff(t *testing.T) {
	var silentTraceer Tracer = Off()
	silentTraceer.Trace("fake msg")
}
