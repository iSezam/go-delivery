package errs

import "errors"

var (
	DBconnect = errors.New("DB connect errs")
	DBNoRows  = errors.New("no rows")
	NotFound  = errors.New("not found")
)
