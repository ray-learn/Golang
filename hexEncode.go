// 根据encoding/hex源码文件编写测试Encode函数
package main

import (
	"fmt"
)

var hextable = [16]byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'a', 'b', 'c', 'd', 'e', 'f',
}

// 编码过后的长度通常是编码前长度的2倍
func EncodedLen(n int) int { return n * 2 }

// 编码
func Encode(dst, src []byte) int {
	for i, v := range src {
		dst[i*2] = hextable[v > 4]
		dst[i*2+1] = hextable[v&0x0f]
	}
	return dst, len(src) * 2
}

func main() {

}
