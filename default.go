package prng

import (
    "crypto/rand"
    "encoding/base64"
)

type DefaultAlg struct {}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateCryptoRandom(n int) ([]byte, error) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    // Note that err == nil only if we read len(b) bytes.
    if err != nil {
        return nil, err
    }

    return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random byte array.
func (alg *DefaultAlg) Rnd(s []byte, outLen uint32) []byte {
    b, _ := GenerateCryptoRandom(int(outLen))
    return []byte(base64.URLEncoding.EncodeToString(b)[:outLen])
}
