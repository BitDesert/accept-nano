package main

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	Index string `json:"index"`
	jwt.StandardClaims
}

func NewToken(index string) (string, error) {
	claims := MyCustomClaims{
		Index: index,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Seed))
}

func NewSeed() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return strings.ToUpper(hex.EncodeToString(b)), nil
}

func NewIndex() (string, error) {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	index := binary.LittleEndian.Uint64(b)
	return strconv.FormatUint(index, 10), nil
}