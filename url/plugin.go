package url

import (
	"github.com/google/go-jsonnet"
	"github.com/marcbran/jpoet/pkg/jpoet"
)

func Plugin(opts ...jpoet.PluginOption) *jpoet.Plugin {
	return jpoet.NewPlugin("url", []jsonnet.NativeFunction{
		PathEscape(),
		PathUnescape(),
		QueryEscape(),
		QueryUnescape(),
		EncodeQuery(),
		ParseQuery(),
		Format(),
		Parse(),
	}, opts...)
}
