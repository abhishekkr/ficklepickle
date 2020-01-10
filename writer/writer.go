package writer

import (
	config "github.com/abhishekkr/ficklepickle/config"
	database "github.com/abhishekkr/ficklepickle/database"

	golfilesystem "github.com/abhishekkr/gol/golfilesystem"
)

var (
	writeInitialized = false
)

func initWrite() {
	if writeInitialized == false {
		golfilesystem.MkDirWithPermission(config.PickleDir, 0700)
		writeInitialized = true
	}
}

// VanillaFile returns error state for creating required pickle file with provided byte array.
func VanillaFile(id string, blob []byte) error {
	initWrite()
	filepath := config.PicklePath(id)
	return golfilesystem.CreateBinaryFile(filepath, blob)
}

// Database returns error state for creating db entry of 'blob' mapped to 'id'.
func Database(id string, blob []byte) error {
	initWrite()
	return database.Write(id, blob)
}
