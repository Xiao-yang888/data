package utils

import  (
	"bytes"
	"encoding/binary"
)

/**
 *将一个int64转换为[]byte
 */
func Int64TOByte(num int64) ([]byte, error) {
	//Buffer: 缓冲区     buff：增益
	buff := new(bytes.Buffer)//通过new来实例化一个缓冲区
	//buff.Write()   通过一系列的Write方法向缓冲区里写入数据
	//buff.Bytes()   通过Bytes方法从缓冲区当中获取数据
	/**
	 *两种排列方式:
	 *  大端位序排列：BigEndian
	 *  小端位序排列：LittlieEndian
	 */
	err := binary.Write(buff, binary.BigEndian, nil)
	if err != nil {
		return nil, err
	}
	//从缓冲区当中读取数据
	return  buff.Bytes(), nil
}

/**
 *将string转换为[]byte
 */
func StringToBytes(data string) []byte {
	return []byte(data)
}