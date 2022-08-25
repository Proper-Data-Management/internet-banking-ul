package main

import (
	"context"
	"runtime"

	"github.com/internet-banking-ul/modules/daylight"
	_ "github.com/mattn/go-oci8"
)

func main() {
	ctx := context.Background()
	runtime.LockOSThread()
	daylight.Start(ctx)
}
