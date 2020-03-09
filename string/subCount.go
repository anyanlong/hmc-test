package main

import (
	"fmt"
	"strings"
)

func main() {
	cont := "   \r\n{{{{{{{   }}}}"
	leftKHCount := strings.Count(cont, "{")
	rightKHCount := strings.Count(cont, "}")
	fmt.Println("leftKHCount :=", leftKHCount, "\trightKHCount:=", rightKHCount)
}
