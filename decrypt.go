package main

import (
	"crypto/aes"
	"fmt"
	"os"
)

func DecryptFile(inputfile, outfile *os.File) {
	decodedChunk := make([]byte,0)	
	chunk := make([]byte, ENCRYPT_CHUNK)

	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}

	for {
		n, err := inputfile.Read(chunk)		
		if err != nil {
			break
		}
		msgByte := make([]byte, ENCRYPT_CHUNK)
		c.Decrypt(msgByte, chunk[:n])
		decodedChunk = append(decodedChunk, msgByte...)
	}
	outfile.WriteString(string(decodedChunk))
	outfile.Sync()

}