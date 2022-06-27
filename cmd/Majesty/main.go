package main

import (
	"github.com/artoy/Majesty"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(Majesty.Analyzer) }
