package errorx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	Info = iota
	Verbose
	Debug
	Trace
)

// verbosity variable stores global verbosity setting for errorx package
// based on it's value different level of error details will be provided
// by Error, Errorf, Json and Jsonf
var verbosity int

// Errorx is a more feature rich implementation of error interface inspired
// by PostgreSQL error style guide
type Errorx struct {
	Code    int      `json:"error_code"`
	Message string   `json:"error_message"`
	Details []string `json:"error_details,omitempty"`
	Cause   error    `json:"cause,omitempty"`
	Stack   Stack    `json:"stack,omitempty"`
}

type StackFrame struct {
	File     string `json:"file,omitempty"`
	Line     int    `json:"line,omitempty"`
	Function string `json:"function,omitempty"`
}

type Stack []StackFrame

func (s Stack) String() string {
	var buffer bytes.Buffer
	for i := 0; i < len(s); i++ {
		buffer.WriteString(fmt.Sprintf("%s:%d %s", s[i].File, s[i].Line, s[i].Function))
	}
	return buffer.String()
}

// New returns an error with error code and error messages provided in
// function params
func New(code int, ErrorMsg ...string) *Errorx {
	e := Errorx{Code: code}

	msgCount := len(ErrorMsg)
	if msgCount > 0 {
		e.Message = ErrorMsg[0]
	}
	if msgCount > 1 {
		e.Details = ErrorMsg[1:]
	}

	return &e
}

// Wrap
func (e Errorx) Wrap(err error) {
	e.Cause = err
}

// SetVerbosity changes global verbosity setting
func SetVerbosity(v int) {
	verbosity = v
}

// ErrorCode returns Errorx error code value. It's intended primarily to allow
// easy error comparison / matching
func (e Errorx) ErrorCode() int {
	return e.Code
}

// Error returns a string representation of errorx. It includes at least
// error code and message. Error details and hint are provided depending
// on verbosity level set
func (e Errorx) Error() string {
	maxMsg := len(e.Details)
	if maxMsg > verbosity {
		maxMsg = verbosity
	}

	switch verbosity {
	case 0:
		return fmt.Sprintf("error %d: %s", e.Code, e.Message)
	case 1:
		if e.Cause == nil {
			return fmt.Sprintf("error %d: %s | %s", e.Code, e.Message, strings.Join(e.Details[0:maxMsg], "; "))
		}
		return fmt.Sprintf("error %d: %s | %s\ncause: %s", e.Code, e.Message, strings.Join(e.Details[0:maxMsg], "; "), e.Cause.Error())
	case 2:
		e.getTrace()

		if e.Cause == nil {
			return fmt.Sprintf("%s:%d: error %d: %s | %s", e.Stack[0].File, e.Stack[0].Line, e.Code, e.Message, strings.Join(e.Details[0:maxMsg], "; "))
		}
		return fmt.Sprintf("%s:%d: error %d: %s | %s\ncause: %s", e.Stack[0].File, e.Stack[0].Line, e.Code, e.Message, strings.Join(e.Details[0:maxMsg], "; "), e.Cause.Error())
	default:
		e.getTrace()

		if e.Cause == nil {
			return fmt.Sprintf("%s:%d: error %d: %s | %s\n%s", e.Stack[0].File, e.Stack[0].Line, e.Code, e.Message, strings.Join(e.Details[0:maxMsg], "; "), e.Stack.String())
		}
		return fmt.Sprintf("%s:%d: error %d: %s | %s\ncause: %s\n%s", e.Stack[0].File, e.Stack[0].Line, e.Code, e.Message, strings.Join(e.Details[0:maxMsg], "; "), e.Cause.Error(), e.Stack.String())
	}
}

// Errorf is a variant of Error that formats according to ErrorMsg
// speficier and returns resulting string. Error details and hint
// will not be formated
/*
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
}*/

// Json returns a json representation (as []byte) of errorx and error
// if marshaling fails
func (e Errorx) Json() ([]byte, error) {
	e.getTrace()
	err := e.verbositySubset()

	return json.Marshal(err)
}

// Jsonf is a variant of Json that formats according to ErrorMsg
// speficier and returns resulting string. Error details and hint
// will not be formated
/*
func (e Errorx) Jsonf(params ...interface{}) ([]byte, error) {
	e.getTrace()
	err := e.verbositySubset()

	return json.Marshal(err)
}*/

func (e Errorx) verbositySubset() Errorx {
	err := Errorx{Code: e.Code, Message: e.Message}
	maxMsg := len(e.Details)
	if maxMsg > verbosity {
		maxMsg = verbosity
	}

	if verbosity > 0 {
		err.Details = e.Details[0:maxMsg]
	}
	if verbosity > 1 {
		if e.Cause != nil {
			if cause, ok := e.Cause.(Errorx); ok {
				err.Cause = cause.verbositySubset()
			} else {
				err.Cause = Errorx{Message: e.Cause.Error()}
			}
		}
	}
	return err
}

func (e Errorx) getTrace() {
	if verbosity < 2 {
		return
	}

	if verbosity == 2 {
		pc, fn, line, ok := runtime.Caller(2)
		if !ok {
			return
		}

		s := StackFrame{}
		s.Function = funcName(pc)
		s.Line = line
		_, s.File = filepath.Split(fn)

		e.Stack = []StackFrame{s}
		return
	}

	for i := 2; ; i++ {
		e.Stack = make([]StackFrame, 0)
		pc, fn, line, ok := runtime.Caller(i)
		if !ok {
			// no more frames - we're done
			break
		}

		f := StackFrame{File: fn, Line: line, Function: funcName(pc)}
		e.Stack = append(e.Stack, f)
	}
}

// funcName gets the name of the function at pointer or "??" if one can't be found
func funcName(pc uintptr) string {
	if f := runtime.FuncForPC(pc); f != nil {
		return f.Name()
	}
	return "??"
}
