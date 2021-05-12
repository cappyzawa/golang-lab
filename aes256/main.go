package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

// original code is https://gist.github.com/cannium/c167a19030f2a3c6adbb5a5174bea3ff
func main() {
	key := []byte("AES256Key-32Characters1234567890")
	plaintext := []byte("cappyzawa")
	fmt.Println("plaintext:", string(plaintext))

	block, err := aes.NewCipher(key)
	if err != nil {
		printErr(fmt.Sprintf("failed to generate cypher: %v", err))
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		printErr(fmt.Sprintf("failed to generate gcm: %v", err))
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		printErr(fmt.Sprintf("failed to read nonce: %v", err))
	}
	fmt.Printf("nonce: %x\n", nonce)
	// encrypt
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	fmt.Println("encrypted:", string(ciphertext))

	storedtext := append(nonce, ciphertext...)
	// decrypt
	decrypted, err := aesgcm.Open(nil, storedtext[:aesgcm.NonceSize()], storedtext[aesgcm.NonceSize():], nil)
	if err != nil {
		printErr(fmt.Sprintf("failed to decrypt: %v", err))
	}
	fmt.Println("decrypted:", string(decrypted))
}

func printErr(msg string) {
	fmt.Fprintf(os.Stderr, "%s\n", msg)
	os.Exit(1)
}
