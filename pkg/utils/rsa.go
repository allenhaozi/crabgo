package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"

	"github.com/pkg/errors"
)

type RSACipher struct {
}

func (rc RSACipher) Decrypt(cipherText []byte, priKey []byte) ([]byte, error) {
	block, _ := pem.Decode(priKey)

	if block == nil {
		return nil, errors.New("invalid private key")

	}

	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, private, cipherText)

}

func (rc RSACipher) Encrypt(plainText []byte, pubKey []byte) ([]byte, error) {

	block, _ := pem.Decode(pubKey)
	if block == nil {
		return nil, errors.New("invalid rsa public key")
	}

	pubInf, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInf.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, plainText)
}

func (rc RSACipher) Signature(plainText []byte, priKey []byte) ([]byte, error) {
	block, _ := pem.Decode(priKey)
	if block == nil {
		return nil, errors.New("invalid rsa public key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	// Before signing, we need to hash our message
	// The hash is what we actually sign
	msgHash := sha256.New()
	_, err = msgHash.Write(plainText)
	if err != nil {
		panic(err)
	}
	msgHashSum := msgHash.Sum(nil)

	// In order to generate the signature, we provide a random number generator,
	// our private key, the hashing algorithm that we used, and the hash sum
	// of our message
	return rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)
}

func (rc RSACipher) VerifySignature(plainText []byte, pubKey, signature []byte) error {
	block, _ := pem.Decode(pubKey)

	if block == nil {
		return errors.New("invalid rsa public key")
	}

	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)

	if err != nil {
		return errors.Wrap(err, "parse public key failure")
	}

	// Before signing, we need to hash our message
	// The hash is what we actually sign
	msgHash := sha256.New()
	_, err = msgHash.Write(plainText)
	if err != nil {
		return errors.Wrap(err, "verify signature failure")
	}
	msgHashSum := msgHash.Sum(nil)

	err = rsa.VerifyPSS(publicKey, crypto.SHA256, msgHashSum, signature, nil)
	if err != nil {
		return err
	}

	return nil
}
