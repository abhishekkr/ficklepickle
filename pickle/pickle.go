package pickle

import (
	"bytes"
	gob "encoding/gob"

	config "github.com/abhishekkr/ficklepickle/config"

	golcompress "github.com/abhishekkr/gol/golcompress"
	golcrypt "github.com/abhishekkr/gol/golcrypt"
)

func PickleGob(data interface{}) (blob []byte, err error) {
	buf := new(bytes.Buffer)
	if err = gob.NewEncoder(buf).Encode(data); err != nil {
		return
	}
	blob = buf.Bytes()

	if config.Encryption == true {
		if blob, err = golcrypt.Encrypt(blob, config.Cookie, config.EncryptionScheme); err != nil {
			return
		}
	}

	if config.Compression == true {
		if blob, err = golcompress.Compress(blob, config.CompressWith); err != nil {
			return
		}
	}
	return
}
