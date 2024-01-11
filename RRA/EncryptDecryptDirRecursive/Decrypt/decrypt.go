package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
<<<<<<< HEAD
	"helpers"
=======
	"fmt"
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
	"os"
	"path/filepath"
)

<<<<<<< HEAD
func decrypt(dirToDecrypt string, c cipher.AEAD, nameOfLogFile string) {
=======
func decrypt(dirToDecrypt string, c cipher.AEAD) {
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
	var filesInDir []string
	err := filepath.Walk(dirToDecrypt, func(path string, info os.FileInfo, err error) error {
		if path != dirToDecrypt {
			filesInDir = append(filesInDir, path)
		}
		return nil
	})
	if err != nil {
<<<<<<< HEAD
		helpers.WriteLog(nameOfLogFile, err.Error(), 1)
=======
		fmt.Println(err)
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
		os.Exit(4)
	}

	for _, file := range filesInDir {
		info, err := os.Stat(file)
		if err != nil {
<<<<<<< HEAD
			helpers.WriteLog(nameOfLogFile, err.Error(), 1)
=======
			fmt.Println(err)
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
			os.Exit(5)
		}

		if !(info.IsDir()) {
			if err != nil {
<<<<<<< HEAD
				helpers.WriteLog(nameOfLogFile, err.Error(), 1)
=======
				fmt.Println(err)
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
				os.Exit(6)
			}
			encryptedText, err := os.ReadFile(file)
			if err != nil {
<<<<<<< HEAD
				helpers.WriteLog(nameOfLogFile, err.Error(), 1)
=======
				fmt.Println(err)
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
				os.Exit(7)
			}

			non, text := encryptedText[:c.NonceSize()], encryptedText[c.NonceSize():]

			decb, err := c.Open(nil, non, text, nil)
			if err != nil {
<<<<<<< HEAD
				helpers.WriteLog(nameOfLogFile, err.Error(), 1)
=======
				fmt.Println(err)
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
				os.Exit(8)
			}
			err = os.WriteFile(file, decb, 0666)
			if err != nil {
<<<<<<< HEAD
				helpers.WriteLog(nameOfLogFile, err.Error(), 1)
=======
				fmt.Println(err)
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
				os.Exit(9)
			}

		} else {
<<<<<<< HEAD
			go decrypt(info.Name(), c, nameOfLogFile)
=======
			go decrypt(info.Name(), c)
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
		}
	}

}

<<<<<<< HEAD
func DecryptDir(dirToDecrypt string, nameOfEncryptionInfoFile string, nameOfLogFile string) {
=======
func DecryptDir(dirToDecrypt string, nameOfEncryptionInfoFile string) {
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
	file, _ := os.ReadFile(nameOfEncryptionInfoFile)
	key, _ := hex.DecodeString(string(file[1:65]))

	block, err := aes.NewCipher(key)
	if err != nil {
<<<<<<< HEAD
		helpers.WriteLog(nameOfLogFile, err.Error(), 1)
=======
		fmt.Println(err)
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
		os.Exit(2)
	}
	c, err := cipher.NewGCM(block)
	if err != nil {
<<<<<<< HEAD
		helpers.WriteLog(nameOfLogFile, err.Error(), 1)
		os.Exit(3)
	}
	decrypt(dirToDecrypt, c, nameOfLogFile)
=======
		fmt.Println(err)
		os.Exit(3)
	}
	decrypt(dirToDecrypt, c)
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
}
