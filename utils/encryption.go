// utils/encryption.go
package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var encryptionKey = []byte("public-vault-key-test-cba-2023") // Clave de 32 bytes

// Encrypt encripta un texto plano utilizando AES-256
func Encrypt(plainText string) (string, error) {
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return "", err
	}

	nonce := encryptionKey[:block.BlockSize()] // Usamos una parte de la clave como nonce
	ciphertext := make([]byte, len(plainText))
	stream := cipher.NewCFBEncrypter(block, nonce)
	stream.XORKeyStream(ciphertext, []byte(plainText))

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt desencripta un texto encriptado utilizando AES-256
func Decrypt(encryptedText string) (string, error) {
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return "", err
	}

	nonce := encryptionKey[:block.BlockSize()] // Usamos la misma parte de la clave como nonce
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	plainText := make([]byte, len(ciphertext))
	stream := cipher.NewCFBDecrypter(block, nonce)
	stream.XORKeyStream(plainText, ciphertext)

	return string(plainText), nil
}
