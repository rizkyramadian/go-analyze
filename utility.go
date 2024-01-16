package main

import "go/ast"

func receiverName(decl *ast.FuncDecl) string {
	if decl.Recv == nil {
		return ""
	}

	if len(decl.Recv.List) < 1 {
		return ""
	}

	return decl.Recv.List[0].Names[0].Name
}

func receiverType(decl *ast.FuncDecl) string {
	if decl.Recv == nil {
		return ""
	}

	if len(decl.Recv.List) < 1 {
		return ""
	}

	switch expr := decl.Recv.List[0].Type.(type) {
	case *ast.Ident:
		return expr.Name
	case *ast.StarExpr:
		if res, ok := expr.X.(*ast.Ident); ok {
			return res.Name
		}
	}
	return ""
}
