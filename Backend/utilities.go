package main

import (
	"strings"
)

func indexOf(data string, element string) int {
	splitData := strings.Split(data, "")

	for i := len(data) - 1; i >= 0; i-- {
		if splitData[i] == element {
			return i //Index Position
		}
	}
	return -1 //not found.
}
