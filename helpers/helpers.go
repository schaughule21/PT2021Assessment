package helpers

import (
	"crypto/rand"
	"fmt"
)

func Uuid(parts int) string {
	len := 4 * parts
	b := make([]byte, len)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return err.Error()
	}
	uuid := fmt.Sprintf("%X", b[0:3])
	for i := 4; i < len-2; i += 3 {
		uuid += fmt.Sprintf("-%X", b[i:i+3])
	}
	return uuid
}
