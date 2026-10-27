package url

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		name    string
		u       map[string]any
		want    string
		wantErr bool
	}{
		{
			name: "empty",
			u:    map[string]any{},
			want: "",
		},
		{
			name: "path only",
			u:    map[string]any{"path": []any{"root", "kubernetes", "pod-1"}},
			want: "/root/kubernetes/pod-1",
		},
		{
			name: "path segments keep sub-delimiters, colon and at",
			u:    map[string]any{"path": []any{"clusterrole", "system:node", "a@b", "a+b", "a=b&c"}},
			want: "/clusterrole/system:node/a@b/a+b/a=b&c",
		},
		{
			name: "slash, question mark and space inside a segment are escaped",
			u:    map[string]any{"path": []any{"a/b", "c?d", "e f"}},
			want: "/a%2Fb/c%3Fd/e%20f",
		},
		{
			name: "params are sorted, escaped and arrays repeat the key",
			u: map[string]any{
				"path":   []any{"x"},
				"params": map[string]any{"q": "a b&c", "ids": []any{"1", "2"}, "page": 2.0, "all": true},
			},
			want: "/x?all=true&ids=1&ids=2&page=2&q=a+b%26c",
		},
		{
			name: "null params are omitted",
			u:    map[string]any{"path": []any{"x"}, "params": map[string]any{"filter": nil}},
			want: "/x",
		},
		{
			name: "scheme, host and fragment",
			u: map[string]any{
				"scheme":   "https",
				"host":     "example.com",
				"path":     []any{"docs"},
				"fragment": "a b",
			},
			want: "https://example.com/docs#a%20b",
		},
		{
			name:    "non-string segment errors",
			u:       map[string]any{"path": []any{1.0}},
			wantErr: true,
		},
		{
			name:    "object param errors",
			u:       map[string]any{"params": map[string]any{"a": map[string]any{}}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := ToURL(tt.u)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, u.String())
		})
	}
}

func TestParseRoundTrip(t *testing.T) {
	in := map[string]any{
		"scheme":   "https",
		"host":     "example.com",
		"path":     []any{"a/b", "system:node", "e f"},
		"params":   map[string]any{"ids": []any{"1", "2"}, "q": []any{"a b&c"}},
		"fragment": "top",
	}
	u, err := ToURL(in)
	require.NoError(t, err)

	parsed, err := Parse().Func([]any{u.String()})
	require.NoError(t, err)
	require.Equal(t, in, parsed)
}

func TestParseEmptyPath(t *testing.T) {
	parsed, err := Parse().Func([]any{"https://example.com"})
	require.NoError(t, err)
	require.Equal(t, []any{}, parsed.(map[string]any)["path"])
}
