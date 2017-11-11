package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

func main() {
	cstr := C.CString("你好，世界")
	C.puts(cstr)
	C.free(unsafe.Pointer(cstr))
}
