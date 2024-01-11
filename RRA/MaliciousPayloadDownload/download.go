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
	nameOfLogFile = helpers.CreateLogFileIfItDoesNotExist("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/MaliciousPayloadDownload/", "DownloadMalicious")

	helpers.WriteLog(nameOfLogFile, "Starting test: DownloadingMaliciousPayload", 1)

	out, err := os.Create("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/MaliciousPayloadDownload/out1.zip")
	if err != nil {
		helpers.WriteLog(nameOfLogFile, err.Error(), 1)
		os.Exit(2)
	}
	defer out.Close()

	resp, err := http.Get("https://github.com/kh4sh3i/Ransomware-Samples/raw/main/Petya/Ransomware.Petya.zip")
	if err != nil {
		helpers.WriteLog(nameOfLogFile, err.Error(), 1)
		os.Exit(3)
=======
	nameOfLogFile = helpers.CreateLogFileIfItDoesNotExist("./", "eicar")

	helpers.WriteLog(nameOfLogFile, "Starting test: DownloadingMaliciousPayload")

	out, err := os.Create("/out1.exe")
	if err != nil {
		os.Exit(1)
		helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
	}
	defer out.Close()

	resp, err := http.Get("https://bazaar.abuse.ch/download/124e57afaa34fee7a6405531b05ceb4466e94f9e1858db976188a36de58156b3/")
	if err != nil {
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
	}

	_, err = os.Open("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/MaliciousPayloadDownload/out1.zip")
	if err != nil {
		helpers.WriteLog(nameOfLogFile, err.Error(), 1)
		os.Exit(1)
	}

	helpers.WriteLog(nameOfLogFile, "Ending test: DownloadingMaliciousPayload", 1)

	os.Exit(0)
}

//a
=======
		helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
		os.Exit(3)
	}

	_, err = os.Open("./out1.exe")
	if err != nil {
		helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
		os.Exit(4)
	}

	helpers.WriteLog(nameOfLogFile, "Ending test: DownloadingMaliciousPayload")
}

//
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
