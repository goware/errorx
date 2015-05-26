[![Build Status](https://travis-ci.org/c2h5oh/errorx.svg?branch=master)](https://travis-ci.org/c2h5oh/errorx)
[![GoDoc](https://godoc.org/github.com/c2h5oh/errorx?status.svg)](https://godoc.org/github.com/c2h5oh/errorx)

# errorx
Feature-rich Golang error interface implementation inspired by Postgres error message style guide http://www.postgresql.org/docs/devel/static/error-style-guide.html

# features
* Error codes
* Verbosity levels
* **File and line on which the error occures** (Debug+ verbosity level). Not 100% accurate, but close enough: shows file/line where errorx is rendered to string/JSON
* error Stack traces (on verbosity level Trace)
* Nested errors (both regular Golang `error` and `Errorx`)
* Everything Golang `error` has - it's a drop-in replacement, because it implements `error` interface
* Everything Golang `errors` package provides
* JSON errors you can just write to your webhandler

# docs
http://godoc.org/github.com/c2h5oh/errorx

# example output
### json, nested error, verbosity: Trace
```json
{
   "error_code":10,
   "error_message":"error message",
   "error_details":[
      "error details",
      "error hint"
   ],
   "cause":{
      "error_code":200,
      "error_message":"wrapped error message",
      "error_details":[
         "wrapped error details",
         "wrapped error hint"
      ]
   },
   "stack":[
      {
         "file":"errorx_test.go",
         "line":175,
         "function":"github.com/c2h5oh/errorx_test.TestJsonErrorEmbedding"
      },
      {
         "file":"testing.go",
         "line":447,
         "function":"testing.tRunner"
      },
      {
         "file":"asm_amd64.s",
         "line":2232,
         "function":"runtime.goexit"
      }
   ]
}
```

### string (via .Error()), verbosity: Debug
```
errorx_test.go:28: error 10: error message | error details; error hint
```
