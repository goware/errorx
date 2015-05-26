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
* Everything Golang `error` has
* Everything Golang `errors` package provides
* JSON errors you can just write to your webhandler

# docs
http://godoc.org/github.com/c2h5oh/errorx
