package url

import (
	"fmt"
	neturl "net/url"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
)

func PathEscape() jsonnet.NativeFunction {
	return stringFunction("pathEscape", func(s string) (any, error) {
		return neturl.PathEscape(s), nil
	})
}

func PathUnescape() jsonnet.NativeFunction {
	return stringFunction("pathUnescape", func(s string) (any, error) {
		return neturl.PathUnescape(s)
	})
}

func QueryEscape() jsonnet.NativeFunction {
	return stringFunction("queryEscape", func(s string) (any, error) {
		return neturl.QueryEscape(s), nil
	})
}

func QueryUnescape() jsonnet.NativeFunction {
	return stringFunction("queryUnescape", func(s string) (any, error) {
		return neturl.QueryUnescape(s)
	})
}

func stringFunction(name string, fn func(s string) (any, error)) jsonnet.NativeFunction {
	return jsonnet.NativeFunction{
		Name:   name,
		Params: ast.Identifiers{"s"},
		Func: func(input []any) (any, error) {
			s, ok := input[0].(string)
			if !ok {
				return nil, fmt.Errorf("s must be a string")
			}
			return fn(s)
		},
	}
}
