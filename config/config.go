package config

import (
	"path"

	golenv "github.com/abhishekkr/gol/golenv"
)

var (
	Compression  = golenv.OverrideIfEnvBool("FICKLEPICKLE_COMPRESSION", true)
	CompressWith = golenv.OverrideIfEnv("FICKLEPICKLE_COMPRESS_WITH", "brotli")

	Encryption       = golenv.OverrideIfEnvBool("FICKLEPICKLE_ENCRYPTION", true)
	Cookie           = []byte(golenv.OverrideIfEnv("FICKLEPICKLE_COOKIE", "secret-cookie"))
	EncryptionScheme = golenv.OverrideIfEnv("FICKLEPICKLE_ENCRYPTION_SCHEME", "aes")

	PickleDir = golenv.OverrideIfEnv("FICKLEPICKLE_DIR", "./ficklepickle-data")
)

func PicklePath(id string) string {
	return path.Join(PickleDir, id)
}
