package utils

import (
	// "github.com/gin-contrib/static"
	// "os"
	// "path/filepath"

	"os"
	"sync"
	// "github.com/kardianos/osext"
)

var path string
var once sync.Once

func GetAppRoot() string {
	once.Do(func() {
		// folderPath, _ := osext.ExecutableFolder()
		folderPath, _ := os.Getwd()
		path = folderPath
	})
	// dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return path
}
