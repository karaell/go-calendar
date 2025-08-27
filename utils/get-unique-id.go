package utils

import "github.com/google/uuid"

func GetUniqId() string {
	return uuid.New().String()
}
