package main

import (
	"crypto/aes"
	"crypto/cipher"
)

func encrypt(masterPassword, localPassword []byte) []byte {
	myCipher, err := aes.NewCipher(masterPassword)
	handle(err)

	gcm, err := cipher.NewGCM(myCipher)
	handle(err)

	nonce := make([]byte, gcm.NonceSize())
	return gcm.Seal(nonce, nonce, localPassword, nil)
}

func decrypt(masterPassword, encrypted []byte) []byte {
	myCipher, err := aes.NewCipher(masterPassword)
	handle(err)

	gcm, err := cipher.NewGCM(myCipher)
	handle(err)

	nonceSize := gcm.NonceSize()
	nonce, encrypted := encrypted[:nonceSize], encrypted[nonceSize:]

	localPassword, err := gcm.Open(nil, nonce, encrypted, nil)
	handle(err)

	return localPassword
}
