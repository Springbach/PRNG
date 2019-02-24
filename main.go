package prng

import (
  "time"  
)

const maxOutputLength uint32 = 1024

//generator Algorithm interface
type ALG interface {
  Rnd([]byte, uint32) []byte
}

type Generator struct {
  Alg ALG
  SymbolSet []byte
  Elapsed time.Duration
  OutputLen uint32
  Type  string
}

func NewGEN(a ALG, set []byte, outLen uint32) *Generator {
  return &Generator{Alg: a, SymbolSet: set, OutputLen: outLen, Type: "Custom" }
}

//Pseudo Random generator
func (g *Generator) Generate() []byte {
  defer tracker(time.Now(), &g.Elapsed)
  return alg.Rnd(g.SymbolSet, g.OutputLen)
}

//elapsed time tracker
func tracker(start time.Time, elapsed *time.Duration) {
    *elapsed = time.Since(start)
}
