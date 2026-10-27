local p = import 'pkg/main.libsonnet';

p.ex({
  pathEscape: p.ex([{
    name: 'colon is kept, slash and space are escaped',
    inputs: ['system:node/a b'],
  }]),
  pathUnescape: p.ex([{
    name: 'escaped slash and space',
    inputs: ['a%2Fb%20c'],
  }]),
  queryEscape: p.ex([{
    name: 'space becomes plus, delimiters are escaped',
    inputs: ['a b&c=d'],
  }]),
  queryUnescape: p.ex([{
    name: 'plus becomes space',
    inputs: ['a+b%26c'],
  }]),
  encodeQuery: p.ex([{
    name: 'arrays repeat the key',
    inputs: [{ ids: ['1', '2'], q: 'a b' }],
  }]),
  parseQuery: p.ex([{
    name: 'repeated keys',
    inputs: ['ids=1&ids=2&q=a+b'],
  }]),
  format: p.ex([{
    name: 'path segments and params',
    inputs: [{ path: ['root', 'clusterrole', 'system:node'], params: { page: 2 } }],
  }, {
    name: 'absolute url',
    inputs: [{ scheme: 'https', host: 'example.com', path: ['docs', 'a/b'], fragment: 'top' }],
  }]),
  parse: p.ex([{
    name: 'absolute url',
    inputs: ['https://example.com/docs/a%2Fb?ids=1&ids=2#top'],
  }]),
})
