package utils

import (
	// "github.com/gin-contrib/static"
	"os"
	"path/filepath"
)

func GetAppRoot() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}
