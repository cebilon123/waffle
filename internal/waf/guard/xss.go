package guard

import (
	"errors"
	"io"

	"waffle/internal/request"
	"waffle/internal/xss"
)

type XSS struct {
}

// Validate validates if given input is XSS. It only returns error
// if given input is XSS, in other cases it returns nil.
func (X *XSS) Validate(rw *request.Wrapper) error {
	body, err := io.ReadAll(rw.Request().Body)
	if err != nil {
		return nil
	}

	if xss.IsXSS(string(body)) {
		return errors.New("body is xss")
	}

	return nil
}
