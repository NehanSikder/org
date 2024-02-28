package filemover

import (
	"os"
	"testing"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createEmptyFile(fileName string) {
	d := []byte("")
	check(os.WriteFile(fileName, d, 0644))
}

func TestMoveFiles(t *testing.T) {

	// SETUP
	// CREATE DIR IN CURRENT DIR
	currentTime := time.Now()
	workingDir := currentTime.Format("20060102150405")
	files := []string{"test1.txt", "test2.log", "test1"}
	err := os.Mkdir(workingDir, 0755)
	check(err)
	for _, file := range files {
		createEmptyFile(workingDir + "/" + file)
	}

	// CLEAN UP
	defer os.RemoveAll(workingDir)
	// EXPECTATION
	fileMapping := make(map[string]string)
	fileMapping["test1.txt"] = "TXT"
	fileMapping["test2.log"] = "LOG"
	fileMapping["test1"] = "NO_EXTENSION"
	// ACTION
	MoveFiles(workingDir, fileMapping, 5)
	// VERIFICATION
	for _, file := range files {
		// FILE SHOULDNT EXIST IN CURRENT DIRECTORY
		_, err := os.Stat(workingDir + "/" + file)
		if err == nil {
			t.Errorf("File %s in dir: %s when it should have moved\n", file, workingDir)
		}
		// FILE SHOULD EXIST IN WORKING DIRECTORY
		_, err = os.Stat(workingDir + "/" + fileMapping[file] + "/" + file)
		if err != nil {
			t.Errorf("File %s should be in dir: %s/%s but not found\n", file, workingDir, fileMapping[file])
		}

	}

}
