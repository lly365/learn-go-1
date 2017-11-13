package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	str := "天涯共此时"
	md5Inst := md5.New()
	md5Inst.Write([]byte(str))
	result := md5Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", result)

	sha1Inst := sha1.New()
	sha1Inst.Write([]byte(str))
	result = sha1Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", result)
}
