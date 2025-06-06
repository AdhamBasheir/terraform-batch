package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/fatih/color"
)

var (
	// Define valid commands and their corresponding actions
	validCommands = map[string]bool{
		"init":     true, // Initialize Terraform configurations
		"validate": true, // Validate Terraform configurations
		"apply":    true, // Apply Terraform configurations
		"destroy":  true, // Destroy Terraform configurations
		"help":     true, // Display help information
	}

	// Define actions for each command
	actions = map[string]func(string) error{
		"init":     initTerraform,
		"validate": validateTerraform,
		"apply":    applyTerraform,
		"destroy":  destroyTerraform,
	}

	// Define color styles for output
	CyanBold  = color.New(color.FgCyan).Add(color.Bold)
	RedBold   = color.New(color.FgRed).Add(color.Bold)
	GreenBold = color.New(color.FgGreen).Add(color.Bold)
)

// Function to initialize Terraform configurations in the specified directory.
func initTerraform(dir string) error {
	// terraform init
	CyanBold.Printf("\nRunning terraform init in '%s'\n", dir)
	initCmd := exec.Command("terraform", "init")
	initCmd.Dir = dir
	initCmd.Stdout = os.Stdout
	initCmd.Stderr = os.Stderr
	if err := initCmd.Run(); err != nil {
		return fmt.Errorf("terraform init failed in '%s': %w", dir, err)
	}

	GreenBold.Printf("\nSuccessfully initialized Terraform in '%s'\n\n", dir)
	return nil
}

// Function to validate Terraform configurations in the specified directory.
func validateTerraform(dir string) error {
	// terraform validate
	CyanBold.Printf("\nRunning terraform validate in '%s'\n", dir)
	validateCmd := exec.Command("terraform", "validate")
	validateCmd.Dir = dir
	validateCmd.Stdout = os.Stdout
	validateCmd.Stderr = os.Stderr
	if err := validateCmd.Run(); err != nil {
		return fmt.Errorf("terraform validate failed in '%s': %w", dir, err)
	}

	return nil
}

// Function to apply Terraform configurations in the specified directory.
func applyTerraform(dir string) error {
	// terraform apply
	CyanBold.Printf("\nRunning terraform apply in '%s'\n", dir)
	applyCmd := exec.Command("terraform", "apply", "-auto-approve")
	applyCmd.Dir = dir
	applyCmd.Stdout = os.Stdout
	applyCmd.Stderr = os.Stderr
	applyCmd.Stdin = os.Stdin
	if err := applyCmd.Run(); err != nil {
		return fmt.Errorf("terraform apply failed in '%s': %w", dir, err)
	}

	GreenBold.Printf("\nSuccessfully applied Terraform in '%s'\n\n", dir)
	return nil
}

// Function to destroy Terraform configurations in the specified directory.
func destroyTerraform(dir string) error {
	// terraform destroy
	CyanBold.Printf("\nRunning terraform destroy in '%s'\n", dir)
	applyCmd := exec.Command("terraform", "destroy", "-auto-approve")
	applyCmd.Dir = dir
	applyCmd.Stdout = os.Stdout
	applyCmd.Stderr = os.Stderr
	applyCmd.Stdin = os.Stdin
	if err := applyCmd.Run(); err != nil {
		return fmt.Errorf("terraform destroy failed in '%s': %w", dir, err)
	}

	GreenBold.Printf("\nSuccessfully destroyed Terraform in '%s'\n\n", dir)
	return nil
}

// Function to displays the help information.
func displayHelp() {
	fmt.Println()
	fmt.Println("Terraform Automation Tool")
	fmt.Println("================================")
	fmt.Println("This Tool provides a simple interface to run Terraform commands in specified directories.")
	fmt.Println("Usage: terraform-batch <command> [directory1] [directory2] ...")
	fmt.Println()
	fmt.Println("Available commands:")
	fmt.Println("  init        Initialize Terraform configurations in the specified directories.")
	fmt.Println("  apply       Apply Terraform configurations in the specified directories.")
	fmt.Println("  destroy     Destroy Terraform configurations in the specified directories.")
	fmt.Println("  validate    Validate Terraform configurations in the specified directories.")
	fmt.Println("  help        Display this help information.")
	fmt.Println()
}

func main() {
	args := os.Args[1:]
	exit := false

	// Check if at least one argument is provided
	if len(args) == 0 {
		fmt.Println("Please provide a command to run and at least one directory.")
		fmt.Println("Usage: go run main.go <command> [directory1] [directory2] ...")
		os.Exit(1)
	}

	command := args[0]
	dirs := args[1:]

	// Check if the first argument is a valid command
	if _, valid := validCommands[command]; !valid {
		fmt.Printf("Invalid command: %s\n", command)
		displayHelp()
		os.Exit(1)
	}

	// If the command is "help", display the help information
	if command == "help" {
		displayHelp()
		os.Exit(0)
	}

	// Check if the directories are provided
	if len(dirs) == 0 && command != "help" {
		fmt.Println("Please provide at least one directory to run the command in.")
		fmt.Println("Usage: terraform-batch <command> [directory1] [directory2] ...")
		os.Exit(1)
	}

	// Check if the directories exist and are valid
	CyanBold.Println("Checking directories...")
	for _, baseDir := range dirs {
		dirInfo, dirErr := os.Stat(baseDir)
		fileInfo, fileErr := os.Stat(filepath.Join(baseDir, "main.tf"))
		if os.IsNotExist(dirErr) || !dirInfo.IsDir() {
			// Check if the directory exists
			fmt.Printf("Directory '%s' does not exist.\n", baseDir)
			exit = true
		} else if os.IsNotExist(fileErr) || fileInfo.IsDir() {
			// Check if the directory contains a main.tf file
			fmt.Printf("Directory '%s' does not have `main.tf` file.\n", baseDir)
			exit = true
		} else {
			// Directory exists and is valid
			fmt.Printf("Directory '%s' is valid.\n", baseDir)
		}
	}

	if exit {
		fmt.Println("Exiting due to invalid structure.")
		os.Exit(1)
	}

	// If all directories are valid, proceed with the command
	GreenBold.Printf("\nAll directories are valid. Proceeding with '%s' command...\n", command)

	// Execute the specified command in each directory
	for _, baseDir := range dirs {
		if err := actions[command](baseDir); err != nil {
			RedBold.Printf("%v\n\n", err)
			os.Exit(1)
		}
	}

	// Print success message
	GreenBold.Printf("\nSuccessfully executed '%s' command in all specified directories.\n\n", command)
}
