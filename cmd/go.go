package cmd

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var goCmd = &cobra.Command{
	Use:   "go [project_name]",
	Short: "Scaffold a go project",
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]
		// default

		fmt.Println("Preparing your go project...")

		// create dir
		dirCmd := exec.Command(
			"mkdir",
			"-p",
			projectName,
		)

		if err := dirCmd.Run(); err != nil {
			return fmt.Errorf("failed to create project dir %w", err)
		}

		// running go mod init
		modCmd := exec.Command(
			"go",
			"mod",
			"init",
			projectName)
		modCmd.Dir = projectName

		if err := modCmd.Run(); err != nil {
			return fmt.Errorf("failed to run go mod init %w", err)
		}

		// copying justfile
		src := "assets/justfile_go"
		out := filepath.Join(projectName, "justfile")

		if err := copyEmbFile(assets, src, out); err != nil {
			return err
		}

		// copying main.go
		src = "assets/main.go"
		out = filepath.Join(projectName, "main.go")

		if err := copyEmbFile(assets, src, out); err != nil {
			return err
		}

		fmt.Println("\nDone!!")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(goCmd)
}
