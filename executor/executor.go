// executor/executor.go

package executor

import (
	"fmt"
	"your_project_name/parser/ast"
	"your_project_name/executor/runtime"
)

type Executor struct {
	runtime *runtime.Runtime
}

func NewExecutor() *Executor {
	return &Executor{
		runtime: runtime.NewRuntime(),
	}
}

func (e *Executor) Execute(node ast.Node) interface{} {
	switch node := node.(type) {
	case *ast.Program:
		return e.executeProgram(node)
	case *ast.IntegerLiteral:
		return node.Value
	case *ast.Boolean:
		return node.Value
	case *ast.Identifier:
		return e.runtime.Get(node.Value)
	case *ast.LetStatement:
		e.runtime.Set(node.Name.Value, e.Execute(node.Value))
	// ... handle more node types ...

	default:
		fmt.Printf("Unknown node type: %T\n", node)
		return nil
	}
}

func (e *Executor) executeProgram(p *ast.Program) interface{} {
	var result interface{}
	for _, statement := range p.Statements {
		result = e.Execute(statement)
	}
	return result
}
