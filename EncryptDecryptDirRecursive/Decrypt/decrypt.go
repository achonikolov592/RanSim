package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
)

func decrypt(dirToDecrypt string, c cipher.AEAD) {
	var filesInDir []string
	err := filepath.Walk(dirToDecrypt, func(path string, info os.FileInfo, err error) error {
		if path != dirToDecrypt {
			filesInDir = append(filesInDir, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}

	for _, file := range filesInDir {
		info, err := os.Stat(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(5)
		}

		if !(info.IsDir()) {
			if err != nil {
				fmt.Println(err)
				os.Exit(6)
			}
			encryptedText, err := os.ReadFile(file)
			if err != nil {
				fmt.Println(err)
				os.Exit(7)
			}

			non, text := encryptedText[:c.NonceSize()], encryptedText[c.NonceSize():]

			decb, err := c.Open(nil, non, text, nil)
			if err != nil {
				fmt.Println(err)
				os.Exit(8)
			}
			err = os.WriteFile(file, decb, 0666)
			if err != nil {
				fmt.Println(err)
				os.Exit(9)
			}

		} else {
			go decrypt(info.Name(), c)
		}
	}

}

func DecryptDir(dirToDecrypt string, nameOfEncryptionInfoFile string) {
	file, _ := os.ReadFile(nameOfEncryptionInfoFile)
	key, _ := hex.DecodeString(string(file[1:65]))

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	c, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	decrypt(dirToDecrypt, c)
}
