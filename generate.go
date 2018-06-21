package uuidgen

import "fmt"

// GenerateUUID will generate a UUID
func GenerateUUID() (string, error) {
	b, err := generate()
	if err != nil {
		return "", err
	}

	s := fmt.Sprintf("%4x-%2x-%2x-%2x-%12x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return s, nil
}
