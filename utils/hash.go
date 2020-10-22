package utils

import (
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
type Block struct {
	Height    int64   //区块的高度，第几个区块
	TimeStamp int64   //时间戳，区块产生的时间
	PrevHash  []byte  //前一个字段的hash
	Data      []byte  //数据字段
	Hash      []byte  //当前字段的hash
	Version   string  //版本号
}

func SHA256HashBlock(bs []byte) []byte {

	//2，将转换后的[]byte字节切片输入Write方法
	sha256Hash := sha256.New()
	sha256Hash.Write(bs)
	hash := sha256Hash.Sum(nil)
	return hash
}