// 根据encoding/hex源码文件编写测试Encode函数。
// 一个byte占用8位bit，hex进制编码为4位bit一个hex位。
// 按照big-endian编码，若src[0]的二进制码为"10100101"，转换为十六进制应该为dst[0]="a",dst[1]="5"

package main

import (
	"fmt"
)

var hextable = [16]byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'a', 'b', 'c', 'd', 'e', 'f',
}

// 编码后的长度是编码前长度的2倍
func EncodedLen(n int) int { return n * 2 }

// 右移运算和位与运算
func Encode(dst, src []byte) ([]byte, int) {
	for i, v := range src {
		dst[i*2] = hextable[v>>4]
		dst[i*2+1] = hextable[v&0x0f]
	}
	return dst, len(src) * 2
}

func main() {
	var testByte []byte = make([]byte, 2)
	var encodedHex []byte = make([]byte, 4)
	var encodedLen int
	testByte[0] = 'a'
	testByte[1] = 'A'
	encodedHex, encodedLen = Encode(encodedHex, testByte)
	fmt.Printf("Before encode: %v\n", testByte)
	fmt.Printf("After encode:%v\n", encodedHex)
	fmt.Printf("Length before encode:%d\n", len(testByte))
	fmt.Printf("Length after encode:%d\n", encodedLen)
}
