package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := ""
	str = "50 4F 53 54 0A 0A 61 70 70 6C 69 63 61 74 69 6F 6E 2F 6F 63 74 65 74 2D 73 74 72 65 61 6D 0A 54 75 65 2C 20 31 32 20 46 65 62 20 32 30 31 39 20 30 37 3A 32 30 3A 35 36 20 47 4D 54 0A 2F 77 69 73 68 73 74 61 72 74 2D 72 64 73 2D 62 61 63 6B 75 70 2F 69 6D 74 65 73 74 2F 74 65 73 74 2E 72 64 62 3F 75 70 6C 6F 61 64 49 64 3D 38 34 31 32 38 30 42 35 45 44 41 34 34 32 31 44 39 45 45 46 43 32 42 31 33 35 45 32 44 33 30 35"


	str = strings.Replace(str, " ", "", -1)
	//println(str)
	test, _ := hex.DecodeString(str)
	fmt.Println(string(test[:]))

	//bt := string_to_byte(strings.Split(str, ","))
	//
	//println(string(bt[:]))
}

func string_to_byte(sli []string) (bli []byte) {
	for i := 0; i < len(sli); i++ {
		iValue, _ := strconv.Atoi(sli[i])
		bli = append(bli, byte(iValue))
	}
	return
}
