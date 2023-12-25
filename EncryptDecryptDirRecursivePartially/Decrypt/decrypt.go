package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"go/scanner"
	"go/token"
	"helpers"
	"os"
	"path/filepath"
	"strconv"
)

type WhereIsEncrypted struct {
	whereToStart       int64
	howMuchIsEncrypted int64
}

func decrypt(dirToDecrypt string, c cipher.AEAD, wherePartsAreEncryptedInLexicalOrder []WhereIsEncrypted, whichIteration int, nameOfLogFile string) {
	var filesInDir []string
	err := filepath.Walk(dirToDecrypt, func(path string, info os.FileInfo, err error) error {
		if path != dirToDecrypt {
			filesInDir = append(filesInDir, path)
		}
		return nil
	})
	if err != nil {
		helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
		os.Exit(4)
	}

	for _, file := range filesInDir {
		info, err := os.Stat(file)
		if err != nil {
			helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
			os.Exit(5)
		}

		if !(info.IsDir()) {
			if err != nil {
				helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
				os.Exit(6)
			}
			encryptedText, err := os.ReadFile(file)
			if err != nil {
				helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
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
				helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
				os.Exit(8)
			}
			whichIteration++
			helpers.WriteLog(nameOfLogFile, "Decrypted: "+file)

		} else {
			decrypt(info.Name(), c, wherePartsAreEncryptedInLexicalOrder, whichIteration, nameOfLogFile)
		}
	}
}

func DecryptDir(dirToDecrypt string, nameOfLogFile string, nameOfEncryptionInfoFile string) {
	var key []byte
	var values []WhereIsEncrypted
	EncryptionInfo, err := os.ReadFile(nameOfEncryptionInfoFile)
	if err != nil {
		helpers.WriteLog(nameOfLogFile, "Error3: "+err.Error())
	}

	fileSet := token.NewFileSet()
	file := fileSet.AddFile("info", fileSet.Base(), len(EncryptionInfo)-65)

	var s scanner.Scanner
	s.Init(file, EncryptionInfo[65:], nil, 0)

	key, _ = hex.DecodeString(string(EncryptionInfo[1:65]))

	var whereToStart, howMuchIsEncrypted, i int
	i = 0
	for {
		_, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		} else if tok == token.INT {
			if i == 0 {
				whereToStart, _ = strconv.Atoi(lit)
				i++
			} else if i == 1 {
				howMuchIsEncrypted, _ = strconv.Atoi(lit)
				values = append(values, WhereIsEncrypted{int64(whereToStart), int64(howMuchIsEncrypted)})
				i = 0
			}
		}
	}

	block, _ := aes.NewCipher(key)
	if err != nil {
		helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
		os.Exit(2)
	}

	c, err := cipher.NewGCM(block)
	if err != nil {
		helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
		os.Exit(3)
	}

	decrypt(dirToDecrypt, c, values, 0, nameOfLogFile)
}
