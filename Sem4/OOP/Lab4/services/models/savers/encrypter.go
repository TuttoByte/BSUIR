package savers

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"encoding/base64"
	"io"
)

const KeySize = 32

type Encrypter struct {
	data []byte
	key  [32]byte
}

func NewEncrypter() *Encrypter {
	return &Encrypter{
		data: make([]byte, 0),
		key:  [32]byte{},
	}

}

func (e *Encrypter) GenerateKey() error {
	key := make([]byte, KeySize)
	_, err := rand.Read(key[:])
	if err != nil {
		return err
	}
	copy(e.key[:], key)
	return nil
}

func (e *Encrypter) GetKey() string {
	return base64.StdEncoding.EncodeToString(e.key[:])
}

func (e *Encrypter) Encrypt(data []byte) (string, error) {
	_, err := aes.NewCipher(e.key[:])
	if err != nil {
		return "", err
	}
	cipherText := make([]byte, aes.BlockSize+len(data))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func pkcs7Pad(data []byte, blockSize int) []byte {
	padLen := blockSize - len(data)%blockSize
	pad := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(data, pad...)
}
