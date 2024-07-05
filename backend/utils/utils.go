package utils

import (
	"fmt"
	"strconv"
)

func ParseInt(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("Error parsing float:", err)
	}
	return i
}

func ParseFloat(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("Error parsing int:", err)
	}
	return f
}
