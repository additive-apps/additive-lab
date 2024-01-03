package file

import (
	"additive-apps/additive-lab/src/print"
	"additive-apps/additive-lab/src/util"
	"io"
	"os"
)

func NotExists(path string, operations ...func()) bool {
	return util.ResourceNotExists("File", path, operations...)
}

func Content(path string) string {
	file, err := os.Open(path)

	if err != nil {
		print.Fatal("Error opening file:", err)
	}

	defer file.Close()
	content, err := io.ReadAll(file)

	if err != nil {
		print.Fatal("Error reading file:", err)
	}

	return string(content)
}

func Write(path string, content string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	defer file.Close()
	if err != nil {
		print.Fatal("Error opening file %s: ", path, err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		print.Fatal("Error writting to file %s: ", path, err)
	}
}
