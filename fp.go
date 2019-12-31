package ficklepickle

import (
	"errors"

	pickle "github.com/abhishekkr/ficklepickle/pickle"
	reader "github.com/abhishekkr/ficklepickle/reader"
	unpickle "github.com/abhishekkr/ficklepickle/unpickle"
	writer "github.com/abhishekkr/ficklepickle/writer"

	gollog "github.com/abhishekkr/gol/gollog"
)

func Pickle(data interface{}) (blob []byte, err error) {
	return pickle.PickleGob(data)
}

func Unpickle(data []byte, skeleton interface{}) error {
	return unpickle.UnpickleGob(data, skeleton)
}

func Write(mode string, id string, data interface{}) error {
	blob, err := Pickle(data)
	if err != nil {
		return err
	}

	switch mode {
	case "vanilla-file":
		return writer.VanillaFile(id, blob)
	default:
		gollog.Err("unsupported write mode")
		return errors.New("ficklepickle: unsupported write mode")
	}
}

func Read(mode string, id string, skeleton interface{}) error {
	var blob []byte
	var err error

	switch mode {
	case "vanilla-file":
		blob, err = reader.VanillaFile(id)
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
