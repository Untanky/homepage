package errors

import (
	"errors"
	"fmt"
	"net/http"
)

type httpError struct {
	parent   error
	httpCode int
}

func (err httpError) Error() string {
	return err.parent.Error()
}

func (err httpError) Unwrap() error {
	return err.parent
}

func (err httpError) Render(writer http.ResponseWriter) {
	writer.WriteHeader(err.httpCode)
	fmt.Fprintln(writer, err.parent.Error())
}

func ResponseError(code int, err error) error {
	return httpError{
		parent:   err,
		httpCode: code,
	}
}

func InternalServerError(err error) error {
	return ResponseError(http.StatusInternalServerError, err)
}

func NotFoundError(err error) error {
	return ResponseError(http.StatusNotFound, err)
}

func RenderError(writer http.ResponseWriter, err error) {
	httpErr, ok := errors.AsType[httpError](err)
	if !ok {
		renderUnknownError(writer)
	}

	httpErr.Render(writer)
}

func renderUnknownError(writer http.ResponseWriter) {

}
