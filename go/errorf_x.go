package demo

import (
	"fmt"
)

func ErrorFormat() error {
	return fmt.Errorf("1+2 != %d", 4)
}
