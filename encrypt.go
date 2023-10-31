package main

import (
	"crypto/aes"
	"fmt"
	"os"
)



func EncryptFile(in *os.File, out *os.File) {
	encryptedChunk := make([]byte,0)	
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}

	for {
		chunk := make([]byte, ENCRYPT_CHUNK)
		_, err := in.Read(chunk)		
		if err != nil {
			break
		}
		msgByte := make([]byte, ENCRYPT_CHUNK)
		c.Encrypt(msgByte, chunk)
		encryptedChunk = append(encryptedChunk, msgByte...)
	}

	_, err = out.Write(encryptedChunk)
	if err != nil {
		fmt.Println(err)
	}

	out.Sync()
}

