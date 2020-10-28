package blockchain

import (
	"bytes"
	"data/utils"
	"fmt"
	"math/big"
)

//设置难度
const DIFFICULTY = 16

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

/**
 *pow算法： 用于寻找符合条件的nonce值
 */
func (p ProofOWork) Run() ([]byte, int64) {
	var nonce int64
	var blockHash []byte
	nonce = 0

	for {
		//不知道什么时候结束，使用无限循环
		block := p.Block
		heightBytes, _ := utils.Int64ToByte(block.Height)
		timeBytes, _ := utils.Int64ToByte(block.TimeStamp)
		versionBytes := utils.StringToBytes(block.Version)

		nonceBytes, _ := utils.Int64ToByte(nonce)
		//已有区块信息和尝试的nonce值的拼接信息
		blockBytes := bytes.Join([][]byte{
			heightBytes,
			timeBytes,
			block.PrevHash,
			block.Data,
			versionBytes,
			nonceBytes,
		}, []byte{})

		//区块和尝试的nonce值拼接后得到的Hash值
		blockHash = utils.SHA256HashBlock(blockBytes)

		target := p.Target     //目标值
		var hashBig *big.Int   //声明和定义
		hashBig = new(big.Int) //分配内存空间，为变量分配地址
		//...:invalid mempry or nil pointer dereference:空指针错误
		hashBig = hashBig.SetBytes(blockHash)
		fmt.Println("当前尝试的nonce值:", nonce)
		if hashBig.Cmp(target) == -1 {
			//停止寻找
			break
		}
		nonce++ //自增，继续寻找
	}
	///将找到的符合规则的nonce值返回
	return blockHash, nonce
}