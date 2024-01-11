package main

import (
	"helpers"
	"io"
	"net/http"
	"os"
)

var nameOfLogFile string

func main() {
<<<<<<< HEAD
	nameOfLogFile = helpers.CreateLogFileIfItDoesNotExist("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/Eicar/", "eicar")

	helpers.WriteLog(nameOfLogFile, "Starting test: EicarTest", 2)

	out, err := os.Create("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/Eicar/out.txt")
	if err != nil {
		os.Exit(2)
		helpers.WriteLog(nameOfLogFile, err.Error(), 1)
=======
	nameOfLogFile = helpers.CreateLogFileIfItDoesNotExist("./", "eicar")

	helpers.WriteLog(nameOfLogFile, "Starting test: EicarTest")

	out, err := os.Create("./out1.txt")
	if err != nil {
		os.Exit(1)
		helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
	}
	defer out.Close()

	resp, err := http.Get("https://www.eicar.org/download/eicar-com-2/?wpdmdl=8842&refresh=656a05f4d5e5d1701447156")
	if err != nil {
<<<<<<< HEAD
		os.Exit(3)
		helpers.WriteLog(nameOfLogFile, err.Error(), 1)
=======
		os.Exit(2)
		helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
<<<<<<< HEAD
		helpers.WriteLog(nameOfLogFile, err.Error(), 1)
		os.Exit(4)
=======
		helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
		os.Exit(3)
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
	}

	_, err = os.Open("./out1.txt")
	if err != nil {
<<<<<<< HEAD
		helpers.WriteLog(nameOfLogFile, err.Error(), 1)
		os.Exit(1)
	}

	helpers.WriteLog(nameOfLogFile, "Ending test: EicarTest", 2)
	os.Exit(0)
}

//as
=======
		helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
		os.Exit(4)
	}

	helpers.WriteLog(nameOfLogFile, "Ending test: EicarTest")
}

//
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
