package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"io"

	"public-vault-ms/database"

	"github.com/google/uuid"
)

// EncryptionKey es la clave utilizada para el cifrado AES-256
var EncryptionKey = []byte("public-vault-key-test-cba-2023")

// GenerateToken genera un UUID como token
func GenerateToken() string {
	return uuid.NewString()
}

// EncryptCard cifra el número de tarjeta utilizando AES-256
func EncryptCard(cardNumber string) (string, error) {
	block, err := aes.NewCipher(EncryptionKey)
	if err != nil {
		return "", err
	}

	plainText := []byte(cardNumber)
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return base64.URLEncoding.EncodeToString(cipherText), nil
}

// DecryptCard descifra el número de tarjeta utilizando AES-256
func DecryptCard(encryptedCard string) (string, error) {
	block, err := aes.NewCipher(EncryptionKey)
	if err != nil {
		return "", err
	}

	cipherText, err := base64.URLEncoding.DecodeString(encryptedCard)
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}

// SaveCard guarda el token y el número de tarjeta cifrado en la base de datos
func SaveCard(token string, encryptedCard string) error {
	_, err := database.DB.Exec("INSERT INTO cards (token, encrypted_card) VALUES ($1, $2)", token, encryptedCard)
	return err
}

// GetCardByToken obtiene el número de tarjeta cifrado relacionado a un token
func GetCardByToken(token string) (string, error) {
	var encryptedCard string
	err := database.DB.QueryRow("SELECT encrypted_card FROM cards WHERE token = $1", token).Scan(&encryptedCard)
	if err == sql.ErrNoRows {
		return "", errors.New("token not found")
	}
	return encryptedCard, err
}

// TokenizeCard procesa un número de tarjeta para generar un token y almacenarlo con la tarjeta encriptada.
func TokenizeCard(cardNumber string) (string, error) {
	if len(cardNumber) != 16 {
		return "", errors.New("el número de tarjeta debe tener 16 dígitos")
	}
	token := GenerateToken()
	encryptedCard, err := EncryptCard(cardNumber)
	if err != nil {
		return "", err
	}
	err = SaveCard(token, encryptedCard)
	if err != nil {
		return "", err
	}
	return token, nil
}

// DetokenizeCard obtiene el número de tarjeta original utilizando un token.
func DetokenizeCard(token string) (string, error) {
	encryptedCard, err := GetCardByToken(token)
	if err != nil {
		return "", err
	}
	cardNumber, err := DecryptCard(encryptedCard)
	if err != nil {
		return "", err
	}
	return cardNumber, nil
}
