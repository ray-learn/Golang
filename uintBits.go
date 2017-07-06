// 一元运算符^，按位补足。下面提供测试机器字长的方法。
// uint在32位操作系统上，使用32位。在64位操作系统上，使用64位。

package main

import "fmt"

func main() {
	var TestBit uint64
	// ^uint(0)>>63,若32位机，返回0,若64位机，返回1
	TestBit = uint64(^uint(0)) >> 63
	fmt.Printf("TestBit is %d:\n ", TestBit)

	// 根据上述返回值，与32基数做相乘即可
	var uintBits uint64
	uintBits = 32 * (1 + TestBit)
	fmt.Printf("BitSize is %d: \n", uintBits)

	// 上述可以简化如下，利用位右移运算
	var ConstuintBits uint64
	ConstuintBits = 32 << (uint64(^uint(0)) >> 63)
	fmt.Printf("ConstuintBits is %d:\n", ConstuintBits)
}
