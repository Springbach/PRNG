package prng

import (
  "math/rand"
  "time"
)

type SimpleAlg struct {}

//simple random generator
func (alg *SimpleAlg) Rnd(set []byte, outLen uint32) []byte {
  rand.Seed(time.Now().UnixNano())
  l := len(set)
  out := make([]byte, outLen)
	for indx := range out {
		out[indx] = set[rand.Intn(l)]
	}
	return out
}
