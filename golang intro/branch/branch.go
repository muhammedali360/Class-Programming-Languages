package branch

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// branchCount counts the number of branching statements in the function.
func branchCount(fn *ast.FuncDecl) uint {
	count :=0
	ast.Inspect (fn, func (node ast.Node) bool  {
        switch node.(type) {
        case *ast.IfStmt:
            count += 1
				case *ast.ForStmt:
						count += 1
      	case *ast.SwitchStmt:
						count += 1
				case *ast.TypeSwitchStmt:
							count += 1
				case *ast.RangeStmt:
							count += 1
			}
			return true
		})
		return uint(count)
	}
// ComputeBranchFactors returns a map from the name of the function in the given
// Go code to the number of branching statements it contains.
func ComputeBranchFactors(src string) map[string]uint {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	m := make(map[string]uint)
	for _, decl := range f.Decls {
		switch fn := decl.(type) {
		case *ast.FuncDecl:
			m[fn.Name.Name] = branchCount(fn)
		}
	}

	return m
}
