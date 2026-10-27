package url

import (
	"fmt"
	neturl "net/url"
	"strings"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
)

func Format() jsonnet.NativeFunction {
	return jsonnet.NativeFunction{
		Name:   "format",
		Params: ast.Identifiers{"u"},
		Func: func(input []any) (any, error) {
			u, err := ToURL(input[0])
			if err != nil {
				return nil, err
			}
			return u.String(), nil
		},
	}
}

func Parse() jsonnet.NativeFunction {
	return jsonnet.NativeFunction{
		Name:   "parse",
		Params: ast.Identifiers{"s"},
		Func: func(input []any) (any, error) {
			s, ok := input[0].(string)
			if !ok {
				return nil, fmt.Errorf("s must be a string")
			}
			u, err := neturl.Parse(s)
			if err != nil {
				return nil, err
			}
			return FromURL(u)
		},
	}
}

// ToURL builds a url.URL from an object with optional scheme, host, path
// (an array of unescaped segments), params and fragment fields.
func ToURL(value any) (*neturl.URL, error) {
	obj, ok := value.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("url must be an object, got %#v", value)
	}
	scheme, err := optionalString(obj, "scheme")
	if err != nil {
		return nil, err
	}
	host, err := optionalString(obj, "host")
	if err != nil {
		return nil, err
	}
	fragment, err := optionalString(obj, "fragment")
	if err != nil {
		return nil, err
	}
	segments, err := pathSegments(obj["path"])
	if err != nil {
		return nil, err
	}
	values, err := ToValues(obj["params"])
	if err != nil {
		return nil, err
	}
	u := &neturl.URL{Scheme: scheme, Host: host, Fragment: fragment, RawQuery: values.Encode()}
	if len(segments) > 0 {
		escaped := make([]string, len(segments))
		for i, segment := range segments {
			escaped[i] = neturl.PathEscape(segment)
		}
		u.Path = "/" + strings.Join(segments, "/")
		u.RawPath = "/" + strings.Join(escaped, "/")
	}
	return u, nil
}

func FromURL(u *neturl.URL) (map[string]any, error) {
	path := []any{}
	if escaped := strings.TrimPrefix(u.EscapedPath(), "/"); escaped != "" {
		for _, part := range strings.Split(escaped, "/") {
			segment, err := neturl.PathUnescape(part)
			if err != nil {
				return nil, err
			}
			path = append(path, segment)
		}
	}
	return map[string]any{
		"scheme":   u.Scheme,
		"host":     u.Host,
		"path":     path,
		"params":   FromValues(u.Query()),
		"fragment": u.Fragment,
	}, nil
}

func optionalString(obj map[string]any, key string) (string, error) {
	switch v := obj[key].(type) {
	case nil:
		return "", nil
	case string:
		return v, nil
	default:
		return "", fmt.Errorf("%s must be a string, got %#v", key, v)
	}
}

func pathSegments(value any) ([]string, error) {
	switch v := value.(type) {
	case nil:
		return nil, nil
	case []any:
		segments := make([]string, len(v))
		for i, elem := range v {
			s, ok := elem.(string)
			if !ok {
				return nil, fmt.Errorf("path segments must be strings, got %#v", elem)
			}
			segments[i] = s
		}
		return segments, nil
	default:
		return nil, fmt.Errorf("path must be an array, got %#v", value)
	}
}
