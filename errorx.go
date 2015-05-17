package errorx

import (
	"encoding/json"
	"fmt"
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
	errorCode    int    `json:"error_code"`
	errorMsg     string `json:"error_msg"`
	errorDetails string `json:"error_details,omitempty"`
	errorHint    string `json:"error_hint,omitempty"`
}

// New returns an error with error code and error messages provided in
// function params
func New(code int, errorMsg ...string) *Errorx {
	e := Errorx{errorCode: code}

	msgCount := len(errorMsg)
	if msgCount > 0 {
		e.errorMsg = errorMsg[0]
	}
	if msgCount > 1 {
		e.errorDetails = errorMsg[1]
	}
	if msgCount > 2 {
		e.errorHint = errorMsg[2]
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
	return e.errorCode
}

// Error returns a string representation of errorx. It includes at least
// error code and message. Error details and hint are provided depending
// on verbosity level set
func (e Errorx) Error() string {
	if verbosity == 0 || (e.errorDetails == "" && e.errorHint == "") {
		return fmt.Sprintf("Error %d: %s", e.errorCode, e.errorMsg)
	}
	if verbosity == 1 || e.errorHint == "" {
		return fmt.Sprintf("Error %d: %s - %s", e.errorCode, e.errorMsg, e.errorDetails)
	}
	return fmt.Sprintf("Error %d: %s - %s - %s", e.errorCode, e.errorMsg, e.errorDetails, e.errorHint)

}

// Errorf is a variant of Error that formats according to errorMsg
// speficier and returns resulting string error details and hint
// will not be formated
func (e Errorx) Errorf(params ...interface{}) string {
	if verbosity == 0 || (e.errorDetails == "" && e.errorHint == "") {
		return fmt.Sprintf("Error %d: %s", e.errorCode, fmt.Sprintf(e.errorMsg, params...))
	}
	if verbosity == 1 || e.errorHint == "" {
		return fmt.Sprintf("Error %d: %s - %s", e.errorCode, fmt.Sprintf(e.errorMsg, params...), e.errorDetails)
	}
	return fmt.Sprintf("Error %d: %s - %s - %s", e.errorCode, fmt.Sprintf(e.errorMsg, params...), e.errorDetails, e.errorHint)
}

// Json returns a json representation (as []byte) of errorx and error
// if marshaling fails
func (e Errorx) Json() ([]byte, error) {
	err := Errorx{errorCode: e.errorCode, errorMsg: e.errorMsg}
	if verbosity > 0 && e.errorDetails != "" {
		err.errorDetails = e.errorDetails
	}
	if verbosity > 1 && e.errorHint != "" {
		err.errorHint = e.errorHint
	}
	return json.Marshal(err)
}

// Jsonf is a variant of Json that formats according to errorMsg
// speficier and returns resulting string error details and hint
// will not be formated
func (e Errorx) Jsonf(params ...interface{}) ([]byte, error) {
	err := Errorx{errorCode: e.errorCode, errorMsg: fmt.Sprintf(e.errorMsg, params...)}
	if verbosity > 0 && e.errorDetails != "" {
		err.errorDetails = e.errorDetails
	}
	if verbosity > 1 && e.errorHint != "" {
		err.errorHint = e.errorHint
	}
	return json.Marshal(err)
}
