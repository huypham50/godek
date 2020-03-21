package db

import (
	"strconv"
	"strings"
)

// GenerateUsername -> generate unique username from email and db count
func GenerateUsername(email string, count int) string {
	components := strings.Split(email, "@")
	emailUsername := components[0]
	count = count + 1

	return emailUsername + strconv.Itoa(count)
}
