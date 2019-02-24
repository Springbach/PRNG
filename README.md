<<<<<<< HEAD
﻿## PSEUDO RANDOM STRING GENERATOR (PRNG)
=======
<<<<<<< HEAD
﻿#PSEUDO RANDOM GENERATOR (prng) FOR GOLANG
=======
PSEUDO RANDOM GENERATOR (prng) FOR GOLANG

with fixed output length
and different algorithm usage possibility
>>>>>>> origin/master

for golang ![](https://golang.org/doc/gopher/frontpage.png =80x)

<<<<<<< HEAD
*fixed output length
different algorithm usage possibility
useful tool for auto generating crypto pass, ids, links and etc.*

**USAGE:**

1. simple non crypto alg
=======
USAGE:

1. simple non crypto alg

alg := &prng.SimpleAlg{}
set := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
gen := prng.NewGEN(set, 32)
token := gen.Generate(alg)
fmt.Println(string(token), len(token))
>>>>>>> origin/master

    alg := &prng.SimpleAlg{}
    set := []byte("abcdefghijklmnopqrstuvwxyz
    ABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
    gen := prng.NewGEN(set, 32)
    token := gen.Generate(alg)
    fmt.Println(string(token), len(token))
OUTPUT: abWWCa6AxRHmQnOdLHYdGMQeEGOvzhij    -32

<<<<<<< HEAD
2. crypto alg

    alg := &prng.CryptoAlg{}
    set := []byte("abcdefghijklmnopqrstuvwxyz
    ABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
    gen := prng.NewGEN(set, 32)
    token := gen.Generate(alg)
     fmt.Println(string(token), len(token))
OUTPUT: ldneTgPaBG3tHspoGZS5ox3ba3g8K19g       -32


3. default crypto base64 encoded alg for zero length symbol set

   alg := &prng.DefaultAlg{}
   set := []byte("") //ZERO LENGTH
   gen := prng.NewGEN(set, 100)
   token := gen.Generate(alg)
   fmt.Println(string(token), len(token))
OUTPUT: 2ZOeBKxUghxty200biKF2EyrRpHpKTA4QnbdISTzpo_
Kt6ZC8roE4Nrz3CXDR11lPxD7m4CUSCna5SqyDdL16d1cRoRQggzG9xqf  -100

**TEST COVERAGE 100%**
      | # | type | length  | elapsed | output
=======

2. crypto alg

alg := &prng.CryptoAlg{}
set := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
gen := prng.NewGEN(set, 32)
token := gen.Generate(alg)
fmt.Println(string(token), len(token))

OUTPUT: ldneTgPaBG3tHspoGZS5ox3ba3g8K19g 32


3. default crypto base64 encoded alg for zero length symbol set

alg := &prng.DefaultAlg{}
set := []byte("") //ZERO LENGTH
gen := prng.NewGEN(set, 100)
token := gen.Generate(alg)
fmt.Println(string(token), len(token))

OUTPUT: 2ZOeBKxUghxty200biKF2EyrRpHpKTA4QnbdISTzpo_Kt6ZC8roE4Nrz3CXDR11lPxD7m4CUSCna5SqyDdL16d1cRoRQggzG9xqf 100

TEST COVERAGE 100%
    type     length   elapsed    output


#   type     length   elapsed    output
>>>>>>> origin/master
* Testing simple non crypto pseudorandom Alg
1 - custom - len=32 - 13.071µs - dW5uaKGOVckzhwP1wsZE31shsvEN7EMX
2 - custom - len=64 - 9.374µs - M6UikIfTRUfHl6hW5CgZnTJd3rertaMdWCWLIRlkendzir7rUX4ECIZh9FQF5ikx
3 - custom - len=128 - 10.99µs - xwOq34g8yPXHwGPKHN7uGNKrNwqjocqSWLoTEowKsIjFMABXinY9PjFUi6q6ERohLHQifrp59UNdqsjWKOVdwU9xy5eg5ekOBhFP3Ym9eUp7XEJFbTx7qjD
D9oL7WmXH
4 - custom - len=16 - 8.301µs - eeNaDUDyEWWbO26U
5 - custom - len=12 - 8.101µs - R8NowhgClPVE
6 - custom - len=32 - 8.713µs - piWqktQSJILwMrMaovdCekvAPgzsMpPt
7 - custom - len=24 - 8.484µs - rnsdNmFgmcWuzYgqjJRoWIgb
8 - custom - len=12 - 8.122µs - nGvWJTFSgWIn
9 - custom - len=64 - 9.434µs - tCySNgpyMfXVpgoItNDhAkNKdsGMVWihsdtyIvXqDlskCAvgIKsMfGVVTcjcABJG
10 - custom - len=25 - 8.435µs - KcFEEHGefBFDEeLgaheaHCGig
11 - custom - len=35 - 8.674µs - JfAgFbHBcCgHgDbBdgHdFeJEGJdJALAFffD
12 - custom - len=45 - 8.893µs - CabjdhIJCFiBGEbjLFcGKieBEJCKgCLbffBdhIHfFadiD
13 - custom - len=1 - 7.839µs - h
14 - custom - len=2 - 7.868µs - Kg

* Testing crypto Alg
1 - custom - len=32 - 24.136µs - KIUOCLRq9M3OhxVczkEXUivRKhLdkgSc
2 - custom - len=64 - 51.448µs - Vullf4ZJt2qLKrDWpvswhujKnUvoGRUh1g9mzWnjh9C8NjqYiVzXTZWGCoUpvMDT
3 - custom - len=128 - 102.786µs - ZvajMriJhRxXyPetzsksMEQi8ywNYmQB92KPNciQ7PCQvKnQd1pS8qRurO8yknLwAy4tqAFZry2R8QziKO916bb7wpP5Tub4xhVgg3D7bvZY9p6mU8Kuv
uGmcCkjXm4W
4 - custom - len=16 - 11.471µs - jZ6TTd2zAnFp72v8
5 - custom - len=12 - 10.003µs - y9b3XIbauvhS
6 - custom - len=32 - 26.938µs - ingAbGxKkgyfHIFdWTTQBAkFkLONXvTw
7 - custom - len=24 - 22.423µs - YwCCGUPWPDEOEHscIqLKXEfW
8 - custom - len=12 - 13.195µs - vDrLbjSuEpwd
9 - custom - len=64 - 56.676µs - qdQvYjTBXSjPzGPHTFYPIcCucHXardSTczREhNRxITKOmroMBVdvAboFAqiXrngM
10 - custom - len=25 - 25.995µs - JCjDcfdgCfCFFHHALjjCDgKDd
11 - custom - len=35 - 35.723µs - aEJFBAaIiIJjfafEfCdggDfcaFiFFagDgde
12 - custom - len=45 - 53.99µs - KgCDHJfILgEABEGdaahGEaLBIAjiGLacKccaFhgBGfJHb
13 - custom - len=1 - 1.453µs - F
14 - custom - len=2 - 2.134µs - AK
15 - default - len=100 - 1.586µs - jSPjwc6ICCfZU4wWy7X8h6kqP2MKplJxgPA7rR_hmfX-At176dvAhx6Emf2AeYrCmxRbMcc8NxxWRsGJgJuPD_TCtpSNgGqsLnRu
PASS
ok      github.com/springbach/prng      0.002s

**TODO:** // MAKE TEST SUITE ANALOGUE OF DIEHARDER for strong crypto random string sequences
agenda: http://webhome.phy.duke.edu/~rgb/General/dieharder.php
Dieharder is a random number generator (rng) testing suite. It is intended to test generators, not files of possibly random numbers as the latter is a fallacious view of what it means to be random. Is the number 7 random? If it is generated by a random process, it might be. If it is made up to serve the purpose of some argument (like this one) it is not. Perfect random number generators produce "unlikely" sequences of random numbers -- at exactly the right average rate. Testing a rng is therefore quite subtle.
