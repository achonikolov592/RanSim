package main

import (
	"fmt"
	"os"
)

func main() {
	if err := os.Remove("./test.txt"); err != nil {
		fmt.Println(err)
	}
}
