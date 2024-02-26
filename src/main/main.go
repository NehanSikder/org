package main

import (
	"flag"
	"fmt"

	"github.com/NehanSikder/org/src/organizers/extension"
)

// func GetNewFileFolderMapping(fileOrganizer FileOrganizer, workdingDir *string) (map[string]string, error) {
// 	return fileOrganizer.GetNewFileFolderMapping(workdingDir)
// }

func main() {
	// Get directory path from first arg
	workingDir := flag.String("dir", "", "Absolute path to dir to organize")
	flag.Parse()
	if *workingDir == "" {
		fmt.Println("Please run command as org -dir=/path/to/dir/to/organize")
		return
	}
	// Read all files in directory

	// Figure out what folders to create and what file should go to what folder - this should be customizeable
	extFileOrganizer := extension.New(10, "NO_EXTENSION")
	fmt.Println("Generating new file folder mapping")
	newFileFolderMapping, err := extFileOrganizer.GetNewFileFolderMapping(*workingDir)
	// newFileFolderMapping, err := GetNewFileFolderMapping(extFileOrganizer, workingDir)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newFileFolderMapping)
	fmt.Println("Complete")
	// Show a simulation to the user
	// 	- how to make the simulation
	// Prompt user for confirmation
	// Push files into folders (create folders if it doesnt exist)

}
