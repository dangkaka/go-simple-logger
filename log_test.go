package log

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func setup() {
	EnableFuncCallDepth(false)
	EnableColorful(false)
	SetLogFuncCallDepth(0)
}

func captureOutput(f func()) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	return string(out)
}

func TestFatal(t *testing.T) {
	setup()
	SetLevel(LevelFatal)
	output := captureOutput(func() {
		Fatal("test")
	})
	expected := "test"
	if !strings.Contains(output, expected) {
		t.Error("Expected got", expected, ", got ", output)
	}
}

func TestFatal_LowerLevel(t *testing.T) {
	setup()
	SetLevel(-1)
	output := captureOutput(func() {
		Error("test")
	})
	expected := ""
	if output != "" {
		t.Error("Expected ", expected, ", got ", output)
	}
}

func TestError(t *testing.T) {
	setup()
	SetLevel(LevelError)
	output := captureOutput(func() {
		Error("test")
	})
	expected := "test"
	if !strings.Contains(output, expected) {
		t.Error("Expected got", expected, ", got ", output)
	}
}

func TestError_LowerLevel(t *testing.T) {
	setup()
	SetLevel(-1)
	output := captureOutput(func() {
		Fatal("test")
	})
	expected := ""
	if output != "" {
		t.Error("Expected ", expected, ", got ", output)
	}
}

func TestWarning(t *testing.T) {
	setup()
	SetLevel(LevelWarning)
	output := captureOutput(func() {
		Warning("test")
	})
	expected := "test"
	if !strings.Contains(output, expected) {
		t.Error("Expected got", expected, ", got ", output)
	}
}

func TestWarning_LowerLevel(t *testing.T) {
	setup()
	SetLevel(-1)
	output := captureOutput(func() {
		Warning("test")
	})
	expected := ""
	if output != "" {
		t.Error("Expected ", expected, ", got ", output)
	}
}

func TestInfo(t *testing.T) {
	setup()
	SetLevel(LevelInfo)
	output := captureOutput(func() {
		Info("test")
	})
	expected := "test"
	if !strings.Contains(output, expected) {
		t.Error("Expected got", expected, ", got ", output)
	}
}

func TestInfo_LowerLevel(t *testing.T) {
	setup()
	SetLevel(-1)
	output := captureOutput(func() {
		Info("test")
	})
	expected := ""
	if output != "" {
		t.Error("Expected ", expected, ", got ", output)
	}
}

func TestDebug(t *testing.T) {
	setup()
	SetLevel(LevelDebug)
	output := captureOutput(func() {
		Debug("test")
	})
	expected := "test"
	if !strings.Contains(output, expected) {
		t.Error("Expected got", expected, ", got ", output)
	}
}

func TestDebug_LowerLevel(t *testing.T) {
	setup()
	SetLevel(-1)
	output := captureOutput(func() {
		Debug("test")
	})
	expected := ""
	if output != "" {
		t.Error("Expected ", expected, ", got ", output)
	}
}
