/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.8
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: tmp.i

package tmp

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


extern void _wrap_Swig_free_tmp_4231490e6333a8e7(uintptr_t arg1);
extern void _wrap_setMyString_tmp_4231490e6333a8e7(swig_voidp arg1);
#undef intgo

#cgo LDFLAGS: -L${SRCDIR}/cc -ltmp -lstdc++

*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"


type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


type _ sync.Mutex

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_tmp_4231490e6333a8e7(C.uintptr_t(_swig_i_0))
}

const X_GLIBCXX_STRING int = 1
func SetMyString(arg1 *string) {
	_swig_i_0 := arg1
	C._wrap_setMyString_tmp_4231490e6333a8e7(C.swig_voidp(_swig_i_0))
}

