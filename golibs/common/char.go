package common

import (
	"crypto/rand"
	"math/big"
)

func RandChar(length int, kind string) string {
	char := make([]rune, length)
	var codeModel []rune
	switch kind {
	case "num":
		codeModel = []rune("0123456789")
	case "char":
		codeModel = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	case "mix":
		codeModel = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	case "advance":
		codeModel = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+=-!@#$%*,.[]")
	default:
		codeModel = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}
	for i := range char {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(codeModel))))
		char[i] = codeModel[int(index.Int64())]
	}
	return string(char)
}

