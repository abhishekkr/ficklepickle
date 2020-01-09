package unpickle

import (
	"bytes"
	gob "encoding/gob"

	config "github.com/abhishekkr/ficklepickle/config"

	golcompress "github.com/abhishekkr/gol/golcompress"
	golcrypt "github.com/abhishekkr/gol/golcrypt"
)

// Gob returns error state for updating provided reference of 'skeleton' inteface from gob-decoded unpickled byte arraycalongwith Decompression and Decryption.
func Gob(data []byte, skeleton interface{}) error {
	var err error

	if config.Compression == true {
		if data, err = golcompress.Decompress(data, config.CompressWith); err != nil {
			return err
		}

	}

	if config.Encryption == true {
		if data, err = golcrypt.Decrypt(data, config.Cookie, config.EncryptionScheme); err != nil {
			return err
		}
	}

	buf := bytes.NewBuffer(data)
	gob.NewDecoder(buf).Decode(skeleton)
	return nil
}
