package errorx

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"
)

const (
	Info    uint8 = 0
	Verbose uint8 = 1
	Debug   uint8 = 2
)

// verbosity variable stores global verbosity setting for errorx package
// based on it's value different level of error details will be provided
// by Error, Errorf, Json and Jsonf
var verbosity uint8

// Errorx is a more feature rich implementation of error interface inspired
// by PostgreSQL error style guide
type Errorx struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMsg     string `json:"error_msg"`
	ErrorDetails string `json:"error_details,omitempty"`
	ErrorHint    string `json:"error_hint,omitempty"`
	File         string `json:"file,omitempty"`
	Line         int    `json:"line,omitempty"`
}

// New returns an error with error code and error messages provided in
// function params
func New(code int, ErrorMsg ...string) *Errorx {
	e := Errorx{ErrorCode: code}

	msgCount := len(ErrorMsg)
	if msgCount > 0 {
		e.ErrorMsg = ErrorMsg[0]
	}
	if msgCount > 1 {
		e.ErrorDetails = ErrorMsg[1]
	}
	if msgCount > 2 {
		e.ErrorHint = ErrorMsg[2]
	}

	return &e
}

// SetVerbosity changes global verbosity setting
func SetVerbosity(v uint8) {
	if v > 3 {
		v = 2
	}
	verbosity = v
}

// Code returns Errorx error code value. It's intended primarily to allow
// easy error comparison / matching
func (e Errorx) Code() int {
	return e.ErrorCode
}

// Error returns a string representation of errorx. It includes at least
// error code and message. Error details and hint are provided depending
// on verbosity level set
func (e Errorx) Error() string {
	if verbosity == 0 || (e.ErrorDetails == "" && e.ErrorHint == "") {
		return fmt.Sprintf("error %d: %s", e.ErrorCode, e.ErrorMsg)
	}
	if verbosity == 1 {
		return fmt.Sprintf("error %d: %s - %s", e.ErrorCode, e.ErrorMsg, e.ErrorDetails)
	}
	_, fn, line, _ := runtime.Caller(1)
	_, file := filepath.Split(fn)
	return fmt.Sprintf("%s:%d: error %d: %s - %s - %s", file, line, e.ErrorCode, e.ErrorMsg, e.ErrorDetails, e.ErrorHint)
}

// Errorf is a variant of Error that formats according to ErrorMsg
// speficier and returns resulting string. Error details and hint
// will not be formated
func (e Errorx) Errorf(params ...interface{}) string {
	if verbosity == 0 || (e.ErrorDetails == "" && e.ErrorHint == "") {
		return fmt.Sprintf("error %d: %s", e.ErrorCode, fmt.Sprintf(e.ErrorMsg, params...))
	}
	if verbosity == 1 {
		return fmt.Sprintf("error %d: %s - %s", e.ErrorCode, fmt.Sprintf(e.ErrorMsg, params...), e.ErrorDetails)
	}
	_, fn, line, _ := runtime.Caller(1)
	_, file := filepath.Split(fn)
	return fmt.Sprintf("%s:%d: error %d: %s - %s - %s", file, line, e.ErrorCode, fmt.Sprintf(e.ErrorMsg, params...), e.ErrorDetails, e.ErrorHint)
}

// Json returns a json representation (as []byte) of errorx and error
// if marshaling fails
func (e Errorx) Json() ([]byte, error) {
	err := Errorx{ErrorCode: e.ErrorCode, ErrorMsg: e.ErrorMsg}
	if verbosity > 0 && e.ErrorDetails != "" {
		err.ErrorDetails = e.ErrorDetails
	}
	if verbosity > 1 {
		err.ErrorHint = e.ErrorHint
		_, fn, line, _ := runtime.Caller(1)
		err.Line = line
		_, err.File = filepath.Split(fn)
	}
	return json.Marshal(&err)
}

// Jsonf is a variant of Json that formats according to ErrorMsg
// speficier and returns resulting string. Error details and hint
// will not be formated
func (e Errorx) Jsonf(params ...interface{}) ([]byte, error) {
	err := Errorx{ErrorCode: e.ErrorCode, ErrorMsg: fmt.Sprintf(e.ErrorMsg, params...)}
	if verbosity > 0 && e.ErrorDetails != "" {
		err.ErrorDetails = e.ErrorDetails
	}
	if verbosity > 1 {
		err.ErrorHint = e.ErrorHint
		_, fn, line, _ := runtime.Caller(1)
		err.Line = line
		_, err.File = filepath.Split(fn)
	}
	return json.Marshal(err)
}
