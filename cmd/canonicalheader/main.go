package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/golangci/canonicalheader"
)

func main() {
	singlechecker.Main(canonicalheader.New())
}
