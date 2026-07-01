package errors

import (
	"fmt"
)

func ErrMaxSizeExceed(algorithm string, size int) error {
	return fmt.Errorf("highest accepted size for %s is %d", algorithm, size)
}