package utils

import (
	"context"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/itchyny/base58-go"

	"github.com/muhammad-asghar-ali/go/sysdesigns/shortly/internal/config"
)

func sha256Of(input string) []byte {
	a := sha256.New()
	a.Write([]byte(input))
	return a.Sum(nil)
}

func base58Encoding(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		return ""
	}

	return string(encoded)
}

func GenerateShortURL(long_url string) string {
	hash := sha256Of(long_url)
	generated := new(big.Int).SetBytes(hash).Uint64()

	encoded := base58Encoding([]byte(fmt.Sprintf("%d", generated)))[:7]
	return strings.ToLower(encoded)
}

func GenerateShortURLFromCounter() string {
	counter, _ := config.IncrementCounter(context.Background())

	hash := sha256Of(strconv.FormatInt(counter, 10))
	generated := new(big.Int).SetBytes(hash).Uint64()

	encoded := base58Encoding([]byte(fmt.Sprintf("%d", generated)))[:7]
	return strings.ToLower(encoded)
}
