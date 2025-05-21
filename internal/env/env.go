package env

import (
	"fmt"
	"os"
	"strings"
)

func Load() error {
	return loadFile(".env")
}

func loadFile(filename string) error {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	for _, line := range strings.Split(strings.Trim(string(contents), "\n"), "\n") {
		if line == "" {
			continue
		}

		if err := loadLine(strings.Trim(line, "\n")); err != nil {
			return err
		}
	}

	return nil
}

func loadLine(line string) error {
	splitLine := strings.Split(line, "=")

	if len(splitLine) != 2 {
		return fmt.Errorf("invalid key/value pair: %s", line)
	}

	return os.Setenv(splitLine[0], splitLine[1])
}
