package main

import (
	"crypto/aes"
	"crypto/cipher"
)

func encrypt(key, plaintext []byte) []byte {
	myCipher, err := aes.NewCipher(key)
	handle(err)

	gcm, err := cipher.NewGCM(myCipher)
	handle(err)

	nonce := make([]byte, gcm.NonceSize())
	return gcm.Seal(nonce, nonce, plaintext, nil)
}

func decrypt(key, ciphertext []byte) []byte {
	myCipher, err := aes.NewCipher(key)
	handle(err)

	gcm, err := cipher.NewGCM(myCipher)
	handle(err)

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	handle(err)

	return plaintext
}
