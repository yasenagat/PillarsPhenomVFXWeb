package main

// #cgo LDFLAGS:  -L . -L ./Lib/linux64 -lr3d  -lR3DSDK -lstdc++ -lpthread -luuid -lm
// #include "r3d.h"
// #include <stdlib.h>
import "C"

import "fmt"
// import "unsafe"

func main() {
	var cmsg *C.char = C.BuildDate()

	var result string = C.GoString(cmsg)
	fmt.Println(result)
	//fmt.Println(BuildDate())
}

//func C.CString
