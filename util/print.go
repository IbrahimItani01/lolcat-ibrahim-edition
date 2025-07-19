package util

import (
	"fmt"
)

func PrintOutput(output []rune) {
    for j := 0; j < len(output); j++ {
        r, g, b := GenerateRGB(j)
        fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[j])
    }
    fmt.Println()
}