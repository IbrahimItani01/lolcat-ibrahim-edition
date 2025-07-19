package main

import (
	"bufio"
	"fmt"
	"io"
	util "lolcat-ibrahim-edition/util"
	"os"
)

func main() {
	fileInfo, _ := os.Stdin.Stat()
	var output []rune

	if fileInfo.Mode() & os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
        fmt.Println("Usage: fortune | gorainbow")
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
            break
        }
		output = append(output, input)
	}
	util.PrintOutput(output)
}