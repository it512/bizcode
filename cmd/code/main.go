package main

import (
	"fmt"

	"github.com/it512/bizcode"
)

func main() {
	c := bizcode.By("0A")

	for {
		code := c()
		c29 := code[29]
		if c29 == '0' {
			fmt.Printf("code = %s, type = %s\n", code, bizcode.CodeType(code))
		}
	}
}
