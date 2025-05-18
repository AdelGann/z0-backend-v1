package gen

import (
	"crypto/rand"
	"math/big"
)

// GenerateCode genera un código aleatorio de la longitud especificada
// compuesto por dígitos (0-9) y letras mayúsculas (A-Z).
func GenerateCode(length int) ([]byte, error) {
	code := make([]byte, length)
	for i := range code {
		for {
			num, err := rand.Int(rand.Reader, big.NewInt(127))
			if err != nil {
				return nil, err
			}
			n := num.Int64()
			if (n >= 48 && n <= 57) || (n >= 65 && n <= 90) {
				code[i] = byte(n)
				break
			}
		}
	}
	return code, nil
}
