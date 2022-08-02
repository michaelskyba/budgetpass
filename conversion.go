package main

import (
	"crypto/aes"
	"crypto/cipher"
)

func encrypt(key, plaintext []byte) []byte {
	myCipher, err := aes.NewCipher(key)
	hdl(err)

	gcm, err := cipher.NewGCM(myCipher)
	hdl(err)

	nonce := make([]byte, gcm.NonceSize())
	return gcm.Seal(nonce, nonce, plaintext, nil)
}

func decrypt(key, ciphertext []byte) []byte {
	myCipher, err := aes.NewCipher(key)
	hdl(err)

	gcm, err := cipher.NewGCM(myCipher)
	hdl(err)

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	hdl(err)

	return plaintext
}
