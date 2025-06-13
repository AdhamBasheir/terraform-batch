package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

var (
	// Define color styles for output
	CyanBold  = color.New(color.FgCyan).Add(color.Bold)
	RedBold   = color.New(color.FgRed).Add(color.Bold)
	GreenBold = color.New(color.FgGreen).Add(color.Bold)

	// Define flags
	Dir          *string
	IgnoreErrors *bool
)

func init() {
	Dir = flag.String("dir", "", "Directories to run the Terraform command in (comma-separated)")
	IgnoreErrors = flag.Bool("ignore-errors", false, "Enable ignore-errors mode")
}

// Function to check if directory exists
func dirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// Function to displays the help information.
func displayHelp() {
	CyanBold.Println("\nterraform-batch - a CLI wrapper for running Terraform commands in multiple directories")
	fmt.Println("Usage:")
	fmt.Println("  terraform-batch -dir=dir1,dir2 -ignore-errors=true plan")
	fmt.Println("\nOptions:")
	fmt.Println("  -dir            Comma-separated list of directories")
	fmt.Println("  -ignore-errors  Continue even if a directory fails")
	fmt.Println()
}

// Function to run terraform command
func terraform(Dirs string, IgnoreErrors bool, tfCmdSlpit []string) error {
	failedDirs := []string{}

	if Dirs == "" {
		tfCmd := exec.Command("terraform", tfCmdSlpit...)
		tfCmd.Stdout = os.Stdout
		tfCmd.Stderr = os.Stderr

		if err := tfCmd.Run(); err != nil {
			return fmt.Errorf("terraform command failed")
		}
		return nil
	}

	for dir := range strings.SplitSeq(Dirs, ",") {
		if !dirExists(dir) {
			CyanBold.Printf("\nThe system cannot find the directory: %s\n", dir)

			if IgnoreErrors {
				failedDirs = append(failedDirs, dir)
				continue
			}
			return fmt.Errorf("failed to find required directory: %s", dir)
		}

		CyanBold.Printf("\nDirectory: %s\n", dir)
		CyanBold.Printf("Command: terraform %s\n", strings.Join(tfCmdSlpit, " "))
		tfCmd := exec.Command("terraform", tfCmdSlpit...)
		tfCmd.Dir = dir
		tfCmd.Stdout = os.Stdout
		tfCmd.Stderr = os.Stderr
		if err := tfCmd.Run(); err != nil {
			if !IgnoreErrors {
				return fmt.Errorf("terraform command failed in directory: %s", dir)
			}
			failedDirs = append(failedDirs, dir)
		}
	}

	if len(failedDirs) > 0 {
		return fmt.Errorf("terraform command failed in directories: %s", strings.Join(failedDirs, ", "))
	}

	return nil
}

func main() {
	flag.Parse()
	tfCmdRaw := strings.Join(flag.Args(), " ")

	// Check if terraform command is provided
	if tfCmdRaw == "" || tfCmdRaw == "help" {
		displayHelp()
		os.Exit(2)
	}

	// Check if the terraform command begins with terraform
	if flag.Arg(0) == "terraform" {
		CyanBold.Println("Don't type 'terraform' — just type the command (e.g., plan, apply)")
		displayHelp()
		os.Exit(1)
	}

	// Run the terraform command across the directories sequentially
	if err := terraform(*Dir, *IgnoreErrors, strings.Fields(tfCmdRaw)); err != nil {
		// Print error message
		RedBold.Printf("\n%v\n\n", err)
		os.Exit(1)
	}

	// Print success message
	GreenBold.Printf("\nSuccessfully executed terraform command in all specified directories!\n\n")
}
