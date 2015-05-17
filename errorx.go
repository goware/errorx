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

var verbosity uint8

type Errorx struct {
	errorCode    int    `json:"error_code"`
	errorMsg     string `json:"error_msg"`
	errorDetails string `json:"error_details,omitempty"`
	errorHint    string `json:"error_hint,omitempty"`
}

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

func SetVerbosity(v uint8) {
	if v > 3 {
		v = 2
	}
	verbosity = v
}

func (e Errorx) Code() int {
	return e.errorCode
}

func (e Errorx) Error() string {
	if verbosity == 0 || (e.errorDetails == "" && e.errorHint == "") {
		return fmt.Sprintf("Error %d: %s", e.errorCode, e.errorMsg)
	}
	if verbosity == 1 || e.errorHint == "" {
		return fmt.Sprintf("Error %d: %s - %s", e.errorCode, e.errorMsg, e.errorDetails)
	}
	return fmt.Sprintf("Error %d: %s - %s - %s", e.errorCode, e.errorMsg, e.errorDetails, e.errorHint)

}

func (e Errorx) Errorf(params ...interface{}) string {
	if verbosity == 0 || (e.errorDetails == "" && e.errorHint == "") {
		return fmt.Sprintf("Error %d: %s", e.errorCode, fmt.Sprintf(e.errorMsg, params...))
	}
	if verbosity == 1 || e.errorHint == "" {
		return fmt.Sprintf("Error %d: %s - %s", e.errorCode, fmt.Sprintf(e.errorMsg, params...), e.errorDetails)
	}
	return fmt.Sprintf("Error %d: %s - %s - %s", e.errorCode, fmt.Sprintf(e.errorMsg, params...), e.errorDetails, e.errorHint)
}

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
