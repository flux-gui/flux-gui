package main

import (
	"fmt"
	"os"

	"github.com/flux-gui/flux-gui/cmd/gitops/root"
)

func main() {
	if err := root.RootCmd().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
