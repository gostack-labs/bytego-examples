package xresult

import (
	"errors"
	"fmt"
	"testing"
)

func Test_result_Cause(t *testing.T) {
	err := errors.New("error 0")
	err = fmt.Errorf("error1: %w", err)
	err = fmt.Errorf("error2: %w", err)
	err = fmt.Errorf("error3: %w", err)
	r := Error(501, err)
	if cause := r.Cause(); errors.Is(cause, err) {
		t.Errorf("Cause() error %v", cause)
	}
}
