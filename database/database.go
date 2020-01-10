package database

import (
	"encoding/base64"
	"errors"

	config "github.com/abhishekkr/ficklepickle/config"

	"github.com/abhishekkr/gol/golkeyval"
)

var (
	// dbDriver is global handle for golkeyval managed databaseengine.
	dbDriver golkeyval.DBEngine

	// dbOpened is global flag to manage state of DB driver initialization.
	dbOpened = false
)

// Open initialized db driver if not already done.
func Open() {
	if dbOpened == false {
		dbDriver = golkeyval.GetDBEngine("leveldb")
		dbDriver.Configure(map[string]string{
			"DBPath": config.DbPath(),
		})
		dbDriver.CreateDB()
		dbOpened = true
	}
}

// Close closes db if open.
func Close() {
	if dbOpened == true {
		dbDriver.CloseDB()
		dbOpened = false
	}
}

// Write returns error state for creating db entry of 'blob' mapped to 'id'.
func Write(id string, blob []byte) error {
	Open()
	encodedVal := base64.StdEncoding.EncodeToString(blob)
	if dbDriver.PushKeyVal(id, encodedVal) {
		return nil
	}
	return errors.New("ficklepickle: database write failed")
}

// Read returns byte array and error state for reading db entry mapped to 'id'.
func Read(id string) ([]byte, error) {
	Open()
	encodedVal := dbDriver.GetVal(id)
	return base64.StdEncoding.DecodeString(encodedVal)
}

// Delete returns error state for deleting db entry mapped to 'id'.
func Delete(id string) error {
	Open()
	if dbDriver.DelKey(id) {
		return nil
	}
	return errors.New("ficklepickle: database delete failed")
}
