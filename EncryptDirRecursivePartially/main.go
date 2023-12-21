package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	mathrand "math/rand"
	"os"
	"path/filepath"
	"strconv"
)

func encrypt(dirToEncrypt string, c cipher.AEAD) {
	var filesInDir []string
	err := filepath.Walk(dirToEncrypt, func(path string, info os.FileInfo, err error) error {
		if path != dirToEncrypt {
			filesInDir = append(filesInDir, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(5)
	}

	for _, file := range filesInDir {
		info, err := os.Stat(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(6)
		}

		f, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(7)
		}

		if !(info.IsDir()) {
			if err != nil {
				fmt.Println(err)
				os.Exit(8)
			}

			howMuchToEncrypt := info.Size() / int64(8)
			whereToStart := int64(howMuchToEncrypt * mathrand.Int63n(8))
			filetext, err := os.ReadFile(file)
			if err != nil {
				fmt.Println(err)
				os.Exit(9)
			}
			textToEncrypt := make([]byte, howMuchToEncrypt)
			_, err = f.ReadAt(textToEncrypt, whereToStart)
			if err != nil {
				fmt.Println(err)
				os.Exit(10)
			}

			nonce := make([]byte, c.NonceSize())
			if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
				fmt.Println(err)
				os.Exit(11)
			}
			encryptedText := c.Seal(nonce, nonce, textToEncrypt, nil)

			fmt.Println("Where it started to encrypt: " + strconv.Itoa(int(whereToStart)))
			fmt.Println("How much bytes it encrypted: " + strconv.Itoa(int(howMuchToEncrypt)))
			fmt.Println("How much bytes is product: " + strconv.Itoa(len(encryptedText)))

			var finaltext []byte
			finaltext = append(finaltext, filetext[:whereToStart]...)
			finaltext = append(finaltext, encryptedText...)
			finaltext = append(finaltext, filetext[whereToStart+howMuchToEncrypt:]...)
			err = os.WriteFile(file, finaltext, 0666)
			if err != nil {
				fmt.Println(err)
				os.Exit(12)
			}

			fmt.Println("Encrypted: " + file)

		} else {
			encrypt(info.Name(), c)
		}
	}
}

func EncryptFilesInDir(dirToEncrypt string) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Println("The key of partial file encryption is: " + hex.EncodeToString(key))

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	c, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}

	encrypt(dirToEncrypt, c)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("invalid number of arguments")
		os.Exit(1)
	}

	fmt.Println("Strating test: EncryptDirPartially")

	EncryptFilesInDir(os.Args[1])
}
