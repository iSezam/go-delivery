// Package app: DEV environment loader (copy file without "_" for DEV environment)
package app

import (
	"fmt"
	"path/filepath"

	"github.com/joho/godotenv"
)

func init() {
	relPath, _ := filepath.Abs(".")
	err := godotenv.Load(fmt.Sprintf("%s%s", relPath, "\\dev.env"))
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}
