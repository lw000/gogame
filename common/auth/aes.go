package tyauth

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
)

type AesStruct struct {
	key []byte
}

func NewAes(key []byte) *AesStruct {
	return &AesStruct{key: key}
}

func (a AesStruct) cbcEncrypt(data []byte, fillMode int) ([]byte, error) {
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()

	switch fillMode {
	case 0:
		data = ZeroPadding(data, blockSize)
	case 5:
		data = Pkcs5Padding(data, blockSize)
	case 7:
		data = Pkcs5Padding(data, blockSize)
	default:
		return nil, errors.New("nonsupport")
	}

	blockMode := cipher.NewCBCEncrypter(block, a.key[:blockSize])
	originData := make([]byte, len(data))
	blockMode.CryptBlocks(originData, data)

	return originData, nil
}

func (a AesStruct) cbcDecrypt(data []byte, fillMode int) ([]byte, error) {
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, a.key[:blockSize])

	var originData []byte
	originData = make([]byte, len(data))
	blockMode.CryptBlocks(originData, []byte(data))
	switch fillMode {
	case 0:
		originData = ZeroUnPadding(originData)
	case 5:
		originData = Pkcs5UnPadding(originData)
	case 7:
		originData = Pkcs5UnPadding(originData)
	default:
		return nil, errors.New("nonsupport")
	}

	return originData, nil
}

func AesEncrypt(key []byte, data []byte) (string, error) {
	aes := NewAes(key)
	originData, err := aes.cbcEncrypt(data, 5)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(originData), err
}

func AesDecrypt(key []byte, str string) ([]byte, error) {
	aes := NewAes(key)
	data, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}
	return aes.cbcDecrypt(data, 5)
}
