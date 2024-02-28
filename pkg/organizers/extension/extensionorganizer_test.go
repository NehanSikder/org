package extension

import (
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
	"time"
)

func TestNew(t *testing.T) {
	expectedNumberOfThreads := 10
	expectedDefaultNoExtensionFoundFolderName := "NO_EXTENSION"
	extFileOrganizer := New(expectedNumberOfThreads, expectedDefaultNoExtensionFoundFolderName)
	if extFileOrganizer.wg == nil {
		t.Error("new method should create ExtFileOrganizer object with wait group but object has wait group as NIL")
	}
	if extFileOrganizer.defaultNoExtensionFoundFolderName != expectedDefaultNoExtensionFoundFolderName {
		t.Errorf("Expected defaultNoExtensionFoundFolderName to be %s, got %s", expectedDefaultNoExtensionFoundFolderName, extFileOrganizer.defaultNoExtensionFoundFolderName)
	}
	if extFileOrganizer.numberOfThreads != expectedNumberOfThreads {
		t.Errorf("Expected numberOfThreads to be %d, got %d", expectedNumberOfThreads, extFileOrganizer.numberOfThreads)
	}

}

func createExpectedFileList(testFs fstest.MapFS) ([]fs.DirEntry, error) {
	fileList := []fs.DirEntry{}
	for k := range testFs {
		fi, err := fs.Stat(testFs, k)
		if err != nil {
			return nil, err
		}
		dirEntry := fs.FileInfoToDirEntry(fi)
		fileList = append(fileList, dirEntry)
	}

	// dirEntry := FileInfoToDirEntry(fi)

	return fileList, nil

}

func TestGetNewFileFolderMapping(t *testing.T) {
	// We are testing that given a mocked up file directory
	// the returned map matches our expected map
	expectedDefaultNoExtensionFoundFolderName := "NO_EXTENSION"
	expectedMap := make(map[string]string)
	expectedMap["test.txt"] = "TXT"
	expectedMap["test.log"] = "LOG"
	expectedMap["test"] = expectedDefaultNoExtensionFoundFolderName
	extFileOrganizer := New(10, expectedDefaultNoExtensionFoundFolderName)
	expectedWorkingDir := "/test"
	var sysValue int
	testFs := fstest.MapFS{
		"test.txt": {
			Data:    []byte("hello, world"),
			Mode:    0,
			ModTime: time.Now(),
			Sys:     &sysValue,
		},
		"test.log": {
			Data:    []byte("hello, world"),
			Mode:    0,
			ModTime: time.Now(),
			Sys:     &sysValue,
		},
		"test": {
			Data:    []byte("hello, world"),
			Mode:    0,
			ModTime: time.Now(),
			Sys:     &sysValue,
		},
	}
	expectedFilesList, err := createExpectedFileList(testFs)
	if err != nil {
		t.Errorf("createExpectedFileList returned error %s", err.Error())
	}
	// MOCK internal function call
	// callOsReadDir = func(name string) ([]os.DirEntry, error) { return expectedFilesList, nil }
	callGetFileNames = func(workdingDir string) []fs.DirEntry {
		return expectedFilesList
	}
	// Execute testing function call
	actualMap, err := extFileOrganizer.GetNewFileFolderMapping(expectedWorkingDir)
	if err != nil {
		t.Errorf("Expected error to be nil, got %s", err.Error())
	}
	// Validate output
	if reflect.DeepEqual(actualMap, expectedMap) == false {
		t.Errorf("Expected output map %s, got %s", expectedMap, actualMap)
	}
}
