package Majesty

import (
	"fmt"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

const doc = "Majesty is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "Majesty",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
}

func dumpFunction(fs []*ssa.Function) error {
	for _, f := range fs {
		fmt.Println(f)
		for _, b := range f.Blocks {
			fmt.Printf("\tBlock %d\n", b.Index)
			for _, instr := range b.Instrs {
				fmt.Printf("\t\t%[1]T\t%[1]v\n", instr)
				for _, v := range instr.Operands(nil) {
					if v != nil {
						fmt.Printf("\t\t\t%[1]T\t%[1]v\n", *v)
					}
				}
			}
		}
	}

	return nil
}

// Entrypoint of analysis
func run(pass *analysis.Pass) (any, error) {
	s := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)

	dumpFunction(s.SrcFuncs)

	return nil, nil
}
