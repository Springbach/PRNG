package prng

import (
  "crypto/rand"
  "math/big"
)

type CryptoAlg struct {}

//simple random generator
func (alg *CryptoAlg) Rnd(set []byte, outLen uint32) []byte {
  out := make([]byte, outLen)
	for indx := range out {
    out[indx] = set[cryptoRandPos(len(set))]
	}
	return out
}

func cryptoRandPos(max int) int {
    nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
    if err != nil {
        panic(err)
    }
    return int(nBig.Int64())
}
