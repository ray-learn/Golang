// 十六进制编码，前缀设为0x。
// 编码规则1：十六进制数据必须有0x前缀；
// 编码规则2：字符切片的十六进制编码长度必须是偶数，空字符切片编码为"0x"
// 编码规则3：整型的十六进制编码长度必须是奇数，数值0编码为"0x0"

package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

// 定义机器字长
const uintBits = 32 << (uint64(^uint(0)) >> 63)

// 定义错误信息结构体
type decError struct {
	msg string
}

// 为decError定义Error方法
func (err decError) Error() string {
	return string(err.msg)
}

// 定义下述输入错误的变量
var (
	ErrEmptyString   = &decError{"empty hex string"}
	ErrSyntax        = &decError{"invalid hex string"}
	ErrMissingPrefix = &decError{"hex string without 0x prefix"}
	ErrOddLength     = &decError{"hex string of odd length"}
	ErrEmptyNumber   = &decError{"hex string \"0x\""}
	ErrLeadingZero   = &decError{"hex number with leading zero digits"}
	ErrUint64Range   = &decError{"hex number > 64 bits"}
	ErrUintRange     = &decError{fmt.Sprintf("hex number > %d bits", uintBits)}
	ErrBig256Range   = &decError{"hex number > 256 bits"}
)

// mapError处理其他解码时候的返回错误
func mapError(err error) error {
	if err, ok := err.(*strconv.NumError); ok {
		switch err.Err {
		case strconv.ErrRange:
			return ErrUint64Range
		case strconv.ErrSyntax:
			return ErrSyntax
		}
	}
	if _, ok := err.(hex.InvalidByteError); ok {
		return ErrSyntax
	}
	if err == hex.ErrLength {
		return ErrOddLength
	}
	return err
}

// has0xPrefix确认输入前缀是否为0x
func has0xPrefix(input string) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
}

// Decode函数解码0x的十六进制字符串
func Decode(input string) ([]byte, error) {
	if len(input) == 0 {
		return nil, ErrEmptyString
	}
	if !has0xPrefix(input) {
		return nil, ErrMissingPrefix
	}
	b, err := hex.DecodeString(input[2:])
	if err != nil {
		err = mapError(err)
	}
	return b, err
}

func main() {
	Result, _ := Decode("0x12345678")
	fmt.Printf("After Decode is %v\n", Result)
}
