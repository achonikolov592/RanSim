package helpers

import (
	"fmt"
	"os"
)

func CreateTestFiles(dir string, logfile string) {
	WriteLog(logfile, "Start creating test files")

	if _, err := os.Stat(dir + "testfiles"); err != nil {
		err := os.Mkdir(dir+"testfiles", 0777)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	if _, err := os.Stat(dir + "testfiles/sub"); err != nil {
		err := os.Mkdir(dir+"testfiles/sub", 0777)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
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

func CreateLogFileIfItDoesNotExist(dir string, name string) string {

	i, err := os.Stat(dir + name + ".log")

	if err != nil {
		_, err := os.Create(dir + name + ".log")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		return dir + name + ".log"
	}

	return i.Name()

}

func WriteLog(logfile string, line string) {
	f, err := os.OpenFile(logfile, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = f.WriteString("\n" + line + "\n")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
