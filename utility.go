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
