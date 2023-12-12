package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
)

type whereIsEncrypted struct {
	whereToStart       int64
	howMuchIsEncrypted int64
}

func decrypt(dirToDecrypt string, c cipher.AEAD, wherePartsAreEncryptedInLexicalOrder []whereIsEncrypted, whichIteration int) {
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
			non, text := encryptedText[wherePartsAreEncryptedInLexicalOrder[whichIteration].whereToStart:wherePartsAreEncryptedInLexicalOrder[whichIteration].whereToStart+int64(c.NonceSize())], encryptedText[wherePartsAreEncryptedInLexicalOrder[whichIteration].whereToStart+int64(c.NonceSize()):wherePartsAreEncryptedInLexicalOrder[whichIteration].whereToStart+wherePartsAreEncryptedInLexicalOrder[whichIteration].howMuchIsEncrypted]

			decb, err := c.Open(nil, non, text, nil)

			var finaltext []byte
			finaltext = append(finaltext, encryptedText[:wherePartsAreEncryptedInLexicalOrder[whichIteration].whereToStart]...)
			finaltext = append(finaltext, decb...)
			finaltext = append(finaltext, encryptedText[wherePartsAreEncryptedInLexicalOrder[whichIteration].whereToStart+wherePartsAreEncryptedInLexicalOrder[whichIteration].howMuchIsEncrypted:]...)

			err = os.WriteFile(file, finaltext, 0666)
			if err != nil {
				fmt.Println(err)
				os.Exit(8)
			}
			whichIteration++

		} else {
			decrypt(info.Name(), c, wherePartsAreEncryptedInLexicalOrder, whichIteration)
		}
	}
}

func DecryptDir(dirToDecrypt string, key string, wherePartsAreEncryptedInLexicalOrder []whereIsEncrypted) {
	k, _ := hex.DecodeString(key)
	block, err := aes.NewCipher(k)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	c, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	decrypt(dirToDecrypt, c, wherePartsAreEncryptedInLexicalOrder, 0)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("invalid number of arguments")
		os.Exit(1)
	}
	var a []whereIsEncrypted
	a = append(a, []whereIsEncrypted{whereIsEncrypted{7, 29}, whereIsEncrypted{0, 29}}...)
	DecryptDir(os.Args[1], os.Args[2], a)
}
