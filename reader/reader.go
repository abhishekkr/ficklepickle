package receiver

import (
	config "github.com/abhishekkr/ficklepickle/config"

	"github.com/abhishekkr/gol/golfilesystem"
)

func VanillaFile(id string) ([]byte, error) {
	filepath := config.PicklePath(id)
	return golfilesystem.ReadBinaryFile(filepath)
}
