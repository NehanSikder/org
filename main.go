package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"sync"
)

type FileOrganizer interface {
	GetNewFileFolderMapping(workdingDir string) (map[string]string, error)
}

type ExtFileOrganizer struct {
	numberOfThreads                   int
	wg                                *sync.WaitGroup
	defaultNoExtensionFoundFolderName string
}

func getFileNames(workdingDir string) []fs.DirEntry {
	files, err := os.ReadDir(workdingDir)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func parseExtension(mapping map[string]string, file fs.DirEntry, wg *sync.WaitGroup, defaultNoExtensionFoundFolderName string, mutex *sync.RWMutex) {
	if !file.IsDir() {
		fileName := file.Name()
		// extract extension
		finalDirectory := path.Ext(fileName)
		if finalDirectory == "" {
			finalDirectory = defaultNoExtensionFoundFolderName
		}
		mutex.Lock()
		mapping[fileName] = finalDirectory
		mutex.Unlock()
	}
	wg.Done()
}

func (ext *ExtFileOrganizer) GetNewFileFolderMapping(workdingDir string) (map[string]string, error) {
	mapping := make(map[string]string)
	// Read all files in working directory
	files := getFileNames(workdingDir)
	i := 0
	numberOfFiles := len(files)
	mutex := sync.RWMutex{}
	for i < numberOfFiles {
		for j := 0; j < ext.numberOfThreads; j++ {
			if i >= numberOfFiles {
				break
			}
			ext.wg.Add(1)
			go parseExtension(mapping, files[i], ext.wg, ext.defaultNoExtensionFoundFolderName, &mutex)
			i += 1
		}
		ext.wg.Wait()
	}

	// Start a for loop and call go threads to read each file name and split the extension if it exists and update the map
	// wait for processing to finish
	return mapping, nil
}

func GetNewFileFolderMapping(fileOrganizer FileOrganizer, workdingDir string) (map[string]string, error) {
	return fileOrganizer.GetNewFileFolderMapping(workdingDir)
}

func main() {
	// Get directory path from first arg
	workingDir := "/Users/arhamsikder/Desktop/go/org"
	// workingDir := flag.String("dir", "", "Absolute path to dir to organize")
	// flag.Parse()
	// if *workingDir == "" {
	// 	fmt.Println("Please run command as org -dir=/path/to/dir/to/organize")
	// }
	// Read all files in directory

	// Figure out what folders to create and what file should go to what folder - this should be customizeable

	extFileOrganizer := &ExtFileOrganizer{
		numberOfThreads:                   10,
		wg:                                new(sync.WaitGroup),
		defaultNoExtensionFoundFolderName: "NO_EXTENSION",
	}

	newFileFolderMapping, err := GetNewFileFolderMapping(extFileOrganizer, workingDir)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newFileFolderMapping)
	// Show a simulation to the user
	// 	- how to make the simulation
	// Prompt user for confirmation
	// Push files into folders (create folders if it doesnt exist)

}
