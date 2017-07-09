// 接口的实现以及使用
package main

import (
	"fmt"
	"io"
)

// 定义一个缓冲器
type buf struct {
	str    []byte
	lhsize int
}

// 缓冲器实现io.Writer接口的方法，返回写入的长度与完成与否
func (self *buf) Write(b []byte) (int, error) {
	self.str = append(self.str, b...)
	return len(b), nil
}

// WriteHello写入"Hello"
func WriteHello(wh io.Writer) {
	// 将whello定义为实现io.Writer接口的*buf
	whello := wh.(*buf)
	// []byte为写入UTF-8的形式
	(*whello).str = []byte{'H', 'e', 'l', 'l', 'o'}
}

func main() {
	var whello buf
	fmt.Println([]byte(whello.str))
	WriteHello(&whello)
	fmt.Println([]byte(whello.str))
}
