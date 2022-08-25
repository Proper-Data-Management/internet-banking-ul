package main

import (
	"context"
	"runtime"

	"github.com/mak-alex/al_hilal_core/modules/daylight"
	_ "github.com/mattn/go-oci8"
)

func main() {
	ctx := context.Background()
	runtime.LockOSThread()
	daylight.Start(ctx)
}
