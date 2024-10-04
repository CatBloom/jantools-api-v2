package utils

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateUUIDWithoutHyphens() (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(uuid.String(), "-", ""), nil
}
