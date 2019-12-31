package writer

import (
	config "github.com/abhishekkr/ficklepickle/config"

	"github.com/abhishekkr/gol/golfilesystem"
)

func VanillaFile(id string, blob []byte) error {
	golfilesystem.MkDirWithPermission(config.PickleDir, 0700)
	filepath := config.PicklePath(id)
	return golfilesystem.CreateBinaryFile(filepath, blob)
}
