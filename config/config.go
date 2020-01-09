package config

import (
	"path"

	golenv "github.com/abhishekkr/gol/golenv"
)

var (
	// Compression is bool config to enable/disable compression of pickled data, default: true. Configurable via env var 'FICKLEPICKLE_COMPRESSION'.
	Compression = golenv.OverrideIfEnvBool("FICKLEPICKLE_COMPRESSION", true)
	// CompressWith configures compression scheme if Compression is true, default: brotli. Configurable via env var 'FICKLEPICKLE_COMPRESS_WITH'.
	CompressWith = golenv.OverrideIfEnv("FICKLEPICKLE_COMPRESS_WITH", "brotli")

	// Encryption is bool config to enable/disable encryption of pickled data, default: true. Configurable via env var 'FICKLEPICKLE_ENCRYPTION'.
	Encryption = golenv.OverrideIfEnvBool("FICKLEPICKLE_ENCRYPTION", true)
	// Cookie is shared secret utilized to encrypt if Encryption is true, default: secret-cookie. Configurable via env var 'FICKLEPICKLE_COOKIE'.
	Cookie = []byte(golenv.OverrideIfEnv("FICKLEPICKLE_COOKIE", "secret-cookie"))
	// EncryptionScheme configures encryption scheme if Encryption is true, default: aes. Configurable via env var 'FICKLEPICKLE_ENCRYPTION_SCHEME'.
	EncryptionScheme = golenv.OverrideIfEnv("FICKLEPICKLE_ENCRYPTION_SCHEME", "aes")

	// PickleDir configures base directory to store pickled data in if 'Read/Write' gets utilized. Configurable via env var 'FICKLEPICKLE_DIR'.
	PickleDir = golenv.OverrideIfEnv("FICKLEPICKLE_DIR", "./ficklepickle-data")
)

// PicklePath returns full file path to persist pickle in PickleDir using 'id' as filename.
func PicklePath(id string) string {
	return path.Join(PickleDir, id)
}
