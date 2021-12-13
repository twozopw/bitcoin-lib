package blob

import (
	"errors"
	"github.com/twozopw/bitcoin-lib/src/utility"
	"io"
)

type Baseblob struct {
	data []byte
}

func DataReverse(dataIn []byte) []byte {
	var dataRet []byte
	dataRet = make([]byte, 0, len(dataIn))
	for i := len(dataIn) - 1; i >= 0; i-- {
		dataRet = append(dataRet, dataIn[i])
	}
	return dataRet
}

func (b *Baseblob) SetData(bytes []byte) {
	b.data = bytes
}

func (b *Baseblob) SetHex(hexStr string) error {
	if !utility.IsValidHex(hexStr) {
		return errors.New("invalid hex string")
	}
	if hexStr[0] == '0' && hexStr[1] == 'x' {
		hexStr = hexStr[2:]
	}
	blobLength := len(hexStr) / 2
	b.data = make([]byte, blobLength, blobLength)
	for i, j := blobLength-1, 0; i >= 0; i, j = i-1, j+1 {
		num1, _ := utility.HexCharToNumber(hexStr[2*i])
		num2, _ := utility.HexCharToNumber(hexStr[2*i+1])
		b.data[j] = byte((num1 << 4) | num2)
	}
	return nil
}

func (b Baseblob) GetHex() string {
	var bytes []byte
	dataRet := DataReverse(b.data)
	bytes = make([]byte, 2*len(dataRet), 2*len(dataRet))
	for i := 0; i < len(dataRet); i++ {
		var h4b byte
		var l4b byte
		h4b, _ = utility.NumberToHexChar((dataRet[i] & 0xf0) >> 4)
		l4b, _ = utility.NumberToHexChar(dataRet[i] & 0x0f)
		bytes[2*i], bytes[2*i+1] = h4b, l4b
	}
	return string(bytes)
}

func (b Baseblob) GetData() []byte {
	return b.data
}

func (b Baseblob) GetDataSize() int {
	return len(b.data)
}

func (b Baseblob) Pack(writer io.Writer, packSize int) error {
	_, err := writer.Write(b.data[0:packSize])
	if err != nil {
		return err
	}
	return nil
}

func (b *Baseblob) UnPack(reader io.Reader, unpackSize int) error {
	dataRead := make([]byte, unpackSize, unpackSize)
	_, err := reader.Read(dataRead[0:unpackSize])
	if err != nil {
		return err
	}
	b.data = dataRead
	return nil
}
