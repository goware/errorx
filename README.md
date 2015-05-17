# errorx
Feature rich Golang error interface implementation inspired by Postgres error message style guide http://www.postgresql.org/docs/devel/static/error-style-guide.html

# features
* Everything Golang `error` has
* Everything Golang `errors` package provides
* 3 levels of error reporting: Info, Verbose, Debug, each providing more information
* **Error line and file** information in Debug. Not 100% accurate - shows file/line where errorx is rendered to string/JSON, but still quite helpful
* Formatted errors with parameters
* JSON errors
* Error codes

# docs
http://godoc.org/github.com/c2h5oh/errorx
