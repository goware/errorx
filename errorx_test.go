package errorx_test

import (
	"errors"
	"testing"

	"github.com/c2h5oh/errorx"
)

func TestErrorVerbosity(t *testing.T) {
	e := errorx.New(10, "error message", "error details", "error hint")

	errorx.SetVerbosity(errorx.Info)
	if e.Error() != "error 10: error message" {
		t.Errorf("Expected 'error 10: error message', got '%s'", e.Error())
	}

	errorx.SetVerbosity(errorx.Verbose)
	if e.Error() != "error 10: error message | error details" {
		t.Errorf("Expected 'error 10: error message | error details', got '%s'", e.Error())
	}

	errorx.SetVerbosity(errorx.Debug)
	err := e.Error()
	if err != "errorx_test.go:24: error 10: error message | error details; error hint" {
		t.Errorf("Expected 'errorx_test.go:23: error 10: error message | error details; error hint', got '%s'", err)
	}

	errorx.SetVerbosity(errorx.Trace)
	err = e.Error()
	if err != "errorx_test.go:24: error 10: error message | error details; error hint" {
		t.Errorf("Expected 'errorx_test.go:23: error 10: error message | error details; error hint', got '%s'", err)
	}
}

/*
func TestErrorfVerbosity(t *testing.T) {
	e := errorx.New(11, "could not find '%s'", "error details", "error hint")

	errorx.SetVerbosity(errorx.Info)
	if e.Errorf("myFile.sh") != "error 11: could not find 'myFile.sh'" {
		t.Errorf("Expected 'error 11: could not find 'myFile.sh'', got '%s'", e.Errorf("myFile.sh"))
	}

	errorx.SetVerbosity(errorx.Verbose)
	if e.Errorf("myFile.sh") != "error 11: could not find 'myFile.sh' - error details" {
		t.Errorf("Expected 'error 11: could not find 'myFile.sh' - error details', got '%s'", e.Errorf("myFile.sh"))
	}

	errorx.SetVerbosity(errorx.Debug)
	err := e.Errorf("myFile.sh")
	if err != "errorx_test.go:43: error 11: could not find 'myFile.sh' - error details - error hint" {
		t.Errorf("Expected 'errorx_test.go:43: error 11: could not find 'myFile.sh' - error details - error hint', got '%s'", err)
	}
}*/

func TestJsonVerbosity(t *testing.T) {
	e := errorx.New(12, "error message", "error details", "error hint")

	errorx.SetVerbosity(errorx.Info)
	err, _ := e.Json()
	if string(err) != `{"error_code":12,"error_message":"error message"}` {
		t.Errorf(`Expected '{"error_code":12,"error_messsage":"error message"}', got '%s'`, string(err))
	}

	errorx.SetVerbosity(errorx.Verbose)
	err, _ = e.Json()
	if string(err) != `{"error_code":12,"error_message":"error message","error_details":["error details"]}` {
		t.Errorf(`Expected '{"error_code":12,"error_message":"error message","error_details":["error details"]}', got '%s'`, string(err))
	}

	errorx.SetVerbosity(errorx.Debug)
	err, _ = e.Json()
	if string(err) != `{"error_code":12,"error_message":"error message","error_details":["error details","error hint"],"stack":[{"file":"errorx_test.go","line":67,"function":"github.com/c2h5oh/errorx_test.TestJsonVerbosity"}]}` {
		t.Errorf(`Expected '{"error_code":12,"error_message":"error message","error_details":["error details","error hint"],"stack":[{"file":"errorx_test.go","line":67,"function":"github.com/c2h5oh/errorx_test.TestJsonVerbosity"}]}', got '%s'`, string(err))
	}
}

/*
func TestJsonfVerbosity(t *testing.T) {
	e := errorx.New(13, "could not find '%s'", "error details", "error hint")

	errorx.SetVerbosity(errorx.Info)
	err, _ := e.Jsonf("myFile.sh")
	if string(err) != `{"error_code":13,"error_msg":"could not find 'myFile.sh'"}` {
		t.Errorf(`Expected '{"error_code":13,"error_msg":"could not find 'myFile.sh'"}', got '%s'`, string(err))
	}

	errorx.SetVerbosity(errorx.Verbose)
	err, _ = e.Jsonf("myFile.sh")
	if string(err) != `{"error_code":13,"error_msg":"could not find 'myFile.sh'","error_details":"error details"}` {
		t.Errorf(`Expected '{"error_code":13,"error_msg":"could not find 'myFile.sh'","error_details":"error details"}', got '%s'`, string(err))
	}

	errorx.SetVerbosity(errorx.Debug)
	err, _ = e.Jsonf("myFile.sh")
	if string(err) != `{"error_code":13,"error_msg":"could not find 'myFile.sh'","error_details":"error details","error_hint":"error hint","file":"errorx_test.go","line":87}` {
		t.Errorf(`Expected '{"error_code":13,"error_msg":"could not find 'myFile.sh'","error_details":"error details","error_hint":"error hint","file":"errorx_test.go","line":87}}', got '%s'`, string(err))
	}
}*/

func TestErrorCode(t *testing.T) {
	e := errorx.New(14, "error message", "error details", "error hint")

	if e.ErrorCode() != 14 {
		t.Errorf(`Invalide error code - expected 14, got %d`, e.ErrorCode())
	}
}

func TestErrorEmbedding(t *testing.T) {
	wrappableErrorx := errorx.New(2, "wrapped error message", "wrapped error details", "wrapped error hint")
	wrappableError := errors.New("wrapped regular error")
	e1 := errorx.New(10, "error message", "error details", "error hint")
	e1.Wrap(wrappableErrorx)
	e2 := errorx.New(10, "error message", "error details", "error hint")
	e2.Wrap(wrappableError)

	errorx.SetVerbosity(errorx.Info)
	if e1.Error() != "error 10: error message" {
		t.Errorf("Expected 'error 10: error message', got '%s'", e1.Error())
	}
	if e2.Error() != "error 10: error message" {
		t.Errorf("Expected 'error 10: error message', got '%s'", e2.Error())
	}

	errorx.SetVerbosity(errorx.Verbose)
	if e1.Error() != "error 10: error message | error details\ncause: error 2: wrapped error message | wrapped error details" {
		t.Errorf("error 10: error message | error details\ncause: error 2: wrapped error message | wrapped error details', got '%s'", e1.Error())
	}

	if e2.Error() != "error 10: error message | error details\ncause: wrapped regular error" {
		t.Errorf("error 10: error message | error details\ncause: wrapped regular error', got '%s'", e2.Error())
	}

	errorx.SetVerbosity(errorx.Debug)
	err := e1.Error()
	if err != "errorx_test.go:23: error 10: error message | error details; error hint" {
		t.Errorf("Expected 'errorx_test.go:23: error 10: error message | error details; error hint', got '%s'", err)
	}
}
