// models/card.go
package models

// Card representa una tarjeta almacenada en la base de datos
type Card struct {
	Token         string // Token único de la tarjeta
	EncryptedCard string // Número de tarjeta encriptado
}
