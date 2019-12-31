package ficklepickle

import (
	receiver "github.com/abhishekkr/ficklepickle/receiver"
	sender "github.com/abhishekkr/ficklepickle/sender"

	golcompress "github.com/abhishekkr/gol/golcompress"
	golcrypt "github.com/abhishekkr/gol/golcrypt"
	golenv "github.com/abhishekkr/gol/golenv"
	gollog "github.com/abhishekkr/gol/gollog"
	jsoniter "github.com/json-iterator/go"
)

var (
	Compression  = golenv.OverrideIfEnvBool("FICKLEPICKLE_COMPRESSION", true)
	CompressWith = golenv.OverrideIfEnv("FICKLEPICKLE_COMPRESS_WITH", "brotli")

	Encryption       = golenv.OverrideIfEnvBool("FICKLEPICKLE_ENCRYPTION", true)
	Cookie           = []byte(golenv.OverrideIfEnv("FICKLEPICKLE_COOKIE", "secret-cookie"))
	EncryptionScheme = golenv.OverrideIfEnv("FICKLEPICKLE_ENCRYPTION_SCHEME", "aes")

	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func Pickle(data interface{}) (blob []byte, err error) {
	gollog.Info("pickle")
	if blob, err = json.Marshal(data); err != nil {
		return
	}
	if Encryption == true {
		if blob, err = golcrypt.Encrypt(blob, Cookie, EncryptionScheme); err != nil {
			return
		}
	}
	if Compression == true {
		if blob, err = golcompress.Compress(blob, CompressWith); err != nil {
			return
		}
	}
	return
}

func Unpickle(data []byte, skeleton interface{}) (interface{}, error) {
	gollog.Info("unpickle")
	var err error
	if Compression == true {
		if data, err = golcompress.Decompress(data, CompressWith); err != nil {
			return skeleton, err
		}

	}
	if Encryption == true {
		if data, err = golcrypt.Decrypt(data, Cookie, EncryptionScheme); err != nil {
			return skeleton, err
		}
	}
	if err = json.Unmarshal(data, &skeleton); err != nil {
		return skeleton, err
	}
	return skeleton, err
}

func Send() {
	gollog.Info("send")
	sender.Execute()
}

func Receive() {
	gollog.Info("receive")
	receiver.Execute()
}
