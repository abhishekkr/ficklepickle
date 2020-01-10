package deleter

import (
	"os"

	config "github.com/abhishekkr/ficklepickle/config"
	database "github.com/abhishekkr/ficklepickle/database"
)

// VanillaFile returns error state for deleting required pickle file inferred by its id.
func VanillaFile(id string) error {
	filepath := config.PicklePath(id)
	return os.Remove(filepath)
}

// Database returns error state for deleting db entry mapped to 'id'.
func Database(id string) error {
	return database.Delete(id)
}
