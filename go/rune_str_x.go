package demo

import (
	"fmt"
)

func RuneStr(s string) {
	for _, r := range s {
		fmt.Print(string(r))
	}
}
