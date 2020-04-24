// extracted so it can be tested without requiring the worker runtime
export function getCodeSandboxLocation(path) {
  const prefix = 'https://codesandbox.io/s/github/'
  const parts = path.substr(1).split('/')

  if (parts.length === 0) return null
  if (parts.length === 1) return null
  if (parts.length === 2) return prefix + parts.join('/')
  if (parts.length === 3) return null

  if (parts[2] === 'tree') return prefix + parts.join('/')

  if (parts[2] === 'blob')
    return (
      prefix +
      [parts[0], parts[1], 'tree', parts[3]].join('/') +
      '?file=/' +
      parts.slice(4).join('/')
    )

  return null
}
