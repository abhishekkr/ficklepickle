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

	// RwDbDriver is db driver to picked if read/write mode is database, default: leveldb. Configurable via env var 'FICKLEPICKLE_DBDRIVER'
	RwDbDriver = golenv.OverrideIfEnv("FICKLEPICKLE_DBDRIVER", "leveldb")
	// DbName is name of database, default: fickle_pickle. Configurable via env var 'FICKLEPICKLE_DBNAME'
	DbName = golenv.OverrideIfEnv("FICKLEPICKLE_DBDRIVER", "fickle_pickle")

	// RpcServerPort configures the listen string to bind for RPC Server, default: localhost:8080. Configurable via env var 'FICKLEPICKLE_RPC_LISTENAT'
	RpcServerPort = golenv.OverrideIfEnv("FICKLEPICKLE_RPC_LISTENAT", "localhost:8080")
)

// PicklePath returns full file path to persist pickle in PickleDir using 'id' as filename.
func PicklePath(id string) string {
	return path.Join(PickleDir, id)
}

// DbPath returns full DB path required by golkeyval packagefor db persistence.
func DbPath() string {
	return path.Join(PickleDir, DbName)
}
