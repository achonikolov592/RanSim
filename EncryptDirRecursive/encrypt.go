package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
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

		if !(info.IsDir()) {
			if err != nil {
				fmt.Println(err)
				os.Exit(7)
			}
			encryptedText, err := os.ReadFile(file)
			if err != nil {
				fmt.Println(err)
				os.Exit(8)
			}

			nonce := make([]byte, c.NonceSize())
			if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
				fmt.Println(err)
				os.Exit(9)
			}
			encryptedText = c.Seal(nonce, nonce, encryptedText, nil)

			err = os.WriteFile(file, encryptedText, 0666)
			if err != nil {
				fmt.Println(err)
				os.Exit(10)
			}

			fmt.Println("Encrypted: " + file)

		} else {
			go encrypt(info.Name(), c)
		}
	}

}

func EncryptDir(dirToEncrypt string) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Println("The key for full file encryption is" + hex.EncodeToString(key))

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

	fmt.Println("Strating test: EncryptDir")

	EncryptDir(os.Args[1])
}

//abc
