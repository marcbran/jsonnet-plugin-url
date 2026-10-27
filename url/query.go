package url

import (
	"fmt"
	neturl "net/url"
	"strconv"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
)

func EncodeQuery() jsonnet.NativeFunction {
	return jsonnet.NativeFunction{
		Name:   "encodeQuery",
		Params: ast.Identifiers{"params"},
		Func: func(input []any) (any, error) {
			values, err := ToValues(input[0])
			if err != nil {
				return nil, err
			}
			return values.Encode(), nil
		},
	}
}

func ParseQuery() jsonnet.NativeFunction {
	return jsonnet.NativeFunction{
		Name:   "parseQuery",
		Params: ast.Identifiers{"query"},
		Func: func(input []any) (any, error) {
			query, ok := input[0].(string)
			if !ok {
				return nil, fmt.Errorf("query must be a string")
			}
			values, err := neturl.ParseQuery(query)
			if err != nil {
				return nil, err
			}
			return FromValues(values), nil
		},
	}
}

// ToValues converts a jsonnet object into url.Values. Arrays become repeated
// keys and null values are omitted.
func ToValues(params any) (neturl.Values, error) {
	if params == nil {
		return neturl.Values{}, nil
	}
	obj, ok := params.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("params must be an object, got %#v", params)
	}
	values := neturl.Values{}
	for key, value := range obj {
		switch v := value.(type) {
		case nil:
			continue
		case []any:
			for _, elem := range v {
				s, err := scalarString(key, elem)
				if err != nil {
					return nil, err
				}
				values.Add(key, s)
			}
		default:
			s, err := scalarString(key, v)
			if err != nil {
				return nil, err
			}
			values.Add(key, s)
		}
	}
	return values, nil
}

func FromValues(values neturl.Values) map[string]any {
	out := make(map[string]any, len(values))
	for key, vs := range values {
		arr := make([]any, len(vs))
		for i, v := range vs {
			arr[i] = v
		}
		out[key] = arr
	}
	return out
}

func scalarString(key string, value any) (string, error) {
	switch v := value.(type) {
	case string:
		return v, nil
	case bool:
		return strconv.FormatBool(v), nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	default:
		return "", fmt.Errorf("invalid value for param %q: %#v", key, value)
	}
}
