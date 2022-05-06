package main

import (
	"tools/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		return
	}
}
