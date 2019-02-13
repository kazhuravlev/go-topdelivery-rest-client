package topdeliveryRestClient

import "fmt"

type ErrBadRequest struct {
	err     error
	message string
}

func (e ErrBadRequest) Error() string {
	return fmt.Sprintf("bad request: %s (%s)", e.err.Error(), e.message)
}

type ErrBadResponse struct {
	err     error
	message string
}

func (e ErrBadResponse) Error() string {
	return fmt.Sprintf("bad response: %s (%s)", e.err.Error(), e.message)
}
