local p = import 'pkg/main.libsonnet';

p.pkg({
  source: 'https://github.com/marcbran/jsonnet-plugin-url',
  repo: 'https://github.com/marcbran/jsonnet.git',
  branch: 'plugin/url',
  path: 'plugin/url',
  target: 'url',
}, |||
  URL building, parsing and escaping backed by Go's `net/url`. Path segments and query parameters are escaped the way Go escapes them: a path segment keeps characters like `:`, `@` and `+`, while query keys and values escape them and encode spaces as `+`. Query parameters are objects whose values are strings, numbers, booleans or arrays of those; arrays become repeated keys and `null` values are omitted.
|||, {
  pathEscape: p.desc(|||
    Escapes `s` so it can be used as a single path segment, like Go's `url.PathEscape`.
  |||),
  pathUnescape: p.desc(|||
    Reverses `pathEscape`, like Go's `url.PathUnescape`.
  |||),
  queryEscape: p.desc(|||
    Escapes `s` so it can be used as a query key or value, like Go's `url.QueryEscape`.
  |||),
  queryUnescape: p.desc(|||
    Reverses `queryEscape`, like Go's `url.QueryUnescape`.
  |||),
  encodeQuery: p.desc(|||
    Encodes a params object into a query string sorted by key, like Go's `url.Values.Encode`.
  |||),
  parseQuery: p.desc(|||
    Parses a query string into an object whose values are always arrays of strings, like Go's `url.ParseQuery`.
  |||),
  format: p.desc(|||
    Formats a URL object with optional `scheme`, `host`, `path` (an array of unescaped segments), `params` and `fragment` into a string.
  |||),
  parse: p.desc(|||
    Parses a URL string into an object with `scheme`, `host`, `path`, `params` and `fragment`, the inverse of `format`.
  |||),
})
