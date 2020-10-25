package config

import (
	"crypto/ecdsa"

	"github.com/dgrijalva/jwt-go"
)

type SigningKey struct {
	publicKey  *ecdsa.PublicKey
	privateKey *ecdsa.PrivateKey
}

func NewSigningKey(private, public []byte) (*SigningKey, error) {
	pub, err := jwt.ParseECPublicKeyFromPEM(public)
	if err != nil {
		return nil, err
	}

	key, err := jwt.ParseECPrivateKeyFromPEM(private)
	if err != nil {
		return nil, err
	}

	return &SigningKey{
		publicKey:  pub,
		privateKey: key,
	}, nil
}

func (sk *SigningKey) GetPublicKey() *ecdsa.PublicKey {
	return sk.publicKey
}

func (sk *SigningKey) SignToken(token *jwt.Token) (string, error) {
	return token.SignedString(sk.privateKey)
}
