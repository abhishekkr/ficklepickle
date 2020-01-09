package receiver

import (
	config "github.com/abhishekkr/ficklepickle/config"

	"github.com/abhishekkr/gol/golfilesystem"
)

// VanillaFile returns byte array and error state for reading required pickle file inferred by its id.
func VanillaFile(id string) ([]byte, error) {
	filepath := config.PicklePath(id)
	return golfilesystem.ReadBinaryFile(filepath)
}
