package model

import "errors"

var (
	ErrorConflictingData = errors.New("data already exists")
	ErrorInternal        = errors.New("internal server error")
)
