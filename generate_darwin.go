package uuidgen

/*
#include <stdlib.h>
*/
import "C"
import (
	"encoding/binary"
	"time"
	"unsafe"
)

func generate() ([16]byte, error) {
	b := [16]byte{}

	seed := time.Now().UnixNano()
	C.srandom(C.uint(seed))

	size := int(unsafe.Sizeof(C.int(0)))
	for i := 0; i < len(b)/size; i++ {
		v := uint32(C.random())
		binary.LittleEndian.PutUint32(b[i*size:(i*size)+size], v)
	}

	return b, nil
}
