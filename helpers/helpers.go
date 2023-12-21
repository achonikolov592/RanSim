package helpers

import (
	"fmt"
	"os"
)

func CreateTestFiles(dir string) {
	err := os.Mkdir(dir+"testfiles", 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = os.Mkdir(dir+"testfiles/sub", 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	f, err := os.Create(dir + "testfiles/a.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	_, err = f.WriteString("asdfghjk")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	f, err = os.Create(dir + "testfiles/sub/b.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	_, err = f.WriteString("asdfghjk")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func CreateLogFile(dir string, name string) {
	_, err := os.Create(dir + name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func WriteLog(file string, line string) {
	f, err := os.OpenFile(file, os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = f.WriteString(line + "\n")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
