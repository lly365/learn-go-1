package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func main() {
	testFile := "foo.txt"

	infile, inerr := os.Open(testFile)

	if inerr == nil {
		md5h := md5.New()
		io.Copy(md5h, infile)
		fmt.Printf("[md5] %x ---> %s\n", md5h.Sum([]byte("")), testFile)

		sha1h := sha1.New()
		io.Copy(sha1h, infile)
		fmt.Printf("[sha1] %x ---> %s\n", sha1h.Sum([]byte("")), testFile)
	} else {
		fmt.Println(inerr)
	}
}
