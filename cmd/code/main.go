package main

import (
	"fmt"

	"github.com/it512/bizcode"
)

func main() {
	c := bizcode.By("0A")

	code := c()

	fmt.Printf("code = %s, type = %s\n", code, bizcode.CodeType(code))
}
