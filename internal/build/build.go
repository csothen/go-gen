package build

import (
	"runtime/debug"
)

// Version is set at runtime
var Version = "dev"

// Date is set at runtime
var Date = "" // DD-MM-YYYY

func init() {
	if Version == "dev" {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "(devel)" {
			Version = info.Main.Version
		}
	}
}
