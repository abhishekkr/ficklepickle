package ficklepickle

import (
	"errors"

	database "github.com/abhishekkr/ficklepickle/database"
	deleter "github.com/abhishekkr/ficklepickle/deleter"
	pickle "github.com/abhishekkr/ficklepickle/pickle"
	reader "github.com/abhishekkr/ficklepickle/reader"
	unpickle "github.com/abhishekkr/ficklepickle/unpickle"
	writer "github.com/abhishekkr/ficklepickle/writer"

	gollog "github.com/abhishekkr/gol/gollog"
)

const (
	// RWVanillaFile is read/write to use simple one file per pickle.
	RwVanillaFile = "vanilla-file"

	// RwFile is alias to RwVanillaFile.
	RwFile = "vanilla-file"

	// RwDb is read/write mode to use a DB for pickle persistence (database type gets managed by config.RwDbDriver).
	RwDb = "database"
)

// Pickle returns pickled byte array and error state for provided 'data' as interface.
func Pickle(data interface{}) (blob []byte, err error) {
	return pickle.Gob(data)
}

// Unpickle returns error state for provided 'data' byte array, updates provided reference to interface for Go type.
func Unpickle(data []byte, skeleton interface{}) error {
	return unpickle.Gob(data, skeleton)
}

// Write returns error state for persisting pickle of provided 'data' as interface mapped to provided 'id' using 'mode'.
func Write(mode string, id string, data interface{}) error {
	blob, err := Pickle(data)
	if err != nil {
		return err
	}

	switch mode {
	case "vanilla-file":
		return writer.VanillaFile(id, blob)
	case "database":
		return writer.Database(id, blob)
	default:
		gollog.Err("unsupported write mode")
		return errors.New("ficklepickle: unsupported write mode")
	}
}

// Read returns error state for restoring pickle of provided 'id' using 'mode'; updates provided reference to 'skeleton' interface Go type.
func Read(mode string, id string, skeleton interface{}) error {
	var blob []byte
	var err error

	switch mode {
	case "vanilla-file":
		blob, err = reader.VanillaFile(id)
	case "database":
		blob, err = reader.Database(id)
	default:
		gollog.Err("unsupported read mode")
		err = errors.New("ficklepickle: unsupported read mode")
	}
	if err != nil {
		return err
	}

	err = Unpickle(blob, skeleton)
	return err
}

// Delete returns error state for removing pickle of provided provided 'id' using 'mode'.
func Delete(mode string, id string) error {
	switch mode {
	case "vanilla-file":
		return deleter.VanillaFile(id)
	case "database":
		return deleter.Database(id)
	default:
		gollog.Err("unsupported delete mode")
		return errors.New("ficklepickle: unsupported delete mode")
	}
}

func CloseDatabase() {
	database.Close()
}
