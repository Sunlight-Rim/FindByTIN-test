package main

import (
	gw "test-rusprofile/internal/api-gw"
	"test-rusprofile/internal/tin"
)

///	ENTRY POINT

func main() {
	tin.Start()
	gw.Start()
}
