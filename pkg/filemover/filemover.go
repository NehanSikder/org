package filemover

import (
	"fmt"
	"os"
)

// parameters fileMapping
// output using the provided number of threads,
func MoveFiles(workingDir string, fileMapping map[string]string, numberOfThreads int) {
	for file, newFolder := range fileMapping {
		targetDir := workingDir + "/" + newFolder
		// if folder doesnt exist create folder
		if _, err := os.Stat(targetDir); os.IsNotExist(err) {
			err := os.Mkdir(targetDir, 0755)
			if err != nil {
				fmt.Printf("ERROR: Unable to create %s \n", targetDir)
				continue
			}
		}
		err := os.Rename(workingDir+"/"+file, targetDir+"/"+file)

		if err != nil {
			fmt.Printf("Unable to move file %s to %s. Error: %s\n", file, targetDir+"/"+file, err)
			continue
		}
		fmt.Printf("Moved file %s to %s\n", file, workingDir+"/"+newFolder+"/"+file)
	}
}
