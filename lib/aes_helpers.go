package lib

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/rabduljamal/backend-iot/config"
)

var aesKey = config.Config("AES_KEY") // Ganti dengan kunci AES-256 yang sesuai
// var aesKey = make([]byte, 32)

// var timestamp = []byte(time.Now().Format("20060102150405"))

func EncryptAES(data []byte, timestamp string) (string, error) {
	// Ubah timestamp menjadi kunci dengan panjang yang sesuai (misalnya, 32 byte untuk AES-256)
	// Konversi string heksadesimal ke byte array
	key, err := hex.DecodeString(aesKey)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Tambahkan padding PKCS#7 ke data
	blockSize := aes.BlockSize
	padding := blockSize - len(data)%blockSize
	if padding > 0 {
		pad := bytes.Repeat([]byte{byte(padding)}, padding)
		data = append(data, pad...)
	}

	ciphertext := make([]byte, blockSize+len(data))
	iv := ciphertext[:blockSize]
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[blockSize:], data)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func DecryptAES(encryptedData string, timestamp string) ([]byte, error) {
	// Ubah timestamp menjadi kunci dengan panjang yang sesuai (misalnya, 32 byte untuk AES-256)
	key, err := hex.DecodeString(aesKey)
	if err != nil {
		return nil, err
	}

	ciphertext, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext is too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return ciphertext, nil
}
