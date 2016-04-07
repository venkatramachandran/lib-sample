// package name: libsample.so
package main

import "C"
import "github.com/venkatramachandran/lib-sample/api"

//export Hello
func Hello(s *C.char) {
	api.Hello(C.GoString(s))
}

//export Hello2
func Hello2(s *C.char) *C.char {
	return C.CString(api.Hello2(C.GoString(s)))
}

func main() {
}
