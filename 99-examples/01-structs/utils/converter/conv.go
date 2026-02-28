package converter

import (
	"fmt"
	"strconv"
)

func ConvToInt(str string) int {
	var i int
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return 0
	}
	return i
}
