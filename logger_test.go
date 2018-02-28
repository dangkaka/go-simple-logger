package log

import (
	"testing"
)

func TestFormatLog(t *testing.T) {
	output := formatLog("a", "b")
	expected := "%v %v "
	if output != expected {
		t.Error("Expected ", expected, ", got ", output)
	}
}
