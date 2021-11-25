package util

import (
	"errors"
	"testing"
)

func TestNewErrorResponse(t *testing.T) {
	err := errors.New("test error")

	res := NewErrorResponse(err.Error())

	if res.Error != err.Error() {
		t.Errorf("expected: %s, got: %s", err.Error(), res.Error)
	}
}
