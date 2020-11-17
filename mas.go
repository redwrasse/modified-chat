// A message authentication scheme consists of three algorithms:
// key generation, signing, and verification.
// code following https://leanpub.com/gocrypto/read#leanpub-auto-authenticity-and-integrity

package main

import (
	"crypto/rand"
	"errors"
	"golang.org/x/crypto/nacl/secretbox"
	"io"
)

const (
	KeySize = 32
	NonceSize = 24
)

// key generation
func GenerateKey() (*[KeySize]byte, error) {
	key := new([KeySize]byte)
	_, err := io.ReadFull(rand.Reader, key[:])
	if err != nil {
		return nil, err
	}
	return key, nil
}

// accompanying nonce for signing, verification
func GenerateNonce() (*[NonceSize]byte, error) {
	nonce := new([NonceSize]byte)
	_, err := io.ReadFull(rand.Reader, nonce[:])
	if err != nil {
		return nil, err
	}
	return nonce, nil
}

var (
	ErrEncrypt = errors.New("secret: encryption failed")
	ErrDecrypt = errors.New("secret: decryption failed")
)

// signing algorithm
func Sign(key *[KeySize]byte, message []byte) ([]byte, error) {
	nonce, err := GenerateNonce()
	if err != nil {
		return nil, ErrEncrypt
	}
	out := make([]byte, len(nonce))
	copy(out, nonce[:])
	out = secretbox.Seal(out, message, nonce, key)
	return out, nil
}

// verification algorithm
func Verify(key *[KeySize]byte, message []byte) ([]byte, error) {
	if len(message) < (NonceSize + secretbox.Overhead) {
		return nil, ErrDecrypt
	}
	var nonce [NonceSize]byte
	copy(nonce[:], message[:NonceSize])
	out, ok := secretbox.Open(nil, message[NonceSize:],
		&nonce, key)
	if !ok {
		return nil, ErrDecrypt
	}
	return out, nil
}


