package savers

import (
	"crypto/aes"
	"crypto/cipher"
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

func (e *Encrypter) Init() {

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
	block, err := aes.NewCipher(e.key[:])
	if err != nil {
		return "", err
	}
	cipherText := make([]byte, aes.BlockSize+len(data))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCBCEncrypter(block, iv)
	stream.CryptBlocks(cipherText[aes.BlockSize:], data)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}
