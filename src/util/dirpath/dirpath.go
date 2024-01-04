package dirpath

import (
	"os"
	"path/filepath"
)

func Join(elem ...string) string {
	return filepath.Join(elem...) + string(os.PathSeparator)
}
