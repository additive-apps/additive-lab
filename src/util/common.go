package util

import (
	"additive-apps/additive-lab/src/print"
	"os"
)

func HomePath() string {
	userHomeDir, err := os.UserHomeDir()

	if err != nil {
		print.Fatal(err)
	}

	return userHomeDir
}

func ResourceNotExists(resourceName string, path string, operations ...func()) bool {
	_, err := os.Stat(path)

	if err != nil {
		for _, operation := range operations {
			operation()
		}
	} else {
		print.Warning("%s %s already exists.", resourceName, path)
	}

	return os.IsNotExist(err)
}
