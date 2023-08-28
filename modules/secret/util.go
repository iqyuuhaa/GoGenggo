package secret

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
)

func decryptGCMCipher(cipherText, secret []byte) (plainText []byte, err error) {
	block, err := aes.NewCipher(secret)
	if err != nil {
		log.Println("[Secret - decryptGCMCipher] Error new chiper secret, err: ", err)
		return plainText, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Println("[Secret - decryptGCMCipher] Error new gcm using block, err: ", err)
		return plainText, err
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		log.Println("[Secret - decryptGCMCipher] Not enough size on cipher text, err: ", err)
		return plainText, err
	}

	nonce, cipherTextOnly := cipherText[:nonceSize], cipherText[nonceSize:]
	plainText, err = gcm.Open(nil, nonce, cipherTextOnly, nil)
	if err != nil {
		log.Println("[Secret - decryptGCMCipher] Error open gcm secret, err: ", err)
		return plainText, err
	}

	return
}
