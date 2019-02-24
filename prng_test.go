package prng

import (
  "fmt"
  "testing"
  "strings"
)

type data struct {
    Num int32
    Set    []byte
    OutLen int
}

var alg ALG

var tests = []data{
    { 1,[]byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"), 32 },
    { 2,[]byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"), 64 },
    { 3,[]byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"), 128 },
    { 4,[]byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"), 16 },
    { 5,[]byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"), 12 },
    { 6,[]byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"), 32 },
    { 7,[]byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"), 24 },
    { 8,[]byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"), 12 },
    { 9,[]byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"), 64 },
    { 10,[]byte("abcdefghijABCDEFGHIJKL"), 25 },
    { 11,[]byte("abcdefghijABCDEFGHIJKL"), 35 },
    { 12,[]byte("abcdefghijABCDEFGHIJKL"), 45 },
    { 13,[]byte("abcdefghijABCDEFGHIJKL"), 1 },
    { 14,[]byte("abcdefghijABCDEFKL"), 2 },
    { 15,[]byte(""), 100 },
}

func TestGenerate(t *testing.T) {
    fmt.Println("--------------------Testing simple non crypto pseudorandom Alg")
    alg = &SimpleAlg{}

    for _, d := range tests {
        // TEST1 check len of output byte array for matching output condition
        gen := NewGEN(d.Set, uint32(d.OutLen))
        token := gen.Generate(alg)
        fmt.Printf("%d - %s - %s - %s\n", d.Num, gen.Type, gen.Elapsed, string(token))
        if len(token) != d.OutLen {
            t.Error(
                "For", d.Num,
                "expected length", d.OutLen,
                "got", len(token),
            )
        }
        // TEST2 check wether token contains only source symbol set
        for _, b := range token {
             //except Generator with default params
             if !strings.Contains(string(d.Set), string(b))&&len(d.Set)!=0 {
               t.Error(
                   "For", d.Num,
                   "expected token contains only source symbol set", string(d.Set),
                   "got", string(b),
               )
              }
        }
    }

    //////////////////////////////////////
    fmt.Println("-------------------------Testing crypto Alg")
    alg = &CryptoAlg{}
    for _, d := range tests {
        // check len of output byte array for matching output condition
        gen := NewGEN(d.Set, uint32(d.OutLen))
        token := gen.Generate(alg)
        fmt.Printf("%d - %s - %s - %s\n", d.Num, gen.Type, gen.Elapsed, string(token))
        if len(token) != d.OutLen {
            t.Error(
                "For", d.Num,
                "expected length", d.OutLen,
                "got", len(token),
            )
        }
        // check wether token contains only source symbol set
        for _, b := range token {
             if !strings.Contains(string(d.Set), string(b))&&len(d.Set)!=0 {
               t.Error(
                   "For", d.Num,
                   "expected token contains only source symbol set", string(d.Set),
                   "got", string(b),
               )
              }
        }
    }

}

///TODO for crypto algs
// Dieharder: A Random Number Test Suite by Robert G. Brown (rgb)
// agenda: http://webhome.phy.duke.edu/~rgb/General/dieharder.php
