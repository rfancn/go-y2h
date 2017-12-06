package y2h

import (
	//"io/ioutil"
	//"fmt"
)

type Y2HError struct {
	errMsg string
	errDetail error
}

func (e *Y2HError) Error() string {
	if e.errDetail != nil {
		return e.errMsg + "\n" + e.errDetail.Error()
	}
	return e.errMsg
}

func NewY2HError(errMsg string, errSlice ...error) *Y2HError {
	var errDetail error = nil
	// now only supports pass only one parent error here
	if len(errSlice) == 1 {
		errDetail = errSlice[0]
	}
	return  &Y2HError{errMsg, errDetail}
}