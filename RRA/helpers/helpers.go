package helpers

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func AppendAtTheEnd(file string, logfile string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		WriteLog(logfile, err.Error(), 1)
	}

	_, err = f.Write([]byte("//a"))
	if err != nil {
		WriteLog(logfile, err.Error(), 1)
	}
}

func CreateTestFiles(dir string, logfile string) {
	WriteLog(logfile, "Start creating test files", 2)

	if _, err := os.Stat(dir + "testfiles"); err != nil {
		err := os.Mkdir(dir+"testfiles", 0777)
		if err != nil {
			WriteLog(logfile, err.Error(), 1)
			os.Exit(1)
		}
	}

	if _, err := os.Stat(dir + "testfiles/sub"); err != nil {
		err := os.Mkdir(dir+"testfiles/sub", 0777)
		if err != nil {
			WriteLog(logfile, err.Error(), 1)
			os.Exit(1)
		}
	}

	f, err := os.Create(dir + "testfiles/c.txt")
	if err != nil {
		WriteLog(logfile, err.Error(), 1)
		os.Exit(2)
	}

	_, err = f.WriteString("asdfghjk")
	if err != nil {
		WriteLog(logfile, err.Error(), 1)
		os.Exit(2)
	}

	f, err = os.Create(dir + "testfiles/sub/b.txt")
	if err != nil {
		WriteLog(logfile, err.Error(), 1)
		os.Exit(2)
	}

	_, err = f.WriteString("asdfghjk")
	if err != nil {
		WriteLog(logfile, err.Error(), 1)
		os.Exit(2)
	}
}

func CreateMultipleTestFiles(dir string, logfile string) {
	var files []*os.File
	//find files

	err := filepath.Walk(dir+"../TestFiles", func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			if path != "..\\TestFiles\\DirCreationInfo.log" {
				f, err := os.OpenFile(path, os.O_RDONLY, 0666)
				if err != nil {
					return err
				}
				files = append(files, f)
			}
		}
		return nil
	})

	if err != nil {
		WriteLog(logfile, err.Error(), 1)
	}

	f, err := os.OpenFile("../TestFiles/DirCreationInfo.log", os.O_RDONLY, 0666)
	if err != nil {
		WriteLog(logfile, err.Error(), 1)
	}

	DirCreationInfoString, err := io.ReadAll(f)
	if err != nil {
		WriteLog(logfile, err.Error(), 1)
	}
	DirCreationInfo := strings.Fields(string(DirCreationInfoString))

	numberOfDirs, _ := strconv.Atoi(DirCreationInfo[0])
	numberOfFIles, _ := strconv.Atoi(DirCreationInfo[1])

	err = os.RemoveAll("./testFilesParent")
	if err == nil {
		err = os.Mkdir("./testFilesParent", 0777)
		if err != nil {
			WriteLog(logfile, err.Error(), 1)
		}
	}

	for i := 0; i < numberOfDirs; i++ {
		testFileDir := dir + "testFilesParent/testfiles" + strconv.Itoa(i) + "/"
		os.Mkdir(testFileDir, 0777)
		for j := 0; j < len(files); j++ {
			fileContent, err := io.ReadAll(files[j])
			if err != nil {
				WriteLog(logfile, err.Error(), 1)
			}
			for iof := 0; iof < numberOfFIles; iof++ {
				f, err := os.Create(testFileDir + strconv.Itoa(iof) + files[j].Name()[13:])
				if err != nil {
					WriteLog(logfile, err.Error(), 1)
				}
				_, err = f.WriteString(string(fileContent))
				if err != nil {
					WriteLog(logfile, err.Error(), 1)
				}
			}
		}
		testFileSubDir := testFileDir + "/sub/"
		os.Mkdir(testFileSubDir, 0777)
		for j := 0; j < len(files); j++ {
			fileContent, err := io.ReadAll(files[j])
			if err != nil {
				WriteLog(logfile, err.Error(), 1)
			}
			for iof := 0; iof < numberOfFIles; iof++ {
				f, err := os.Create(testFileSubDir + files[j].Name() + strconv.Itoa(iof))
				if err != nil {
					WriteLog(logfile, err.Error(), 1)
				}
				_, err = f.WriteString(string(fileContent))
				if err != nil {
					WriteLog(logfile, err.Error(), 1)
				}
			}
		}
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
	os.RemoveAll(dir + "testfiles")
	os.RemoveAll(dir + "testFilesParent")
}

func WriteLog(logfile string, line string, opt int) {
	f, err := os.OpenFile(logfile, os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(f)
		os.Exit(100)
	}

	var stringToWrite string
	if opt == 0 {
		stringToWrite = line + "\n"
	} else if opt == 1 { //err
		stringToWrite = "[ERROR]: " + line + " AT " + time.Now().Format(time.RFC822) + "\n"
	} else { //info
		stringToWrite = "[INFO]: " + line + " AT " + time.Now().Format(time.RFC822) + "\n"
	}
	_, err = f.WriteString(stringToWrite)
	if err != nil {
		fmt.Println(err)
		os.Exit(101)
	}
}
