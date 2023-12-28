package utils

import (
	// "github.com/gin-contrib/static"
	// "os"
	// "path/filepath"

	"github.com/kardianos/osext"
)

func GetAppRoot() string {
	folderPath, _ := osext.ExecutableFolder()
	// dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return folderPath
}
