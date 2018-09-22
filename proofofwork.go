package main

import (
    "bytes"
    "crypto/sha256"
    "fmt"
    "math"
    "math/big"
    "strconv"
)

var (
    maxNonce = math.MaxInt64
)

const targetBits = 24

type ProofOfWork struct {
    block *Block
    // 요구사항
    target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
    target := big.NewInt(1)
    //bigInt.Lsh is >> left shift calculator
    target.Lsh(target, uint(256-targetBits))

    pow := &ProofOfWork{b, target}
    return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte{
    data := bytes.Join(
        [][]byte{
            pow.block.PrevBlockHash,
            pow.block.Data,
            []byte(strconv.FormatInt(pow.block.Timestamp, 16)),
            []byte(strconv.FormatInt(int64(targetBits), 16)),
            []byte(strconv.FormatInt(int64(nonce), 16)),
        },
        []byte{},
    )
    return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("golockchain block을 채굴합니다.  \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
    var hashInt big.Int

    data := pow.prepareData(pow.block.Nonce)
    hash := sha256.Sum256(data)
    hashInt.SetBytes(hash[:])

    isValid := hashInt.Cmp(pow.target) == -1
    return isValid
}
