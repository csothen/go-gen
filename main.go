package main

import (
	"github.com/csothen/go-gen/internal/build"
	"github.com/csothen/go-gen/pkg/cmd/root"
)

func main() {

	rootCmd := root.NewCmdRoot(build.Version, build.Date)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}

}
