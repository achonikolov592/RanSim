package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./test.txt", os.O_RDONLY|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	stats, err := file.Stat()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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
			fmt.Println(err)
			os.Exit(2)
		}

		position += howMuchToChange
	}

	file.Close()

	if err = os.Remove(stats.Name()); err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
}
