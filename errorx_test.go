package errorx_test

import (
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
	if e.Error() != "error 10: error message - error details" {
		t.Errorf("Expected 'error 10: error message - error details', got '%s'", e.Error())
	}

	errorx.SetVerbosity(errorx.Debug)
	if e.Error() != "error 10: error message - error details - error hint" {
		t.Errorf("Expected 'error 10: error message - error details - error hint', got '%s'", e.Error())
	}
}

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
	if e.Errorf("myFile.sh") != "error 11: could not find 'myFile.sh' - error details - error hint" {
		t.Errorf("Expected 'error 11: could not find 'myFile.sh' - error details - error hint', got '%s'", e.Errorf("myFile.sh"))
	}
}

func TestJsonVerbosity(t *testing.T) {
	e := errorx.New(12, "error message", "error details", "error hint")

	errorx.SetVerbosity(errorx.Info)
	err, _ := e.Json()
	if string(err) != `{"error_code":12,"error_msg":"error message"}` {
		t.Errorf(`Expected '{"error_code":12,"error_msg":"error message"}', got '%s'`, string(err))
	}

	errorx.SetVerbosity(errorx.Verbose)
	err, _ = e.Json()
	if string(err) != `{"error_code":12,"error_msg":"error message","error_details":"error details"}` {
		t.Errorf(`Expected '{"error_code":12,"error_msg":"error message","error_details":"error details"}', got '%s'`, string(err))
	}

	errorx.SetVerbosity(errorx.Debug)
	err, _ = e.Json()
	if string(err) != `{"error_code":12,"error_msg":"error message","error_details":"error details","error_hint":"error hint"}` {
		t.Errorf(`Expected '{"error_code":12,"error_msg":"error message","error_details":"error details","error_hint":"error hint"}}', got '%s'`, string(err))
	}
}

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
	if string(err) != `{"error_code":13,"error_msg":"could not find 'myFile.sh'","error_details":"error details","error_hint":"error hint"}` {
		t.Errorf(`Expected '{"error_code":13,"error_msg":"could not find 'myFile.sh'","error_details":"error details","error_hint":"error hint"}}', got '%s'`, string(err))
	}
}

func TestErrorCode(t *testing.T) {
	e := errorx.New(14, "error message", "error details", "error hint")

	if e.Code() != 14 {
		t.Errorf(`Invalide error code - expected 14, got %d`, e.Code())
	}
}
