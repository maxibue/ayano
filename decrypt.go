package main

import (
	"crypto/aes"
	"crypto/cipher"

	"golang.org/x/crypto/scrypt"
)

// Decrypt decrypts the ciphertext using the provided password, salt, and nonce,
// and returns the plaintext.
func Decrypt(password, salt, nonce, ciphertext []byte) ([]byte, error) {
	key, err := scrypt.Key(password, salt, 32768, 8, 1, 32)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
