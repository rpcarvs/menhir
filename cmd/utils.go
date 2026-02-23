package cmd

import (
	"embed"
	"fmt"
	"os"
)

func copyEmbFile(fs embed.FS, srcPath string, outPath string) error {
	data, err := fs.ReadFile(srcPath)
	if err != nil {
		return fmt.Errorf("error reading embedded justfile")
	}

	if err := os.WriteFile(outPath, data, 0o644); err != nil {
		return fmt.Errorf("error writing justfile")
	}
	return nil
}
