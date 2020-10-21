package blockchain

import (
	"math/big"
)

/**
 *工作量证明算法结构体
 */
type ProofOWork struct {
	Target *big.Int //系统的目标值
	Block  Block   //要找的noce值对应的区块

}

/**
 *实例化一个Pow算法的实例
 */
func NewPoW(block Block) ProofOWork {
	t := big.NewInt(1)
	t = t.Lsh(t, 255)
	pow := ProofOWork{
		Target: t,
		Block:  block,
	}
	return pow
}
