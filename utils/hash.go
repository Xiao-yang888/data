package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	 "encoding/hex"
	"io"
	"io/ioutil"
)

/**
 *对一个字符串数据进行Md5哈希计算
 */
func MD5HashString(data string) (string) {
	md5Hash := md5.New()
	md5Hash.Write([]byte(data))
	bytes := md5Hash.Sum(nil)
	return hex.EncodeToString(bytes)
}
/**
 *io:input output 输入与输出
 */
func MD5HashReader(reader io.Reader) (string, error) {
	md5Hash := md5.New()
	readerBytes, err := ioutil.ReadAll(reader)
	//fmt.Println("读取到的文件：", readerBytes)
	if err != nil {
		return " ", err
	}
	md5Hash.Write(readerBytes)
	hashBytes := md5Hash.Sum(readerBytes)
	return hex.EncodeToString(hashBytes), nil
}

/*
 *读取io流当中的数据，并对数据进行hash计算，返回hash256 hash值
 */
func SHA256HashReader(reader io.Reader) (string, error) {
	sha256Hash := sha256.New()
	readerBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return " ", err
	}
	sha256Hash.Write(readerBytes)
	hashBytes := sha256Hash.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}

/**
 *对区块数据进行SHA256hash计算
 */
func SHA256HashBlock(block blockchain.Block) []byte {
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

	//2，将转换后的[]byte字节切片输入Write方法
	sha256Hash := sha256.New()
	sha256Hash.Write(blockBytes)
	hash := sha256Hash.Sum(nil)
	return hash
}