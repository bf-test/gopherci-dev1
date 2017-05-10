package main

/*
#include <gsl/gsl_version.h>

const char* ver() {
	return GSL_VERSION;
}
*/
import "C"

func main() {
	println(C.GoString(C.ver()))
}
