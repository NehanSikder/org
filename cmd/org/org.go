package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/NehanSikder/org/pkg/filemover"
	"github.com/NehanSikder/org/pkg/organizers/extension"
)

// func GetNewFileFolderMapping(fileOrganizer FileOrganizer, workdingDir *string) (map[string]string, error) {
// 	return fileOrganizer.GetNewFileFolderMapping(workdingDir)
// }

func showSimulation(workdingDir string, newFileFolderMapping map[string]string) (bool, error) {
	fmt.Println("Files will be moved to following directories")
	for file, folder := range newFileFolderMapping {
		fmt.Printf("%s to %s/%s\n", file, workdingDir, folder)
	}
	counter := 0
	for {
		fmt.Print("Do you want to proceed? (y/n): ")
		reader := bufio.NewReader(os.Stdin)
		userInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error getting user input: %s\n", err)
			return false, err
		}
		userInput = strings.ToLower(strings.TrimSpace(userInput))
		if userInput == "y" {
			fmt.Println("Proceeding...")
			// Add your code for "yes" option here
			return true, nil
		} else if userInput == "n" {
			fmt.Println("Exiting.")
			// Add your code for "no" option here
			return false, nil
		} else {
			counter += 1
			if counter == 3 {
				fmt.Println("Exiting after 3 attempts to get valid input")
				return false, nil
			}
			fmt.Println("Invalid input. Please enter 'y' or 'n'.")

		}
	}
	return false, nil
}

func main() {
	// Get directory path from first arg
	workingDir := flag.String("dir", "", "Absolute path to dir to organize")
	numberOfThreads := flag.Int("threadCount", 1, "Number of threads. Default: 1")
	flag.Parse()
	if *workingDir == "" {
		fmt.Println("Please run command as org -dir=/path/to/dir/to/organize")
		return
	}
	// Read all files in directory

	// Figure out what folders to create and what file should go to what folder - this should be customizeable
	extFileOrganizer := extension.New(*numberOfThreads, "NO_EXTENSION")
	fmt.Println("Generating new file folder mapping")
	newFileFolderMapping, err := extFileOrganizer.GetNewFileFolderMapping(*workingDir)
	// newFileFolderMapping, err := GetNewFileFolderMapping(extFileOrganizer, workingDir)
	if err != nil {
		fmt.Println(err)
	}

	// Show a simulation to the user
	// Prompt user for confirmation
	accept, err := showSimulation(*workingDir, newFileFolderMapping)
	if err != nil {
		fmt.Printf("Error while showing simulation to user. Error: %s", err)
		return
	}
	if !accept {
		return
	}

	// Push files into folders (create folders if it doesnt exist)
	filemover.MoveFiles(*workingDir, newFileFolderMapping, *numberOfThreads)
	fmt.Println("Complete")
}
