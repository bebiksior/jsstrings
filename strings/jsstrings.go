package jsstrings

import (
	"github.com/t14raptor/go-fast/ast"
	"github.com/t14raptor/go-fast/parser"
)

type StringExtractVisitor struct {
	ast.NoopVisitor
	Strings []JSStringWithURL
	URL     string
}

func (v *StringExtractVisitor) VisitExpression(n *ast.Expression) {
	if n == nil || n.Expr == nil {
		return
	}

	if strLit, ok := n.Expr.(*ast.StringLiteral); ok {
		if strLit != nil && strLit.Value != "" {
			startIdx := int(strLit.Idx0())
			endIdx := int(strLit.Idx1())
			locations := []Location{
				{StartIdx: startIdx,
					EndIdx: endIdx,
				},
			}
			v.Strings = append(v.Strings, JSStringWithURL{
				Value:     strLit.Value,
				Locations: locations,
				SourceURL: v.URL,
			})
		}
	}

	if tempLit, ok := n.Expr.(*ast.TemplateLiteral); ok {
		if tempLit != nil {
			var combinedValue string
			var locations []Location
			for _, elem := range tempLit.Elements {
				if elem.Valid {
					combinedValue += elem.Parsed
					locations = append(locations, Location{
						StartIdx: int(elem.Idx),
						EndIdx:   int(elem.Idx) + len(elem.Literal),
					})
				}
			}
			if combinedValue != "" {
				v.Strings = append(v.Strings, JSStringWithURL{
					Value:     combinedValue,
					Locations: locations,
					SourceURL: v.URL,
				})
			}
		}
	}

	n.VisitChildrenWith(v)
}

func (v *StringExtractVisitor) VisitProgram(program *ast.Program) {
	if program == nil {
		return
	}
	program.VisitChildrenWith(v)
}

// ExtractStrings parses JS content and extracts all string literals
func ExtractStrings(jsContent string, url string) ([]JSStringWithURL, error) {
	// Parse the JS content
	program, err := parser.ParseFile(jsContent)
	if err != nil {
		return nil, err
	}

	// Create and run the visitor
	visitor := &StringExtractVisitor{
		NoopVisitor: ast.NoopVisitor{},
		Strings:     make([]JSStringWithURL, 0),
		URL:         url,
	}
	visitor.V = visitor
	visitor.VisitProgram(program)

	return visitor.Strings, nil
}
