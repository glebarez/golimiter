package linters

import (
	"go/ast"

	"github.com/mirecl/golimiter/internal/analysis"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/packages"
)

const (
	messageNoDefer = "a `defer` statement forbidden to use"
)

// NewNoDefer create instance linter for check defer.
func NewNoDefer() *analysis.Linter {
	return &analysis.Linter{
		Name: "NoDefer",
		Run: func(cfg *analysis.Config, pkgs []*packages.Package) []analysis.Issue {
			issues := make([]analysis.Issue, 0)

			for _, pkg := range pkgs {
				pkgIssues := runNoDefer(&cfg.NoDefer, pkg)
				issues = append(issues, pkgIssues...)
			}

			return issues
		},
	}
}

// TODO: check defer in func with name.
func runNoDefer(cfg *analysis.ConfigDefaultLinter, pkg *packages.Package) []analysis.Issue {
	nodeFilter := []ast.Node{(*ast.DeferStmt)(nil)}

	inspect := inspector.New(pkg.Syntax)

	var pkgIssues []analysis.Issue

	inspect.Preorder(nodeFilter, func(node ast.Node) {
		hash := analysis.GetHashFromBody(pkg.Fset, node)
		if cfg.IsVerifyHash(hash) {
			return
		}

		position := pkg.Fset.Position(node.Pos())

		pkgIssues = append(pkgIssues, analysis.Issue{
			Message:  messageNoDefer,
			Line:     position.Line,
			Filename: position.Filename,
			Hash:     hash,
		})
	})

	return pkgIssues
}
