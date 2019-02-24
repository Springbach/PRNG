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
  SymbolSet []byte
  Elapsed   time.Duration
  OutputLen uint32
  Type      string
}

func NewGEN(set []byte, outLen uint32) *Generator {
  return &Generator{SymbolSet: set, OutputLen: outLen, Type: "default" }

}

//Pseudo Random generator
func (g *Generator) Generate(alg ALG) []byte {
  if len(g.SymbolSet) == 0 {
    alg = &DefaultAlg{}
  } else { g.Type = "custom"}
  defer tracker(time.Now(), &g.Elapsed)
  return alg.Rnd(g.SymbolSet, g.OutputLen)
}

//time tracker for generator
func tracker(start time.Time, elapsed *time.Duration) {
    *elapsed = time.Since(start)
}
