package extension

import (
	"io/fs"
	"log"
	"os"
	"path"
	"strings"
	"sync"
)

type ExtFileOrganizer struct {
	numberOfThreads                   int
	wg                                *sync.WaitGroup
	defaultNoExtensionFoundFolderName string
}

func New(numberOfThreads int, defaultNoExtensionFoundFolderName string) *ExtFileOrganizer {
	return &ExtFileOrganizer{
		numberOfThreads:                   numberOfThreads,
		wg:                                new(sync.WaitGroup),
		defaultNoExtensionFoundFolderName: defaultNoExtensionFoundFolderName,
	}
}

var callOsReadDir = os.ReadDir

func getFiles(workdingDir string) []fs.DirEntry {
	files, err := callOsReadDir(workdingDir)
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
		} else {
			finalDirectory = strings.ToUpper(finalDirectory[1:])
		}
		mutex.Lock()
		mapping[fileName] = finalDirectory
		mutex.Unlock()
	}
	wg.Done()
}

var callGetFileNames = getFiles
var callParseExtension = parseExtension

func (ext *ExtFileOrganizer) GetNewFileFolderMapping(workdingDir string) (map[string]string, error) {
	mapping := make(map[string]string)
	// Read all files in working directory
	files := callGetFileNames(workdingDir)
	i := 0
	numberOfFiles := len(files)
	mutex := sync.RWMutex{}
	for i < numberOfFiles {
		for j := 0; j < ext.numberOfThreads; j++ {
			if i >= numberOfFiles {
				break
			}
			ext.wg.Add(1)
			go callParseExtension(mapping, files[i], ext.wg, ext.defaultNoExtensionFoundFolderName, &mutex)
			i += 1
		}
		ext.wg.Wait()
	}

	// Start a for loop and call go threads to read each file name and split the extension if it exists and update the map
	// wait for processing to finish
	return mapping, nil
}
