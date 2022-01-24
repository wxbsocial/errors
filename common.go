package errors

import "errors"

var (
	ErrIdempotent = errors.New("idempotent failed")
)

func IsIdempotentError(err error) bool {
	return errors.Is(err, ErrIdempotent)
}
