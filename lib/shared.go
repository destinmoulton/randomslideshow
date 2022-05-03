package lib

import (
	"crypto/sha1"
	"fmt"
)

// Get the sha1 hash hex of a directory
func HashString(dir string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(dir)))
}
