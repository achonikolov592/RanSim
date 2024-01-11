package SecureDeleteFile

import (
	"helpers"
	"os"
)

func SecureDelete(filename string, nameOfLogFile string) {
	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_WRONLY, 0666)

	if err != nil {
<<<<<<< HEAD
		helpers.WriteLog(nameOfLogFile, err.Error(), 1)
=======
		helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
		os.Exit(5)
	}

	stats, err := file.Stat()

	if err != nil {
<<<<<<< HEAD
		helpers.WriteLog(nameOfLogFile, err.Error(), 1)
=======
		helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
		os.Exit(6)
	}

	sizeOfFile := stats.Size()
	chunk := int64(1 * (1 << 18))

	numberOfParts := (sizeOfFile / chunk) + 1
	position := int64(0)

	for i := int64(0); i < numberOfParts; i++ {
		var howMuchToChange int64
		if sizeOfFile-(i+1)*chunk < 0 {
			howMuchToChange = sizeOfFile - i*chunk
		} else {
			howMuchToChange = chunk
		}

		zeros := make([]byte, howMuchToChange)

		_, err = file.WriteAt([]byte(zeros), position)
		if err != nil {
<<<<<<< HEAD
			helpers.WriteLog(nameOfLogFile, err.Error(), 1)
=======
			helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
			os.Exit(7)
		}

		position += howMuchToChange
	}

	file.Close()

<<<<<<< HEAD
	if err = os.Remove(filename); err != nil {
		helpers.WriteLog(nameOfLogFile, err.Error(), 1)
=======
	if err = os.Remove(stats.Name()); err != nil {
		helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
		os.Exit(8)
	}

}
