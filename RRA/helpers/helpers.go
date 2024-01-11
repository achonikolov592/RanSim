package helpers

import (
	"fmt"
	"os"
<<<<<<< HEAD
	"time"
=======
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
)

func AppendAtTheEnd(file string, logfile string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
<<<<<<< HEAD
		WriteLog(logfile, err.Error(), 1)
=======
		WriteLog(logfile, "Error: "+err.Error())
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
	}

	_, err = f.Write([]byte("//a"))
	if err != nil {
<<<<<<< HEAD
		WriteLog(logfile, err.Error(), 1)
=======
		WriteLog(logfile, "Error: "+err.Error())
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
	}
}

func CreateTestFiles(dir string, logfile string) {
<<<<<<< HEAD
	WriteLog(logfile, "Start creating test files", 2)
=======
	WriteLog(logfile, "Start creating test files")
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350

	if _, err := os.Stat(dir + "testfiles"); err != nil {
		err := os.Mkdir(dir+"testfiles", 0777)
		if err != nil {
<<<<<<< HEAD
			WriteLog(logfile, err.Error(), 1)
=======
			WriteLog(logfile, "Error: "+err.Error())
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
			os.Exit(1)
		}
	}

	if _, err := os.Stat(dir + "testfiles/sub"); err != nil {
		err := os.Mkdir(dir+"testfiles/sub", 0777)
		if err != nil {
<<<<<<< HEAD
			WriteLog(logfile, err.Error(), 1)
=======
			WriteLog(logfile, "Error: "+err.Error())
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
			os.Exit(1)
		}
	}

<<<<<<< HEAD
	f, err := os.Create(dir + "testfiles/c.txt")
	if err != nil {
		WriteLog(logfile, err.Error(), 1)
=======
	f, err := os.Create(dir + "testfiles/a.txt")
	if err != nil {
		WriteLog(logfile, "Error: "+err.Error())
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
		os.Exit(2)
	}

	_, err = f.WriteString("asdfghjk")
	if err != nil {
<<<<<<< HEAD
		WriteLog(logfile, err.Error(), 1)
=======
		WriteLog(logfile, "Error: "+err.Error())
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
		os.Exit(2)
	}

	f, err = os.Create(dir + "testfiles/sub/b.txt")
	if err != nil {
<<<<<<< HEAD
		WriteLog(logfile, err.Error(), 1)
=======
		WriteLog(logfile, "Error: "+err.Error())
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
		os.Exit(2)
	}

	_, err = f.WriteString("asdfghjk")
	if err != nil {
<<<<<<< HEAD
		WriteLog(logfile, err.Error(), 1)
=======
		WriteLog(logfile, "Error: "+err.Error())
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
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

func RemoveTestFilesIfExists(dir string) {
	_, err := os.Stat(dir + "testfiles")
	if err == nil {
		os.RemoveAll(dir + "testfiles")
	}
}

<<<<<<< HEAD
func WriteLog(logfile string, line string, opt int) {

	f, err := os.OpenFile(logfile, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}

	var stringToWrite string
	if opt == 0 {
		stringToWrite = "\n" + line + "\n"
	}
	if opt == 1 { //err
		stringToWrite = "[ERROR]: " + line + " AT " + time.Now().Format(time.RFC822) + "\n"
	} else { //info
		stringToWrite = "[INFO]: " + line + " AT " + time.Now().Format(time.RFC822) + "\n"
	}
	_, err = f.WriteString(stringToWrite)
	if err != nil {
		fmt.Println(err)
		os.Exit(101)
=======
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
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
	}
}
