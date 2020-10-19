package blockchain

import "time"

/**
 *定义区块结构体，用于表示区块
 */
type Block struct {
	Height int64  //表示区块的高度，第几个区块
	TimeStamp int64 //时间戳，区块产生的时间
	PrevHash []byte  //前一个字段的hash
	Data []byte  //数据字段
	Hash []byte  //当前字段的hash
	Version string  //版本号
}

/**
 *创建一个新区快
 */
func NewBlock(height int64, prevHash []byte, data []byte) Block {
    block := Block{
		Height:    height,
		TimeStamp: time.Now().Unix(),
		PrevHash:  prevHash,
		Data:      data,
		Version:   "0*01",
	}
	//block.Hash =
	return block
}

/**
 *创建创世区块
 */
func CreateGenesisBlock() Block {
	genesisBlock := NewBlock(0,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},nil)
    return genesisBlock
}
