// utils/uuid_generator.go
package utils

import (
	"github.com/google/uuid"
)

// GenerateUUID genera un UUID único
func GenerateUUID() string {
	return uuid.New().String()
}
