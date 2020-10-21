package blockchain

import (
	"data/utils"
	"time"
)

/**
 *定义区块结构体，用于表示区块
 */
type Block struct {
	Height    int64   //区块的高度，第几个区块
	TimeStamp int64   //时间戳，区块产生的时间
	PrevHash  []byte  //前一个字段的hash
	Data      []byte  //数据字段
	Hash      []byte  //当前字段的hash
	Version   string  //版本号
	Nonce     int64   //区块对应的nonce值
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

	//1,将block结构体数据转换为[]byte类型
	heightBytes, _ := Int64TOByte(block.Height)
	timeStampBytes, _ := Int64TOByte(block.TimeStamp)
	versionBytes := StringToBytes(block.Version)

	var blockBytes []byte
	//bytes.Join 拼接
	bytes.Join([] []byte{
		heightBytes,
		timeStampBytes,
		block.PrevHash,
		block.Data,
		versionBytes,
	}, []byte{})
	//调用Hash计算，对区块链进行SHA256计算
	block.Hash = utils.SHA256HashBlock(blockBytes)

	//挖矿竞争获得记账权

	return block
}

/**
 *创建创世区块
 */
func CreateGenesisBlock() Block {
	genesisBlock := NewBlock(0,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},nil)//创世区块
    return genesisBlock
}
