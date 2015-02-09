package r3dOperation

// #cgo LDFLAGS:  -L ../cppFile -L ../cppFile/Lib/linux64 -lr3d  -lR3DSDK -lstdc++ -lpthread -luuid -lm 
// #include "../cppFile/r3d.h"
// #include <stdlib.h>
import "C"

// import "fmt"
// import "unsafe"

func BuildDate() string {
	//the pointer returned can not be destroyed, they may use c_str() function, which is noe safe
	var cmsg *C.char = C.BuildDate()

	var result string = C.GoString(cmsg)
	return result
	//fmt.Println(BuildDate())
}


//func C.CString(goString string) *C.char
//func C.GoString(cString *C.char) string
//func C.GoString(cString *C.char, length C.int) string

//C.CString() need free mannul
//var cmsg * C.char = C.CString("hi")
//defer C.free(unsafe.Pointer(cmsg))
