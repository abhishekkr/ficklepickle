package reader

import (
	config "github.com/abhishekkr/ficklepickle/config"
	database "github.com/abhishekkr/ficklepickle/database"
	ficklerpc "github.com/abhishekkr/ficklepickle/ficklerpc"

	"github.com/abhishekkr/gol/golfilesystem"
)

// VanillaFile returns byte array and error state for reading required pickle file inferred by its id.
func VanillaFile(id string) ([]byte, error) {
	filepath := config.PicklePath(id)
	return golfilesystem.ReadBinaryFile(filepath)
}

// Database returns byte array and error state for reading db entry mapped to 'id'.
func Database(id string) ([]byte, error) {
	return database.Read(id)
}

// Rpc returns byte array and error state for reading remote entry mapped to 'id'.
func Rpc(id string) ([]byte, error) {
	return ficklerpc.Read(id)
}
